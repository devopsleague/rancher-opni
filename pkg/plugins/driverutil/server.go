package driverutil

import (
	"context"

	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Implements a subset of methods usually required by a driver which uses a DefaultingConfigTracker
// to manage its configuration. These implementations should not vary between drivers, so they are
// provided here as a convenience.
//
// Usage example:
//
//	type Driver struct {
//		foo.UnsafeFooServer
//
//		// package foo;
//		// message GetRequest {
//		//   core.Revision revision = 1;
//		// }
//		*driverutil.BaseConfigServer[*foo.ResetRequest, *foo.HistoryResponse, T]
//	}
//
//	func NewDriver() *Driver {
//		defaultStore := ...
//		activeStore := ...
//		return &Driver{
//			BaseConfigServer: driverutil.NewBaseConfigServer[*foo.ResetRequest, *foo.HistoryResponse](defaultStore, activeStore, flagutil.LoadDefaults)
//		}
//	}
func NewBaseConfigServer[
	R ResetRequestType[T],
	HR HistoryResponseType[T],
	T InstallableConfigType[T],
](
	defaultStore, activeStore storage.ValueStoreT[T],
	loadDefaultsFunc DefaultLoaderFunc[T],
) *BaseConfigServer[R, HR, T] {
	return &BaseConfigServer[R, HR, T]{
		tracker: NewDefaultingConfigTracker[T](defaultStore, activeStore, loadDefaultsFunc),
	}
}

type BaseConfigServer[
	R ResetRequestType[T],
	HR HistoryResponseType[T],
	T InstallableConfigType[T],
] struct {
	tracker *DefaultingConfigTracker[T]
}

// GetConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) GetConfiguration(ctx context.Context, in *GetRequest) (T, error) {
	return d.tracker.GetConfigOrDefault(ctx, in.GetRevision())

}

// GetDefaultConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) GetDefaultConfiguration(ctx context.Context, in *GetRequest) (T, error) {
	return d.tracker.GetDefaultConfig(ctx, in.GetRevision())
}

// Install implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) Install(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	var t T
	t = t.ProtoReflect().New().Interface().(T)
	d.setEnabled(t, lo.ToPtr(true))
	err := d.tracker.ApplyConfig(ctx, t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to install monitoring cluster: %s", err.Error())
	}
	return &emptypb.Empty{}, nil
}

// ResetDefaultConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) ResetDefaultConfiguration(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	if err := d.tracker.ResetDefaultConfig(ctx); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// ResetConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) ResetConfiguration(ctx context.Context, in R) (*emptypb.Empty, error) {
	if in.GetMask() != nil {
		if enabledField := util.FieldByName[T]("enabled"); enabledField != nil {
			var t T
			in.GetMask().Append(t, "enabled")
			patch := in.GetPatch()
			if patch.ProtoReflect().Has(enabledField) {
				patch.ProtoReflect().Clear(enabledField)
			}
		}
	}
	if err := d.tracker.ResetConfig(ctx, in.GetMask(), in.GetPatch()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SetConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) SetConfiguration(ctx context.Context, t T) (*emptypb.Empty, error) {
	d.setEnabled(t, nil)
	if err := d.tracker.ApplyConfig(ctx, t); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SetDefaultConfiguration implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) SetDefaultConfiguration(ctx context.Context, t T) (*emptypb.Empty, error) {
	d.setEnabled(t, nil)
	if err := d.tracker.SetDefaultConfig(ctx, t); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// Uninstall implements ConfigurableServerInterface.
func (d *BaseConfigServer[R, HR, T]) Uninstall(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	var t T
	t = t.ProtoReflect().New().Interface().(T)
	d.setEnabled(t, lo.ToPtr(false))
	err := d.tracker.ApplyConfig(ctx, t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to uninstall monitoring cluster: %s", err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (d *BaseConfigServer[R, HR, T]) setEnabled(t T, enabled *bool) {
	field := util.FieldByName[T]("enabled")
	if field == nil {
		return
	}
	msg := t.ProtoReflect()
	if msg.Has(field) && enabled == nil {
		msg.Clear(field)
	} else if enabled != nil {
		msg.Set(field, protoreflect.ValueOfBool(*enabled))
	}
}

func (k *BaseConfigServer[R, HR, T]) ConfigurationHistory(ctx context.Context, req *ConfigurationHistoryRequest) (HR, error) {
	revisions, err := k.tracker.History(ctx, req.GetTarget(), storage.IncludeValues(req.GetIncludeValues()))
	resp := util.NewMessage[HR]()
	if err != nil {
		return resp, err
	}
	entries := resp.ProtoReflect().Mutable(util.FieldByName[HR]("entries")).List()
	for _, rev := range revisions {
		if req.IncludeValues {
			spec := rev.Value()
			SetRevision(spec, rev.Revision(), rev.Timestamp())
			entries.Append(protoreflect.ValueOfMessage(spec.ProtoReflect()))
		} else {
			newSpec := util.NewMessage[T]()
			SetRevision(newSpec, rev.Revision(), rev.Timestamp())
			entries.Append(protoreflect.ValueOfMessage(newSpec.ProtoReflect()))
		}
	}
	return resp, nil
}

func (k *BaseConfigServer[R, HR, T]) Tracker() *DefaultingConfigTracker[T] {
	return k.tracker
}
