// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/internal/cortex/config/compactor/compactor.proto

package compactor

import (
	_ "github.com/rancher/opni/internal/codegen/cli"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of compaction time ranges.
	BlockRanges []*durationpb.Duration `protobuf:"bytes,1,rep,name=block_ranges,json=blockRanges,proto3" json:"block_ranges,omitempty"`
	// Number of Go routines to use when syncing block index and chunks files from the long term storage.
	BlockSyncConcurrency *int32 `protobuf:"varint,2,opt,name=block_sync_concurrency,json=blockSyncConcurrency,proto3,oneof" json:"block_sync_concurrency,omitempty"`
	// Number of Go routines to use when syncing block meta files from the long term storage.
	MetaSyncConcurrency *int32 `protobuf:"varint,3,opt,name=meta_sync_concurrency,json=metaSyncConcurrency,proto3,oneof" json:"meta_sync_concurrency,omitempty"`
	// Minimum age of fresh (non-compacted) blocks before they are being processed. Malformed blocks older than the maximum of consistency-delay and 48h0m0s will be removed.
	ConsistencyDelay *durationpb.Duration `protobuf:"bytes,4,opt,name=consistency_delay,json=consistencyDelay,proto3" json:"consistency_delay,omitempty"`
	// The frequency at which the compaction runs
	CompactionInterval *durationpb.Duration `protobuf:"bytes,5,opt,name=compaction_interval,json=compactionInterval,proto3" json:"compaction_interval,omitempty"`
	// How many times to retry a failed compaction within a single compaction run.
	CompactionRetries *int32 `protobuf:"varint,6,opt,name=compaction_retries,json=compactionRetries,proto3,oneof" json:"compaction_retries,omitempty"`
	// Max number of concurrent compactions running.
	CompactionConcurrency *int32 `protobuf:"varint,7,opt,name=compaction_concurrency,json=compactionConcurrency,proto3,oneof" json:"compaction_concurrency,omitempty"`
	// How frequently compactor should run blocks cleanup and maintenance, as well as update the bucket index.
	CleanupInterval *durationpb.Duration `protobuf:"bytes,8,opt,name=cleanup_interval,json=cleanupInterval,proto3" json:"cleanup_interval,omitempty"`
	// Max number of tenants for which blocks cleanup and maintenance should run concurrently.
	CleanupConcurrency *int32 `protobuf:"varint,9,opt,name=cleanup_concurrency,json=cleanupConcurrency,proto3,oneof" json:"cleanup_concurrency,omitempty"`
	// Time before a block marked for deletion is deleted from bucket. If not 0, blocks will be marked for deletion and compactor component will permanently delete blocks marked for deletion from the bucket. If 0, blocks will be deleted straight away. Note that deleting blocks immediately can cause query failures.
	DeletionDelay *durationpb.Duration `protobuf:"bytes,10,opt,name=deletion_delay,json=deletionDelay,proto3" json:"deletion_delay,omitempty"`
	// For tenants marked for deletion, this is time between deleting of last block, and doing final cleanup (marker files, debug files) of the tenant.
	TenantCleanupDelay *durationpb.Duration `protobuf:"bytes,11,opt,name=tenant_cleanup_delay,json=tenantCleanupDelay,proto3" json:"tenant_cleanup_delay,omitempty"`
	// When enabled, mark blocks containing index with out-of-order chunks for no compact instead of halting the compaction.
	SkipBlocksWithOutOfOrderChunksEnabled *bool `protobuf:"varint,12,opt,name=skip_blocks_with_out_of_order_chunks_enabled,json=skipBlocksWithOutOfOrderChunksEnabled,proto3,oneof" json:"skip_blocks_with_out_of_order_chunks_enabled,omitempty"`
	// Number of goroutines to use when fetching/uploading block files from object storage.
	BlockFilesConcurrency *int32 `protobuf:"varint,13,opt,name=block_files_concurrency,json=blockFilesConcurrency,proto3,oneof" json:"block_files_concurrency,omitempty"`
	// Number of goroutines to use when fetching blocks from object storage when compacting.
	BlocksFetchConcurrency *int32 `protobuf:"varint,14,opt,name=blocks_fetch_concurrency,json=blocksFetchConcurrency,proto3,oneof" json:"blocks_fetch_concurrency,omitempty"`
	// When enabled, at compactor startup the bucket will be scanned and all found deletion marks inside the block location will be copied to the markers global location too. This option can (and should) be safely disabled as soon as the compactor has successfully run at least once.
	BlockDeletionMarksMigrationEnabled *bool `protobuf:"varint,15,opt,name=block_deletion_marks_migration_enabled,json=blockDeletionMarksMigrationEnabled,proto3,oneof" json:"block_deletion_marks_migration_enabled,omitempty"`
	// Comma separated list of tenants that can be compacted. If specified, only these tenants will be compacted by compactor, otherwise all tenants can be compacted. Subject to sharding.
	EnabledTenants []string `protobuf:"bytes,16,rep,name=enabled_tenants,json=enabledTenants,proto3" json:"enabled_tenants,omitempty"`
	// Comma separated list of tenants that cannot be compacted by this compactor. If specified, and compactor would normally pick given tenant for compaction (via -compactor.enabled-tenants or sharding), it will be ignored instead.
	DisabledTenants []string `protobuf:"bytes,17,rep,name=disabled_tenants,json=disabledTenants,proto3" json:"disabled_tenants,omitempty"`
	// How long block visit marker file should be considered as expired and able to be picked up by compactor again.
	BlockVisitMarkerTimeout *durationpb.Duration `protobuf:"bytes,18,opt,name=block_visit_marker_timeout,json=blockVisitMarkerTimeout,proto3" json:"block_visit_marker_timeout,omitempty"`
	// How frequently block visit marker file should be updated duration compaction.
	BlockVisitMarkerFileUpdateInterval *durationpb.Duration `protobuf:"bytes,19,opt,name=block_visit_marker_file_update_interval,json=blockVisitMarkerFileUpdateInterval,proto3" json:"block_visit_marker_file_update_interval,omitempty"`
	// When enabled, index verification will ignore out of order label names.
	AcceptMalformedIndex *bool `protobuf:"varint,20,opt,name=accept_malformed_index,json=acceptMalformedIndex,proto3,oneof" json:"accept_malformed_index,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetBlockRanges() []*durationpb.Duration {
	if x != nil {
		return x.BlockRanges
	}
	return nil
}

func (x *Config) GetBlockSyncConcurrency() int32 {
	if x != nil && x.BlockSyncConcurrency != nil {
		return *x.BlockSyncConcurrency
	}
	return 0
}

func (x *Config) GetMetaSyncConcurrency() int32 {
	if x != nil && x.MetaSyncConcurrency != nil {
		return *x.MetaSyncConcurrency
	}
	return 0
}

func (x *Config) GetConsistencyDelay() *durationpb.Duration {
	if x != nil {
		return x.ConsistencyDelay
	}
	return nil
}

func (x *Config) GetCompactionInterval() *durationpb.Duration {
	if x != nil {
		return x.CompactionInterval
	}
	return nil
}

func (x *Config) GetCompactionRetries() int32 {
	if x != nil && x.CompactionRetries != nil {
		return *x.CompactionRetries
	}
	return 0
}

func (x *Config) GetCompactionConcurrency() int32 {
	if x != nil && x.CompactionConcurrency != nil {
		return *x.CompactionConcurrency
	}
	return 0
}

func (x *Config) GetCleanupInterval() *durationpb.Duration {
	if x != nil {
		return x.CleanupInterval
	}
	return nil
}

func (x *Config) GetCleanupConcurrency() int32 {
	if x != nil && x.CleanupConcurrency != nil {
		return *x.CleanupConcurrency
	}
	return 0
}

func (x *Config) GetDeletionDelay() *durationpb.Duration {
	if x != nil {
		return x.DeletionDelay
	}
	return nil
}

func (x *Config) GetTenantCleanupDelay() *durationpb.Duration {
	if x != nil {
		return x.TenantCleanupDelay
	}
	return nil
}

func (x *Config) GetSkipBlocksWithOutOfOrderChunksEnabled() bool {
	if x != nil && x.SkipBlocksWithOutOfOrderChunksEnabled != nil {
		return *x.SkipBlocksWithOutOfOrderChunksEnabled
	}
	return false
}

func (x *Config) GetBlockFilesConcurrency() int32 {
	if x != nil && x.BlockFilesConcurrency != nil {
		return *x.BlockFilesConcurrency
	}
	return 0
}

func (x *Config) GetBlocksFetchConcurrency() int32 {
	if x != nil && x.BlocksFetchConcurrency != nil {
		return *x.BlocksFetchConcurrency
	}
	return 0
}

func (x *Config) GetBlockDeletionMarksMigrationEnabled() bool {
	if x != nil && x.BlockDeletionMarksMigrationEnabled != nil {
		return *x.BlockDeletionMarksMigrationEnabled
	}
	return false
}

func (x *Config) GetEnabledTenants() []string {
	if x != nil {
		return x.EnabledTenants
	}
	return nil
}

func (x *Config) GetDisabledTenants() []string {
	if x != nil {
		return x.DisabledTenants
	}
	return nil
}

func (x *Config) GetBlockVisitMarkerTimeout() *durationpb.Duration {
	if x != nil {
		return x.BlockVisitMarkerTimeout
	}
	return nil
}

func (x *Config) GetBlockVisitMarkerFileUpdateInterval() *durationpb.Duration {
	if x != nil {
		return x.BlockVisitMarkerFileUpdateInterval
	}
	return nil
}

func (x *Config) GetAcceptMalformedIndex() bool {
	if x != nil && x.AcceptMalformedIndex != nil {
		return *x.AcceptMalformedIndex
	}
	return false
}

var File_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6c, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x0e,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x1c, 0x8a, 0xc0, 0x0c, 0x18, 0x0a,
	0x16, 0x32, 0x68, 0x30, 0x6d, 0x30, 0x73, 0x2c, 0x31, 0x32, 0x68, 0x30, 0x6d, 0x30, 0x73, 0x2c,
	0x32, 0x34, 0x68, 0x30, 0x6d, 0x30, 0x73, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x8a, 0xc0, 0x0c, 0x04, 0x0a, 0x02, 0x32, 0x30, 0x48, 0x00,
	0x52, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x15, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x8a, 0xc0, 0x0c, 0x04, 0x0a, 0x02,
	0x32, 0x30, 0x48, 0x01, 0x52, 0x13, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x50, 0x0a, 0x11,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x08, 0x8a, 0xc0, 0x0c, 0x04, 0x0a, 0x02, 0x30, 0x73, 0x52, 0x10, 0x63, 0x6f,
	0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x58,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x8a, 0xc0, 0x0c, 0x08, 0x0a, 0x06, 0x31, 0x68,
	0x30, 0x6d, 0x30, 0x73, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0x8a, 0xc0, 0x0c, 0x03, 0x0a, 0x01, 0x33, 0x48, 0x02, 0x52,
	0x11, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0x8a, 0xc0, 0x0c, 0x03, 0x0a, 0x01, 0x31, 0x48, 0x03,
	0x52, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x10, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0b, 0x8a, 0xc0, 0x0c, 0x07, 0x0a, 0x05, 0x31, 0x35, 0x6d, 0x30, 0x73, 0x52, 0x0f, 0x63, 0x6c,
	0x65, 0x61, 0x6e, 0x75, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x3e, 0x0a,
	0x13, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x8a, 0xc0, 0x0c, 0x04,
	0x0a, 0x02, 0x32, 0x30, 0x48, 0x04, 0x52, 0x12, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x43,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x4f, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0d, 0x8a, 0xc0, 0x0c, 0x09, 0x0a, 0x07, 0x31, 0x32, 0x68, 0x30, 0x6d, 0x30, 0x73, 0x52,
	0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x59,
	0x0a, 0x14, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70,
	0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x8a, 0xc0, 0x0c, 0x08, 0x0a, 0x06, 0x36,
	0x68, 0x30, 0x6d, 0x30, 0x73, 0x52, 0x12, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x6d, 0x0a, 0x2c, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x0b, 0x8a, 0xc0, 0x0c, 0x07, 0x0a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x48, 0x05, 0x52, 0x25,
	0x73, 0x6b, 0x69, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x75,
	0x74, 0x4f, 0x66, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x17, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x8a, 0xc0, 0x0c, 0x04, 0x0a,
	0x02, 0x31, 0x30, 0x48, 0x06, 0x52, 0x15, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x46, 0x0a, 0x18, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f,
	0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0x8a, 0xc0, 0x0c, 0x03, 0x0a, 0x01, 0x33, 0x48, 0x07, 0x52, 0x16, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x64, 0x0a, 0x26, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0x8a, 0xc0, 0x0c, 0x07, 0x0a, 0x05, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x48, 0x08, 0x52, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x12, 0x62, 0x0a, 0x1a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x74,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0a, 0x8a, 0xc0, 0x0c, 0x06, 0x0a, 0x04, 0x35, 0x6d, 0x30, 0x73, 0x52, 0x17, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x56, 0x69, 0x73, 0x69, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x7a, 0x0a, 0x27, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x8a, 0xc0, 0x0c, 0x06, 0x0a, 0x04, 0x31, 0x6d, 0x30, 0x73, 0x52, 0x22, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x69, 0x73, 0x69, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x46, 0x0a, 0x16, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x6d, 0x61, 0x6c, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x0b, 0x8a, 0xc0, 0x0c, 0x07, 0x0a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x48, 0x09,
	0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65,
	0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x15,
	0x0a, 0x13, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x42, 0x16, 0x0a, 0x14, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x2f, 0x0a, 0x2d, 0x5f, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x42, 0x29, 0x0a, 0x27, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5f, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x19, 0x0a,
	0x17, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d,
	0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x42, 0x82, 0xc0, 0x0c, 0x04, 0x08, 0x01,
	0x10, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x74, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescData = file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDesc
)

func file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescData)
	})
	return file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDescData
}

var file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_goTypes = []interface{}{
	(*Config)(nil),              // 0: compactor.Config
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_depIdxs = []int32{
	1, // 0: compactor.Config.block_ranges:type_name -> google.protobuf.Duration
	1, // 1: compactor.Config.consistency_delay:type_name -> google.protobuf.Duration
	1, // 2: compactor.Config.compaction_interval:type_name -> google.protobuf.Duration
	1, // 3: compactor.Config.cleanup_interval:type_name -> google.protobuf.Duration
	1, // 4: compactor.Config.deletion_delay:type_name -> google.protobuf.Duration
	1, // 5: compactor.Config.tenant_cleanup_delay:type_name -> google.protobuf.Duration
	1, // 6: compactor.Config.block_visit_marker_timeout:type_name -> google.protobuf.Duration
	1, // 7: compactor.Config.block_visit_marker_file_update_interval:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_init() }
func file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_init() {
	if File_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto = out.File
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_rawDesc = nil
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_goTypes = nil
	file_github_com_rancher_opni_internal_cortex_config_compactor_compactor_proto_depIdxs = nil
}