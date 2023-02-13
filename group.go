package sdk

// Group types
const (
	// Primary group
	GroupTypePrimary = iota + 1
	// Secondary group
	GroupTypeSecondary
	// Membership only, no settings are inherited from this group type
	GroupTypeMembership
)

// BaseGroupUserSettings defines the base settings to apply to users
type BaseGroupUserSettings struct {
	HomeDir string `json:"home_dir"`
	// Maximum concurrent sessions. 0 means unlimited
	MaxSessions int `json:"max_sessions"`
	// Maximum size allowed as bytes. 0 means unlimited
	QuotaSize int64 `json:"quota_size"`
	// Maximum number of files allowed. 0 means unlimited
	QuotaFiles int `json:"quota_files"`
	// List of permissions granted per-directory
	Permissions map[string][]string `json:"permissions"`
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
	// Defines account expiration in number of days from creation.
	// 0 means no expiration
	ExpiresIn int `json:"expires_in,omitempty"`
	// Additional restrictions
	Filters BaseUserFilters `json:"filters"`
}

// GroupUserSettings defines the settings to apply to users
type GroupUserSettings struct {
	BaseGroupUserSettings
	// Filesystem configuration details
	FsConfig Filesystem `json:"filesystem"`
}

// BaseGroup defines the shared group fields
type BaseGroup struct {
	// Data provider unique identifier
	ID int64 `json:"id"`
	// Group name
	Name string `json:"name"`
	// optional description
	Description string `json:"description,omitempty"`
	// Creation time as unix timestamp in milliseconds
	CreatedAt int64 `json:"created_at"`
	// last update time as unix timestamp in milliseconds
	UpdatedAt int64 `json:"updated_at"`
	// list of usernames associated with this group
	Users []string `json:"users,omitempty"`
	// list of admins associated with this group
	Admins []string `json:"admins,omitempty"`
}

// Group defines an SFTPGo group.
// Groups are used to easily configure similar users
type Group struct {
	BaseGroup
	// settings to apply to users for whom this is a primary group
	UserSettings GroupUserSettings `json:"user_settings,omitempty"`
	// Mapping between virtual paths and virtual folders
	VirtualFolders []VirtualFolder `json:"virtual_folders,omitempty"`
}

// GroupMapping defines the mapping between an SFTPGo user and a group
type GroupMapping struct {
	Name string `json:"name"` // group name
	Type int    `json:"type"`
}
