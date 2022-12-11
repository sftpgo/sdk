package sdk

import (
	"path"
	"strings"

	"github.com/sftpgo/sdk/kms"
)

// Web Client/user REST API restrictions
const (
	WebClientPubKeyChangeDisabled     = "publickey-change-disabled"
	WebClientWriteDisabled            = "write-disabled"
	WebClientMFADisabled              = "mfa-disabled"
	WebClientPasswordChangeDisabled   = "password-change-disabled"
	WebClientAPIKeyAuthChangeDisabled = "api-key-auth-change-disabled"
	WebClientInfoChangeDisabled       = "info-change-disabled"
	WebClientSharesDisabled           = "shares-disabled"
	WebClientPasswordResetDisabled    = "password-reset-disabled"
	WebClientShareNoPasswordDisabled  = "shares-without-password-disabled"
)

const (
	// DenyPolicyDefault means that denied files matching the filters are visible in directory
	// listing but cannot be uploaded/downloaded/overwritten/renamed
	DenyPolicyDefault = iota
	// DenyPolicyHide applies the same restrictions as DenyPolicyDefault and denied files/directories
	// matching the filters will also be hidden in directory listing.
	// This mode may cause performance issues for large directories
	DenyPolicyHide
)

var (
	// WebClientOptions defines the available options for the web client interface/user REST API
	WebClientOptions = []string{WebClientWriteDisabled, WebClientPasswordChangeDisabled, WebClientPasswordResetDisabled,
		WebClientPubKeyChangeDisabled, WebClientMFADisabled, WebClientAPIKeyAuthChangeDisabled, WebClientInfoChangeDisabled,
		WebClientSharesDisabled, WebClientShareNoPasswordDisabled}
	// UserTypes defines the supported user type hints for auth plugins
	UserTypes = []string{string(UserTypeLDAP), string(UserTypeOS)}
)

// TLSUsername defines the TLS certificate attribute to use as username
type TLSUsername string

// Supported certificate attributes to use as username
const (
	TLSUsernameNone TLSUsername = "None"
	TLSUsernameCN   TLSUsername = "CommonName"
)

// UserType defines the supported user types.
// This is an hint for external auth plugins, is not used in SFTPGo directly
type UserType string

// User types, auth plugins could use this info to choose the correct authentication backend
const (
	UserTypeLDAP UserType = "LDAPUser"
	UserTypeOS   UserType = "OSUser"
)

// DirectoryPermissions defines permissions for a directory virtual path
type DirectoryPermissions struct {
	Path        string
	Permissions []string
}

// PatternsFilter defines filters based on shell like patterns.
// System commands such as Git and rsync interacts with the filesystem directly
// and they are not aware about these restrictions so they are not allowed
// inside paths with extensions filters
type PatternsFilter struct {
	// Virtual path, if no other specific filter is defined, the filter applies for
	// sub directories too.
	// For example if filters are defined for the paths "/" and "/sub" then the
	// filters for "/" are applied for any file outside the "/sub" directory
	Path string `json:"path"`
	// files/dir with these, case insensitive, patterns are allowed.
	// Denied file patterns are evaluated before the allowed ones
	AllowedPatterns []string `json:"allowed_patterns,omitempty"`
	// files/dir with these, case insensitive, patterns are not allowed.
	// Denied file patterns are evaluated before the allowed ones
	DeniedPatterns []string `json:"denied_patterns,omitempty"`
	// Deny policy
	DenyPolicy int `json:"deny_policy,omitempty"`
}

// GetCommaSeparatedPatterns returns the first non empty patterns list comma separated
func (p *PatternsFilter) GetCommaSeparatedPatterns() string {
	if len(p.DeniedPatterns) > 0 {
		return strings.Join(p.DeniedPatterns, ",")
	}
	return strings.Join(p.AllowedPatterns, ",")
}

// IsDenied returns true if the patterns has one or more denied patterns
func (p *PatternsFilter) IsDenied() bool {
	return len(p.DeniedPatterns) > 0
}

// IsAllowed returns true if the patterns has one or more allowed patterns
func (p *PatternsFilter) IsAllowed() bool {
	return len(p.AllowedPatterns) > 0
}

// CheckAllowed returns true if the specified item is allowed
func (p *PatternsFilter) CheckAllowed(item string) bool {
	if p.Path != "" {
		toMatch := strings.ToLower(item)
		for _, denied := range p.DeniedPatterns {
			matched, err := path.Match(denied, toMatch)
			if err != nil || matched {
				return false
			}
		}
		for _, allowed := range p.AllowedPatterns {
			matched, err := path.Match(allowed, toMatch)
			if err == nil && matched {
				return true
			}
		}
		return len(p.AllowedPatterns) == 0
	}
	return true
}

// HooksFilter defines user specific overrides for global hooks
type HooksFilter struct {
	ExternalAuthDisabled  bool `json:"external_auth_disabled"`
	PreLoginDisabled      bool `json:"pre_login_disabled"`
	CheckPasswordDisabled bool `json:"check_password_disabled"`
}

// RecoveryCode defines a 2FA recovery code
type RecoveryCode struct {
	Secret kms.BaseSecret `json:"secret"`
	Used   bool           `json:"used,omitempty"`
}

// TOTPConfig defines the time-based one time password configuration
type TOTPConfig struct {
	Enabled    bool           `json:"enabled,omitempty"`
	ConfigName string         `json:"config_name,omitempty"`
	Secret     kms.BaseSecret `json:"secret,omitempty"`
	// TOTP will be required for the specified protocols.
	// SSH protocol (SFTP/SCP/SSH commands) will ask for the TOTP passcode if the client uses keyboard interactive
	// authentication.
	// FTP have no standard way to support two factor authentication, if you
	// enable the support for this protocol you have to add the TOTP passcode after the password.
	// For example if your password is "password" and your one time passcode is
	// "123456" you have to use "password123456" as password.
	Protocols []string `json:"protocols,omitempty"`
}

// BandwidthLimit defines a per-source bandwidth limit
type BandwidthLimit struct {
	// Source networks in CIDR notation as defined in RFC 4632 and RFC 4291
	// for example "192.0.2.0/24" or "2001:db8::/32". The limit applies if the
	// defined networks contain the client IP
	Sources []string `json:"sources"`
	// Maximum upload bandwidth as KB/s
	UploadBandwidth int64 `json:"upload_bandwidth,omitempty"`
	// Maximum download bandwidth as KB/s
	DownloadBandwidth int64 `json:"download_bandwidth,omitempty"`
}

// GetSourcesAsString returns the sources as comma separated string
func (l *BandwidthLimit) GetSourcesAsString() string {
	return strings.Join(l.Sources, ",")
}

// DataTransferLimit defines a per-source data transfer limits
type DataTransferLimit struct {
	// Source networks in CIDR notation as defined in RFC 4632 and RFC 4291
	// for example "192.0.2.0/24" or "2001:db8::/32". The limit applies if the
	// defined networks contain the client IP
	Sources []string `json:"sources"`
	// Maximum data transfer allowed for uploads as MB. 0 means no limit.
	UploadDataTransfer int64 `json:"upload_data_transfer"`
	// Maximum data transfer allowed for downloads as MB. 0 means no limit.
	DownloadDataTransfer int64 `json:"download_data_transfer"`
	// Maximum total data transfer as MB. 0 means unlimited.
	// You can set a total data transfer instead of the individual values
	// for uploads and downloads
	TotalDataTransfer int64 `json:"total_data_transfer"`
}

// GetSourcesAsString returns the sources as comma separated string
func (l *DataTransferLimit) GetSourcesAsString() string {
	return strings.Join(l.Sources, ",")
}

// BaseUserFilters defines additional restrictions for a user
type BaseUserFilters struct {
	// only clients connecting from these IP/Mask are allowed.
	// IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291
	// for example "192.0.2.0/24" or "2001:db8::/32"
	AllowedIP []string `json:"allowed_ip,omitempty"`
	// clients connecting from these IP/Mask are not allowed.
	// Denied rules will be evaluated before allowed ones
	DeniedIP []string `json:"denied_ip,omitempty"`
	// these login methods are not allowed.
	// If null or empty any available login method is allowed
	DeniedLoginMethods []string `json:"denied_login_methods,omitempty"`
	// these protocols are not allowed.
	// If null or empty any available protocol is allowed
	DeniedProtocols []string `json:"denied_protocols,omitempty"`
	// filter based on shell patterns.
	// Please note that these restrictions can be easily bypassed.
	FilePatterns []PatternsFilter `json:"file_patterns,omitempty"`
	// max size allowed for a single upload, 0 means unlimited
	MaxUploadFileSize int64 `json:"max_upload_file_size,omitempty"`
	// TLS certificate attribute to use as username.
	// For FTP clients it must match the name provided using the
	// "USER" command
	TLSUsername TLSUsername `json:"tls_username,omitempty"`
	// user specific hook overrides
	Hooks HooksFilter `json:"hooks,omitempty"`
	// Disable checks for existence and automatic creation of home directory
	// and virtual folders.
	// SFTPGo requires that the user's home directory, virtual folder root,
	// and intermediate paths to virtual folders exist to work properly.
	// If you already know that the required directories exist, disabling
	// these checks will speed up login.
	// You could, for example, disable these checks after the first login
	DisableFsChecks bool `json:"disable_fs_checks,omitempty"`
	// WebClient related configuration options
	WebClient []string `json:"web_client,omitempty"`
	// API key auth allows to impersonate this user with an API key
	AllowAPIKeyAuth bool `json:"allow_api_key_auth,omitempty"`
	// UserType is an hint for authentication plugins.
	// It is ignored when using SFTPGo internal authentication
	UserType string `json:"user_type,omitempty"`
	// Per-source bandwidth limits
	BandwidthLimits []BandwidthLimit `json:"bandwidth_limits,omitempty"`
	// Per-source data transfer limits
	DataTransferLimits []DataTransferLimit `json:"data_transfer_limits,omitempty"`
	// Defines the cache time, in seconds, for users authenticated using
	// an external auth hook. 0 means no cache
	ExternalAuthCacheTime int64 `json:"external_auth_cache_time,omitempty"`
	// Specifies an alternate starting directory. If not set, the default is "/".
	// This option is supported for SFTP/SCP, FTP and HTTP (WebClient/REST API) protocols.
	// Relative paths will use this directory as base
	StartDirectory string `json:"start_directory,omitempty"`
	// TwoFactorAuthProtocols defines protocols that require two factor authentication
	TwoFactorAuthProtocols []string `json:"2fa_protocols,omitempty"`
	// Define the FTP security mode. Set to 1 to require TLS for both data and control
	// connection. This setting is useful if you want to allow both encrypted and plain text
	// FTP sessions globally and then you want to require encrypted sessions on a per-user
	// basis.
	// It has no effect if TLS is already required for all users in the configuration file.
	FTPSecurity int `json:"ftp_security,omitempty"`
	// If enabled the user can login with any password or no password at all.
	// Anonymous users are supported for FTP and WebDAV protocols and
	// permissions will be automatically set to "list" and "download" (read only)
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Defines the default expiration for newly created shares as number of days.
	// 0 means no expiration
	DefaultSharesExpiration int `json:"default_shares_expiration,omitempty"`
	// The password expires after the defined number of days. 0 means no expiration
	PasswordExpiration int `json:"password_expiration,omitempty"`
}

// GetFlatFilePatterns returns file patterns as flat list
// duplicating a path if it has both allowed and denied patterns
func (f *BaseUserFilters) GetFlatFilePatterns() []PatternsFilter {
	result := make([]PatternsFilter, 0, len(f.FilePatterns))

	for _, pattern := range f.FilePatterns {
		if len(pattern.AllowedPatterns) > 0 {
			result = append(result, PatternsFilter{
				Path:            pattern.Path,
				AllowedPatterns: pattern.AllowedPatterns,
				DenyPolicy:      pattern.DenyPolicy,
			})
		}
		if len(pattern.DeniedPatterns) > 0 {
			result = append(result, PatternsFilter{
				Path:           pattern.Path,
				DeniedPatterns: pattern.DeniedPatterns,
				DenyPolicy:     pattern.DenyPolicy,
			})
		}
	}
	return result
}

// UserFilters defines additional restrictions for a user
// TODO: rename to UserOptions in v3
type UserFilters struct {
	BaseUserFilters
	// User must change password from WebClient/REST API at next login.
	RequirePasswordChange bool `json:"require_password_change,omitempty"`
	// Time-based one time passwords configuration
	TOTPConfig TOTPConfig `json:"totp_config,omitempty"`
	// Recovery codes to use if the user loses access to their second factor auth device.
	// Each code can only be used once, you should use these codes to login and disable or
	// reset 2FA for your account
	RecoveryCodes []RecoveryCode `json:"recovery_codes,omitempty"`
}

// BaseUser defines the shared user fields
type BaseUser struct {
	// Data provider unique identifier
	ID int64 `json:"id"`
	// 1 enabled, 0 disabled (login is not allowed)
	Status int `json:"status"`
	// Username
	Username string `json:"username"`
	// Email
	Email string `json:"email,omitempty"`
	// Account expiration date as unix timestamp in milliseconds. An expired account cannot login.
	// 0 means no expiration
	ExpirationDate int64 `json:"expiration_date,omitempty"`
	// Password used for password authentication.
	// For users created using SFTPGo REST API the password is be stored using bcrypt or argon2id hashing algo.
	// Checking passwords stored with pbkdf2, md5crypt and sha512crypt is supported too.
	Password string `json:"password,omitempty"`
	// PublicKeys used for public key authentication. At least one between password and a public key is mandatory
	PublicKeys []string `json:"public_keys,omitempty"`
	// The user cannot upload or download files outside this directory. Must be an absolute path
	HomeDir string `json:"home_dir"`
	// If sftpgo runs as root system user then the created files and directories will be assigned to this system UID
	UID int `json:"uid"`
	// If sftpgo runs as root system user then the created files and directories will be assigned to this system GID
	GID int `json:"gid"`
	// Maximum concurrent sessions. 0 means unlimited
	MaxSessions int `json:"max_sessions"`
	// Maximum size allowed as bytes. 0 means unlimited
	QuotaSize int64 `json:"quota_size"`
	// Maximum number of files allowed. 0 means unlimited
	QuotaFiles int `json:"quota_files"`
	// List of permissions granted per-directory
	Permissions map[string][]string `json:"permissions"`
	// Used quota as bytes
	UsedQuotaSize int64 `json:"used_quota_size,omitempty"`
	// Used quota as number of files
	UsedQuotaFiles int `json:"used_quota_files,omitempty"`
	// Last quota update as unix timestamp in milliseconds
	LastQuotaUpdate int64 `json:"last_quota_update,omitempty"`
	// Maximum upload bandwidth as KB/s, 0 means unlimited.
	// This is the default if no per-source limit match
	UploadBandwidth int64 `json:"upload_bandwidth,omitempty"`
	// Maximum download bandwidth as KB/s, 0 means unlimited.
	// This is the default if no per-source limit match
	DownloadBandwidth int64 `json:"download_bandwidth,omitempty"`
	// Maximum data transfer allowed for uploads as MB. 0 means no limit.
	// You can periodically reset the data related transfer fields for example
	// each month
	UploadDataTransfer int64 `json:"upload_data_transfer"`
	// Maximum data transfer allowed for downloads as MB. 0 means no limit.
	DownloadDataTransfer int64 `json:"download_data_transfer"`
	// Maximum total data transfer as MB. 0 means unlimited.
	// You can set a total data transfer instead of the individual values
	// for uploads and downloads
	TotalDataTransfer int64 `json:"total_data_transfer"`
	// Uploaded size, as bytes, since the last reset
	UsedUploadDataTransfer int64 `json:"used_upload_data_transfer,omitempty"`
	// Downloaded size, as bytes, since the last reset
	UsedDownloadDataTransfer int64 `json:"used_download_data_transfer,omitempty"`
	// Last login as unix timestamp in milliseconds
	LastLogin int64 `json:"last_login,omitempty"`
	// Creation time as unix timestamp in milliseconds. It will be 0 for admins created before v2.2.0
	CreatedAt int64 `json:"created_at"`
	// last update time as unix timestamp in milliseconds
	UpdatedAt int64 `json:"updated_at"`
	// first download time as unix timestamp in milliseconds
	FirstDownload int64 `json:"first_download,omitempty"`
	// first upload time as unix timestamp in milliseconds
	FirstUpload int64 `json:"first_upload,omitempty"`
	// last password change as unix timestamp in milliseconds
	LastPasswordChange int64 `json:"last_password_change,omitempty"`
	// optional description, for example full name
	Description string `json:"description,omitempty"`
	// free form text field for external systems
	AdditionalInfo string `json:"additional_info,omitempty"`
	// groups associated with this user
	Groups []GroupMapping `json:"groups,omitempty"`
	// This field is passed to the pre-login hook if custom OIDC fields have been configured.
	// Field values can be of any type (this is a free form object) and depend on the type
	// of the configured OIDC fields.
	// This fields are never saved or returned in anything other than the pre-login hook
	OIDCCustomFields *map[string]interface{} `json:"oidc_custom_fields,omitempty"`
	// Role name
	Role string `json:"role,omitempty"`
}

// User defines a SFTPGo user
type User struct {
	BaseUser
	// Additional restrictions
	Filters UserFilters `json:"filters"`
	// Mapping between virtual paths and virtual folders
	VirtualFolders []VirtualFolder `json:"virtual_folders,omitempty"`
	// Filesystem configuration details
	FsConfig Filesystem `json:"filesystem"`
}
