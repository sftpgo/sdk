// Copyright (C) 2022-2023 Nicola Murino
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import "github.com/sftpgo/sdk/kms"

// FilesystemProvider defines the supported storage filesystems
type FilesystemProvider int

// supported values for FilesystemProvider
const (
	LocalFilesystemProvider     FilesystemProvider = iota // Local
	S3FilesystemProvider                                  // AWS S3 compatible
	GCSFilesystemProvider                                 // Google Cloud Storage
	AzureBlobFilesystemProvider                           // Azure Blob Storage
	CryptedFilesystemProvider                             // Local encrypted
	SFTPFilesystemProvider                                // SFTP
	HTTPFilesystemProvider                                // HTTP
)

// IsProviderSupported returns true if the specified provider is supported.
func IsProviderSupported(provider FilesystemProvider) bool {
	return provider >= LocalFilesystemProvider && provider <= HTTPFilesystemProvider
}

// OSFsConfig defines the configuration for local filesystem
type OSFsConfig struct {
	// Read buffer size as bytes, 0 means no buffering
	ReadBufferSize int `json:"read_buffer_size,omitempty"`
	// Write buffer size as bytes, 0 means no buffering
	WriteBufferSize int `json:"write_buffer_size,omitempty"`
}

// BaseS3FsConfig defines the base configuration for S3 based filesystems
type BaseS3FsConfig struct {
	Bucket string `json:"bucket,omitempty"`
	// KeyPrefix is similar to a chroot directory for local filesystem.
	// If specified then the SFTP user will only see objects that starts
	// with this prefix and so you can restrict access to a specific
	// folder. The prefix, if not empty, must not start with "/" and must
	// end with "/".
	// If empty the whole bucket contents will be available
	KeyPrefix string `json:"key_prefix,omitempty"`
	Region    string `json:"region,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
	// IAM Role ARN to assume
	RoleARN string `json:"role_arn,omitempty"`
	// Optional Session token that is a part of temporary security credentials provisioned by AWS STS.
	SessionToken string `json:"session_token,omitempty"`
	Endpoint     string `json:"endpoint,omitempty"`
	StorageClass string `json:"storage_class,omitempty"`
	// The canned ACL to apply to uploaded objects. Leave empty to use the default ACL.
	// For more information and available ACLs, see here:
	// https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	ACL string `json:"acl,omitempty"`
	// The buffer size (in MB) to use for multipart uploads. The minimum allowed part size is 5MB,
	// and if this value is set to zero, the default value (5MB) for the AWS SDK will be used.
	// The minimum allowed value is 5.
	// Please note that if the upload bandwidth between the SFTP client and SFTPGo is greater than
	// the upload bandwidth between SFTPGo and S3 then the SFTP client have to wait for the upload
	// of the last parts to S3 after it ends the file upload to SFTPGo, and it may time out.
	// Keep this in mind if you customize these parameters.
	UploadPartSize int64 `json:"upload_part_size,omitempty"`
	// How many parts are uploaded in parallel. 0 means the default (5)
	UploadConcurrency int `json:"upload_concurrency,omitempty"`
	// The buffer size (in MB) to use for multipart downloads. The minimum allowed part size is 5MB,
	// and if this value is set to zero, the default value (5MB) for the AWS SDK will be used.
	// The minimum allowed value is 5. Ignored for partial downloads.
	DownloadPartSize int64 `json:"download_part_size,omitempty"`
	// UploadPartMaxTime defines the maximum time allowed, in seconds, to upload a single chunk.
	// 0 means no timeout.
	UploadPartMaxTime int `json:"upload_part_max_time,omitempty"`
	// How many parts are downloaded in parallel.  0 means the default (5). Ignored for partial downloads.
	DownloadConcurrency int `json:"download_concurrency,omitempty"`
	// DownloadPartMaxTime defines the maximum time allowed, in seconds, to download a single chunk.
	// 0 means no timeout. Ignored for partial downloads.
	DownloadPartMaxTime int `json:"download_part_max_time,omitempty"`
	// Set this to `true` to force the request to use path-style addressing,
	// i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client
	// will use virtual hosted bucket addressing when possible
	// (`http://BUCKET.s3.amazonaws.com/KEY`)
	ForcePathStyle bool `json:"force_path_style,omitempty"`
	// If enabled the S3 client accepts any TLS certificate presented by
	// the server and any host name in that certificate.
	// In this mode, TLS is susceptible to man-in-the-middle attacks.
	// This should be used only for testing.
	SkipTLSVerify bool `json:"skip_tls_verify,omitempty"`
}

// S3FsConfig defines the base configuration for S3 based filesystems
type S3FsConfig struct {
	BaseS3FsConfig
	AccessSecret   kms.BaseSecret `json:"access_secret,omitempty"`
	SSECustomerKey kms.BaseSecret `json:"sse_customer_key,omitempty"`
}

// BaseGCSFsConfig defines the base configuration for Google Cloud Storage based filesystems
type BaseGCSFsConfig struct {
	Bucket string `json:"bucket,omitempty"`
	// KeyPrefix is similar to a chroot directory for local filesystem.
	// If specified then the SFTP user will only see objects that starts
	// with this prefix and so you can restrict access to a specific
	// folder. The prefix, if not empty, must not start with "/" and must
	// end with "/".
	// If empty the whole bucket contents will be available
	KeyPrefix      string `json:"key_prefix,omitempty"`
	CredentialFile string `json:"-"`
	// 0 explicit, 1 automatic
	AutomaticCredentials int    `json:"automatic_credentials,omitempty"`
	StorageClass         string `json:"storage_class,omitempty"`
	// The ACL to apply to uploaded objects. Leave empty to use the default ACL.
	// For more information and available ACLs, refer to the JSON API here:
	// https://cloud.google.com/storage/docs/access-control/lists#predefined-acl
	ACL string `json:"acl,omitempty"`
	// The buffer size (in MB) to use for multipart uploads. The default value is 16MB.
	// 0 means use the default
	UploadPartSize int64 `json:"upload_part_size,omitempty"`
	// UploadPartMaxTime defines the maximum time allowed, in seconds, to upload a single chunk.
	// The default value is 32. 0 means use the default
	UploadPartMaxTime int `json:"upload_part_max_time,omitempty"`
}

// GCSFsConfig defines the configuration for Google Cloud Storage based filesystems
type GCSFsConfig struct {
	BaseGCSFsConfig
	Credentials kms.BaseSecret `json:"credentials,omitempty"`
}

// BaseAzBlobFsConfig defines the base configuration for Azure Blob Storage based filesystem
type BaseAzBlobFsConfig struct {
	Container string `json:"container,omitempty"`
	// Storage Account Name, leave blank to use SAS URL
	AccountName string `json:"account_name,omitempty"`
	// Optional endpoint. Default is "blob.core.windows.net".
	// If you use the emulator the endpoint must include the protocol,
	// for example "http://127.0.0.1:10000"
	Endpoint string `json:"endpoint,omitempty"`
	// KeyPrefix is similar to a chroot directory for local filesystem.
	// If specified then the SFTPGo user will only see objects that starts
	// with this prefix and so you can restrict access to a specific
	// folder. The prefix, if not empty, must not start with "/" and must
	// end with "/".
	// If empty the whole bucket contents will be available
	KeyPrefix string `json:"key_prefix,omitempty"`
	// The buffer size (in MB) to use for multipart uploads.
	// If this value is set to zero, the default value (5MB) will be used.
	// Please note that if the upload bandwidth between the SFTPGo client and SFTPGo server is
	// greater than the upload bandwidth between SFTPGo and Azure then the SFTP client have
	// to wait for the upload of the last parts to Azure after it ends the file upload to SFTPGo,
	// and it may time out.
	// Keep this in mind if you customize these parameters.
	UploadPartSize int64 `json:"upload_part_size,omitempty"`
	// How many parts are uploaded in parallel. Default: 5
	UploadConcurrency int `json:"upload_concurrency,omitempty"`
	// The buffer size (in MB) to use for multipart downloads.
	// If this value is set to zero, the default value (5MB) will be used.
	DownloadPartSize int64 `json:"download_part_size,omitempty"`
	// How many parts are downloaded in parallel. 0 means the default (5)
	DownloadConcurrency int `json:"download_concurrency,omitempty"`
	// Set to true if you use an Azure emulator such as Azurite
	UseEmulator bool `json:"use_emulator,omitempty"`
	// Blob Access Tier
	AccessTier string `json:"access_tier,omitempty"`
}

// AzBlobFsConfig defines the configuration for Azure Blob Storage based filesystem
type AzBlobFsConfig struct {
	BaseAzBlobFsConfig
	// Storage Account Key leave blank to use SAS URL.
	// The access key is stored encrypted based on the kms configuration
	AccountKey kms.BaseSecret `json:"account_key,omitempty"`
	// Shared access signature URL, leave blank if using account/key
	SASURL kms.BaseSecret `json:"sas_url,omitempty"`
}

// CryptFsConfig defines the configuration to store local files as encrypted
type CryptFsConfig struct {
	OSFsConfig
	Passphrase kms.BaseSecret `json:"passphrase,omitempty"`
}

// BaseSFTPFsConfig defines the base configuration for SFTP based filesystem
type BaseSFTPFsConfig struct {
	Endpoint     string   `json:"endpoint,omitempty"`
	Username     string   `json:"username,omitempty"`
	Fingerprints []string `json:"fingerprints,omitempty"`
	// Prefix is the path prefix to strip from SFTP resource paths.
	Prefix string `json:"prefix,omitempty"`
	// Concurrent reads are safe to use and disabling them will degrade performance.
	// Some servers automatically delete files once they are downloaded.
	// Using concurrent reads is problematic with such servers.
	DisableCouncurrentReads bool `json:"disable_concurrent_reads,omitempty"`
	// The buffer size (in MB) to use for transfers.
	// Buffering could improve performance for high latency networks.
	// With buffering enabled upload resume is not supported and a file
	// cannot be opened for both reading and writing at the same time
	// 0 means disabled.
	BufferSize int64 `json:"buffer_size,omitempty"`
	// Defines how to check if this config points to the same
	// server as another config. By default both the endpoint and
	// the username must match. 1 means that only the endpoint must match.
	// If different configs point to the same server the renaming
	// between the fs configs is allowed.
	EqualityCheckMode int `json:"equality_check_mode,omitempty"`
}

// SFTPFsConfig defines the configuration for SFTP based filesystem
type SFTPFsConfig struct {
	BaseSFTPFsConfig
	Password      kms.BaseSecret `json:"password,omitempty"`
	PrivateKey    kms.BaseSecret `json:"private_key,omitempty"`
	KeyPassphrase kms.BaseSecret `json:"key_passphrase,omitempty"`
}

// BaseHTTPFsConfig defines the base configuration for HTTP based filesystem
type BaseHTTPFsConfig struct {
	// HTTP/S endpoint URL. SFTPGo will use this URL as base, for example for the
	// "stat" API, SFTPGo will add "/stat/{name}"
	Endpoint string `json:"endpoint,omitempty"`
	Username string `json:"username,omitempty"`
	// if enabled the HTTP client accepts any TLS certificate presented by
	// the server and any host name in that certificate.
	// In this mode, TLS is susceptible to man-in-the-middle attacks.
	// This should be used only for testing.
	SkipTLSVerify bool `json:"skip_tls_verify,omitempty"`
	// Defines how to check if this config points to the same
	// server as another config. By default both the endpoint and
	// the username must match. 1 means that only the endpoint must match.
	// If different configs point to the same server the renaming
	// between the fs configs is allowed.
	EqualityCheckMode int `json:"equality_check_mode,omitempty"`
}

// HTTPFsConfig defines the configuration for HTTP based filesystem
type HTTPFsConfig struct {
	BaseHTTPFsConfig
	Password kms.BaseSecret `json:"password,omitempty"`
	APIKey   kms.BaseSecret `json:"api_key,omitempty"`
}

// Filesystem defines filesystem details
type Filesystem struct {
	Provider     FilesystemProvider `json:"provider"`
	OSConfig     OSFsConfig         `json:"osconfig,omitempty"`
	S3Config     S3FsConfig         `json:"s3config,omitempty"`
	GCSConfig    GCSFsConfig        `json:"gcsconfig,omitempty"`
	AzBlobConfig AzBlobFsConfig     `json:"azblobconfig,omitempty"`
	CryptConfig  CryptFsConfig      `json:"cryptconfig,omitempty"`
	SFTPConfig   SFTPFsConfig       `json:"sftpconfig,omitempty"`
	HTTPConfig   HTTPFsConfig       `json:"httpconfig,omitempty"`
}
