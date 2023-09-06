// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/internal/cortex/config/storage/storage.proto

package storage

import (
	storage "github.com/rancher/opni/pkg/storage"
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	lo "github.com/samber/lo"
	pflag "github.com/spf13/pflag"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	strings "strings"
	time "time"
)

func (in *AzureConfig) DeepCopyInto(out *AzureConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *AzureConfig) DeepCopy() *AzureConfig {
	return proto.Clone(in).(*AzureConfig)
}

func (in *Config) DeepCopyInto(out *Config) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *Config) DeepCopy() *Config {
	return proto.Clone(in).(*Config)
}

func (in *FilesystemConfig) DeepCopyInto(out *FilesystemConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *FilesystemConfig) DeepCopy() *FilesystemConfig {
	return proto.Clone(in).(*FilesystemConfig)
}

func (in *GcsConfig) DeepCopyInto(out *GcsConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *GcsConfig) DeepCopy() *GcsConfig {
	return proto.Clone(in).(*GcsConfig)
}

func (in *HttpConfig) DeepCopyInto(out *HttpConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *HttpConfig) DeepCopy() *HttpConfig {
	return proto.Clone(in).(*HttpConfig)
}

func (in *S3Config) DeepCopyInto(out *S3Config) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *S3Config) DeepCopy() *S3Config {
	return proto.Clone(in).(*S3Config)
}

func (in *S3SSEConfig) DeepCopyInto(out *S3SSEConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *S3SSEConfig) DeepCopy() *S3SSEConfig {
	return proto.Clone(in).(*S3SSEConfig)
}

func (in *SwiftConfig) DeepCopyInto(out *SwiftConfig) {
	out.Reset()
	proto.Merge(out, in)
}

func (in *SwiftConfig) DeepCopy() *SwiftConfig {
	return proto.Clone(in).(*SwiftConfig)
}

func (in *Config) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Config", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("s3"), &in.Backend), strings.Join(append(prefix, "backend"), "."), "Backend storage to use. Supported backends are: s3, gcs, azure, swift, filesystem.")
	if in.S3 == nil {
		in.S3 = &S3Config{}
	}
	fs.AddFlagSet(in.S3.FlagSet(append(prefix, "s3")...))
	if in.Gcs == nil {
		in.Gcs = &GcsConfig{}
	}
	fs.AddFlagSet(in.Gcs.FlagSet(append(prefix, "gcs")...))
	if in.Azure == nil {
		in.Azure = &AzureConfig{}
	}
	fs.AddFlagSet(in.Azure.FlagSet(append(prefix, "azure")...))
	if in.Swift == nil {
		in.Swift = &SwiftConfig{}
	}
	fs.AddFlagSet(in.Swift.FlagSet(append(prefix, "swift")...))
	if in.Filesystem == nil {
		in.Filesystem = &FilesystemConfig{}
	}
	fs.AddFlagSet(in.Filesystem.FlagSet(append(prefix, "filesystem")...))
	return fs
}

func (in *Config) RedactSecrets() {
	if in == nil {
		return
	}
	in.S3.RedactSecrets()
	in.Gcs.RedactSecrets()
	in.Azure.RedactSecrets()
	in.Swift.RedactSecrets()
}

func (in *Config) UnredactSecrets(unredacted *Config) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if err := in.S3.UnredactSecrets(unredacted.GetS3()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "s3." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Gcs.UnredactSecrets(unredacted.GetGcs()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "gcs." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Azure.UnredactSecrets(unredacted.GetAzure()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "azure." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if err := in.Swift.UnredactSecrets(unredacted.GetSwift()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "swift." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *S3Config) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("S3Config", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.Endpoint), strings.Join(append(prefix, "endpoint"), "."), "The S3 bucket endpoint. It could be an AWS S3 endpoint listed at https://docs.aws.amazon.com/general/latest/gr/s3.html or the address of an S3-compatible service in hostname:port format.")
	fs.Var(flagutil.StringPtrValue(nil, &in.Region), strings.Join(append(prefix, "region"), "."), "S3 region. If unset, the client will issue a S3 GetBucketLocation API call to autodetect it.")
	fs.Var(flagutil.StringPtrValue(nil, &in.BucketName), strings.Join(append(prefix, "bucket-name"), "."), "S3 bucket name")
	fs.Var(flagutil.StringPtrValue(nil, &in.SecretAccessKey), strings.Join(append(prefix, "secret-access-key"), "."), "\x1b[31m[secret]\x1b[0m S3 secret access key")
	fs.Var(flagutil.StringPtrValue(nil, &in.AccessKeyId), strings.Join(append(prefix, "access-key-id"), "."), "S3 access key ID")
	fs.Var(flagutil.BoolPtrValue(flagutil.Ptr(false), &in.Insecure), strings.Join(append(prefix, "insecure"), "."), "If enabled, use http:// for the S3 endpoint instead of https://. This could be useful in local dev/test environments while using an S3-compatible backend storage, like Minio.")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("v4"), &in.SignatureVersion), strings.Join(append(prefix, "signature-version"), "."), "The signature version to use for authenticating against S3. Supported values are: v4, v2.")
	fs.Var(flagutil.StringPtrValue(flagutil.Ptr("auto"), &in.BucketLookupType), strings.Join(append(prefix, "bucket-lookup-type"), "."), "The s3 bucket lookup style. Supported values are: auto, virtual-hosted, path.")
	if in.Sse == nil {
		in.Sse = &S3SSEConfig{}
	}
	fs.AddFlagSet(in.Sse.FlagSet(append(prefix, "sse")...))
	if in.Http == nil {
		in.Http = &HttpConfig{}
	}
	fs.AddFlagSet(in.Http.FlagSet(append(prefix, "http")...))
	return fs
}

func (in *S3Config) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetSecretAccessKey() != "" {
		in.SecretAccessKey = flagutil.Ptr("***")
	}
	in.Sse.RedactSecrets()
}

func (in *S3Config) UnredactSecrets(unredacted *S3Config) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetSecretAccessKey() == "***" {
		if unredacted.GetSecretAccessKey() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "secret_access_key"},
			})
		} else {
			*in.SecretAccessKey = *unredacted.SecretAccessKey
		}
	}
	if err := in.Sse.UnredactSecrets(unredacted.GetSse()); storage.IsDiscontinuity(err) {
		for _, sd := range status.Convert(err).Details() {
			if info, ok := sd.(*errdetails.ErrorInfo); ok {
				info.Metadata["field"] = "sse." + info.Metadata["field"]
				details = append(details, info)
			}
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *S3SSEConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("S3SSEConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.Type), strings.Join(append(prefix, "type"), "."), "Enable AWS Server Side Encryption. Supported values: SSE-KMS, SSE-S3.")
	fs.Var(flagutil.StringPtrValue(nil, &in.KmsKeyId), strings.Join(append(prefix, "kms-key-id"), "."), "KMS Key ID used to encrypt objects in S3")
	fs.Var(flagutil.StringPtrValue(nil, &in.KmsEncryptionContext), strings.Join(append(prefix, "kms-encryption-context"), "."), "\x1b[31m[secret]\x1b[0m KMS Encryption Context used for object encryption. It expects JSON formatted string.")
	return fs
}

func (in *S3SSEConfig) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetKmsEncryptionContext() != "" {
		in.KmsEncryptionContext = flagutil.Ptr("***")
	}
}

func (in *S3SSEConfig) UnredactSecrets(unredacted *S3SSEConfig) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetKmsEncryptionContext() == "***" {
		if unredacted.GetKmsEncryptionContext() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "kms_encryption_context"},
			})
		} else {
			*in.KmsEncryptionContext = *unredacted.KmsEncryptionContext
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *HttpConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("HttpConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](1*time.Minute+30*time.Second), &in.IdleConnTimeout), strings.Join(append(prefix, "idle-conn-timeout"), "."), "The time an idle connection will remain idle before closing.")
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](2*time.Minute), &in.ResponseHeaderTimeout), strings.Join(append(prefix, "response-header-timeout"), "."), "The amount of time the client will wait for a servers response headers.")
	fs.Var(flagutil.BoolPtrValue(flagutil.Ptr(false), &in.InsecureSkipVerify), strings.Join(append(prefix, "insecure-skip-verify"), "."), "If the client connects via HTTPS and this option is enabled, the client will accept any certificate and hostname.")
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](10*time.Second), &in.TlsHandshakeTimeout), strings.Join(append(prefix, "tls-handshake-timeout"), "."), "Maximum time to wait for a TLS handshake. 0 means no limit.")
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](1*time.Second), &in.ExpectContinueTimeout), strings.Join(append(prefix, "expect-continue-timeout"), "."), "The time to wait for a server's first response headers after fully writing the request headers if the request has an Expect header. 0 to send the request body immediately.")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](100), &in.MaxIdleConnections), strings.Join(append(prefix, "max-idle-connections"), "."), "Maximum number of idle (keep-alive) connections across all hosts. 0 means no limit.")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](100), &in.MaxIdleConnectionsPerHost), strings.Join(append(prefix, "max-idle-connections-per-host"), "."), "Maximum number of idle (keep-alive) connections to keep per-host. If 0, a built-in default value is used.")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](0), &in.MaxConnectionsPerHost), strings.Join(append(prefix, "max-connections-per-host"), "."), "Maximum number of connections per host. 0 means no limit.")
	return fs
}

func (in *GcsConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("GcsConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.BucketName), strings.Join(append(prefix, "bucket-name"), "."), "GCS bucket name")
	fs.Var(flagutil.StringPtrValue(nil, &in.ServiceAccount), strings.Join(append(prefix, "service-account"), "."), "\x1b[31m[secret]\x1b[0m JSON representing either a Google Developers Console client_credentials.json file or a Google Developers service account key file. If empty, fallback to Google default logic.")
	return fs
}

func (in *GcsConfig) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetServiceAccount() != "" {
		in.ServiceAccount = flagutil.Ptr("***")
	}
}

func (in *GcsConfig) UnredactSecrets(unredacted *GcsConfig) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetServiceAccount() == "***" {
		if unredacted.GetServiceAccount() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "service_account"},
			})
		} else {
			*in.ServiceAccount = *unredacted.ServiceAccount
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *AzureConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("AzureConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.StringPtrValue(nil, &in.AccountName), strings.Join(append(prefix, "account-name"), "."), "Azure storage account name")
	fs.Var(flagutil.StringPtrValue(nil, &in.AccountKey), strings.Join(append(prefix, "account-key"), "."), "\x1b[31m[secret]\x1b[0m Azure storage account key")
	fs.Var(flagutil.StringPtrValue(nil, &in.ContainerName), strings.Join(append(prefix, "container-name"), "."), "Azure storage container name")
	fs.Var(flagutil.StringPtrValue(nil, &in.EndpointSuffix), strings.Join(append(prefix, "endpoint-suffix"), "."), "Azure storage endpoint suffix without schema. The account name will be prefixed to this value to create the FQDN")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](20), &in.MaxRetries), strings.Join(append(prefix, "max-retries"), "."), "Number of retries for recoverable errors")
	fs.Var(flagutil.StringPtrValue(nil, &in.MsiResource), strings.Join(append(prefix, "msi-resource"), "."), "\x1b[31m[secret]\x1b[0m Azure storage MSI resource. Either this or account key must be set.")
	fs.Var(flagutil.StringPtrValue(nil, &in.UserAssignedId), strings.Join(append(prefix, "user-assigned-id"), "."), "Azure storage MSI resource managed identity client Id. If not supplied system assigned identity is used")
	if in.Http == nil {
		in.Http = &HttpConfig{}
	}
	fs.AddFlagSet(in.Http.FlagSet(append(prefix, "http")...))
	return fs
}

func (in *AzureConfig) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetAccountKey() != "" {
		in.AccountKey = flagutil.Ptr("***")
	}
	if in.GetMsiResource() != "" {
		in.MsiResource = flagutil.Ptr("***")
	}
}

func (in *AzureConfig) UnredactSecrets(unredacted *AzureConfig) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetAccountKey() == "***" {
		if unredacted.GetAccountKey() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "account_key"},
			})
		} else {
			*in.AccountKey = *unredacted.AccountKey
		}
	}
	if in.GetMsiResource() == "***" {
		if unredacted.GetMsiResource() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "msi_resource"},
			})
		} else {
			*in.MsiResource = *unredacted.MsiResource
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *SwiftConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("SwiftConfig", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](0), &in.AuthVersion), strings.Join(append(prefix, "auth-version"), "."), "OpenStack Swift authentication API version. 0 to autodetect.")
	fs.Var(flagutil.StringPtrValue(nil, &in.AuthUrl), strings.Join(append(prefix, "auth-url"), "."), "OpenStack Swift authentication URL")
	fs.Var(flagutil.StringPtrValue(nil, &in.Username), strings.Join(append(prefix, "username"), "."), "OpenStack Swift username.")
	fs.Var(flagutil.StringPtrValue(nil, &in.UserDomainName), strings.Join(append(prefix, "user-domain-name"), "."), "OpenStack Swift user's domain name.")
	fs.Var(flagutil.StringPtrValue(nil, &in.UserDomainId), strings.Join(append(prefix, "user-domain-id"), "."), "OpenStack Swift user's domain ID.")
	fs.Var(flagutil.StringPtrValue(nil, &in.UserId), strings.Join(append(prefix, "user-id"), "."), "OpenStack Swift user ID.")
	fs.Var(flagutil.StringPtrValue(nil, &in.Password), strings.Join(append(prefix, "password"), "."), "\x1b[31m[secret]\x1b[0m OpenStack Swift API key.")
	fs.Var(flagutil.StringPtrValue(nil, &in.DomainId), strings.Join(append(prefix, "domain-id"), "."), "OpenStack Swift user's domain ID.")
	fs.Var(flagutil.StringPtrValue(nil, &in.DomainName), strings.Join(append(prefix, "domain-name"), "."), "OpenStack Swift user's domain name.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ApplicationCredentialId), strings.Join(append(prefix, "application-credential-id"), "."), "OpenStack Swift application credential ID.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ApplicationCredentialName), strings.Join(append(prefix, "application-credential-name"), "."), "OpenStack Swift application credential name.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ApplicationCredentialSecret), strings.Join(append(prefix, "application-credential-secret"), "."), "\x1b[31m[secret]\x1b[0m OpenStack Swift application credential secret.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ProjectId), strings.Join(append(prefix, "project-id"), "."), "OpenStack Swift project ID (v2,v3 auth only).")
	fs.Var(flagutil.StringPtrValue(nil, &in.ProjectName), strings.Join(append(prefix, "project-name"), "."), "OpenStack Swift project name (v2,v3 auth only).")
	fs.Var(flagutil.StringPtrValue(nil, &in.ProjectDomainId), strings.Join(append(prefix, "project-domain-id"), "."), "ID of the OpenStack Swift project's domain (v3 auth only), only needed if it differs the from user domain.")
	fs.Var(flagutil.StringPtrValue(nil, &in.ProjectDomainName), strings.Join(append(prefix, "project-domain-name"), "."), "Name of the OpenStack Swift project's domain (v3 auth only), only needed if it differs from the user domain.")
	fs.Var(flagutil.StringPtrValue(nil, &in.RegionName), strings.Join(append(prefix, "region-name"), "."), "OpenStack Swift Region to use (v2,v3 auth only).")
	fs.Var(flagutil.StringPtrValue(nil, &in.ContainerName), strings.Join(append(prefix, "container-name"), "."), "Name of the OpenStack Swift container to put chunks in.")
	fs.Var(flagutil.IntPtrValue(flagutil.Ptr[int32](3), &in.MaxRetries), strings.Join(append(prefix, "max-retries"), "."), "Max retries on requests error.")
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](10*time.Second), &in.ConnectTimeout), strings.Join(append(prefix, "connect-timeout"), "."), "Time after which a connection attempt is aborted.")
	fs.Var(flagutil.DurationpbValue(flagutil.Ptr[time.Duration](5*time.Second), &in.RequestTimeout), strings.Join(append(prefix, "request-timeout"), "."), "Time after which an idle request is aborted. The timeout watchdog is reset each time some data is received, so the timeout triggers after X time no data is received on a request.")
	return fs
}

func (in *SwiftConfig) RedactSecrets() {
	if in == nil {
		return
	}
	if in.GetPassword() != "" {
		in.Password = flagutil.Ptr("***")
	}
	if in.GetApplicationCredentialSecret() != "" {
		in.ApplicationCredentialSecret = flagutil.Ptr("***")
	}
}

func (in *SwiftConfig) UnredactSecrets(unredacted *SwiftConfig) error {
	if in == nil {
		return nil
	}
	var details []protoiface.MessageV1
	if in.GetPassword() == "***" {
		if unredacted.GetPassword() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "password"},
			})
		} else {
			*in.Password = *unredacted.Password
		}
	}
	if in.GetApplicationCredentialSecret() == "***" {
		if unredacted.GetApplicationCredentialSecret() == "" {
			details = append(details, &errdetails.ErrorInfo{
				Reason:   "DISCONTINUITY",
				Metadata: map[string]string{"field": "application_credential_secret"},
			})
		} else {
			*in.ApplicationCredentialSecret = *unredacted.ApplicationCredentialSecret
		}
	}
	if len(details) == 0 {
		return nil
	}
	return lo.Must(status.New(codes.InvalidArgument, "cannot unredact: missing values for secret fields").WithDetails(details...)).Err()
}

func (in *FilesystemConfig) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("FilesystemConfig", pflag.ExitOnError)
	fs.SortFlags = true
	return fs
}
