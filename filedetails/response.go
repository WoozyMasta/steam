package filedetails

import (
	"time"

	json "github.com/json-iterator/go"
)

// FileDetail represents detailed information about a file (mod) published in the Steam Workshop.
type FileDetail struct {
	TimeCreated                time.Time   `json:"time_created"`                                // The timestamp when the file was created.
	TimeUpdated                time.Time   `json:"time_updated"`                                // The timestamp when the file was last updated.
	AppName                    string      `json:"app_name,omitempty"`                          // The name of the application associated with the file.
	BanReason                  string      `json:"ban_reason,omitempty"`                        // The reason why the file was banned (if applicable).
	FileDescription            string      `json:"file_description,omitempty"`                  // A textual description of the file.
	Filename                   string      `json:"filename,omitempty"`                          // The name of the file.
	PreviewURL                 string      `json:"preview_url,omitempty"`                       // The URL for the file's preview image.
	Title                      string      `json:"title"`                                       // The title of the file.
	URL                        string      `json:"url,omitempty"`                               // The URL to access the file in the Steam Workshop.
	YoutubeVideoID             string      `json:"youtubevideoid,omitempty"`                    // YouTube video ID for the preview (if applicable).}
	Children                   []Children  `json:"children,omitempty"`                          // A list of child files associated with this file.
	KVTags                     []KVTags    `json:"kvtags,omitempty"`                            // Key-value tags for categorization or metadata.
	Previews                   []Previews  `json:"previews,omitempty"`                          // A list of preview images or videos.
	Reactions                  []Reactions `json:"reactions,omitempty"`                         // User reactions or feedback for the file.
	Tags                       []Tags      `json:"tags,omitempty"`                              // Tags applied to the file.
	VoteData                   []VoteData  `json:"vote_data,omitempty"`                         // Voting statistics for the file.
	Banner                     uint64      `json:"banner,string,omitempty"`                     // Banner ID or reference.
	BanTextCheckResult         int         `json:"ban_text_check_result,omitempty"`             // Result of the text check for banning.
	ConsumerAppID              uint64      `json:"consumer_appid,omitempty"`                    // App ID of the consumer.
	ConsumerShortcutID         int         `json:"consumer_shortcutid,omitempty"`               // Shortcut ID associated with the consumer.
	Creator                    uint64      `json:"creator,string,omitempty"`                    // ID of the file's creator.
	CreatorAppID               uint64      `json:"creator_appid,omitempty"`                     // App ID of the creator.
	Favorited                  int         `json:"favorited,omitempty"`                         // Number of times the file has been favorited.
	FileSize                   uint64      `json:"file_size,string,omitempty"`                  // The size of the file in bytes.
	FileType                   int         `json:"file_type,omitempty"`                         // The type of file.
	Flags                      int         `json:"flags,omitempty"`                             // File flags indicating special properties.
	Followers                  int         `json:"followers,omitempty"`                         // Number of followers for the file.
	HContentFile               uint64      `json:"hcontent_file,string,omitempty"`              // Content hash or reference ID for the file.
	HContentPreview            uint64      `json:"hcontent_preview,string,omitempty"`           // Content hash or reference ID for the file's preview.
	ImageHeight                int         `json:"image_height,omitempty"`                      // Image height in pixels
	ImageWidth                 int         `json:"image_width,omitempty"`                       // Image width in pixels
	Language                   int         `json:"language,omitempty"`                          // The language of the file.
	LifetimeFavorited          int         `json:"lifetime_favorited,omitempty"`                // Total number of times the file has been favorited.
	LifetimeFollowers          int         `json:"lifetime_followers,omitempty"`                // Total number of followers for the file.
	LifetimePlaytime           uint64      `json:"lifetime_playtime,string,omitempty"`          // Total playtime across all users.
	LifetimePlaytimeSessions   uint64      `json:"lifetime_playtime_sessions,string,omitempty"` // Total playtime sessions across all users.
	LifetimeSubscriptions      int         `json:"lifetime_subscriptions,omitempty"`            // Total number of subscriptions for the file.
	NumChildren                int         `json:"num_children,omitempty"`                      // Number of child files associated with this file.
	NumCommentsPublic          int         `json:"num_comments_public,omitempty"`               // Number of public comments on the file.
	NumReports                 int         `json:"num_reports,omitempty"`                       // Number of reports submitted for the file.
	PreviewFileSize            uint64      `json:"preview_file_size,string,omitempty"`          // The size of the preview file in bytes.
	PublishedFileID            uint64      `json:"publishedfileid,string"`                      // The unique ID of the published file.
	Result                     int         `json:"result,omitempty"`                            // The result code for the operation.
	Revision                   int         `json:"revision,omitempty"`                          // The revision number of the file.
	RevisionChangeNumber       uint64      `json:"revision_change_number,string,omitempty"`     // The revision change number for tracking updates.
	Subscriptions              int         `json:"subscriptions,omitempty"`                     // Number of subscriptions to the file.
	Views                      int         `json:"views,omitempty"`                             // Number of views for the file.
	Visibility                 int         `json:"visibility,omitempty"`                        // The visibility level of the file.
	Banned                     bool        `json:"banned,omitempty"`                            // Indicates if the file is banned.
	CanBeDeleted               bool        `json:"can_be_deleted,omitempty"`                    // Indicates if the file can be deleted.
	CanSubscribe               bool        `json:"can_subscribe,omitempty"`                     // Indicates if users can subscribe to the file.
	MaybeInappropriateSex      bool        `json:"maybe_inappropriate_sex,omitempty"`           // Indicates if the file may contain inappropriate sexual content.
	MaybeInappropriateViolence bool        `json:"maybe_inappropriate_violence,omitempty"`      // Indicates if the file may contain inappropriate violent content.
	ShowSubscribeAll           bool        `json:"show_subscribe_all,omitempty"`                // Indicates if "subscribe to all" is available for the file.
	WorkshopAccepted           bool        `json:"workshop_accepted,omitempty"`                 // Indicates if the file was accepted in the workshop.
	WorkshopFile               bool        `json:"workshop_file"`                               // Indicates if the file is a workshop file.
}

// Custom implementation for deserializing UnixTime in `FileDetail`
func (fd *FileDetail) UnmarshalJSON(data []byte) error {
	type Alias FileDetail

	aux := &struct {
		*Alias
		TimeCreated int64 `json:"time_created"`
		TimeUpdated int64 `json:"time_updated"`
	}{
		Alias: (*Alias)(fd),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	fd.TimeCreated = time.Unix(aux.TimeCreated, 0)
	fd.TimeUpdated = time.Unix(aux.TimeUpdated, 0)

	return nil
}

// Children represents a child file associated with a parent file in the workshop.
type Children struct {
	PublishedFileID uint64 `json:"publishedfileid,string"` // The unique ID of the child file.
	FileType        int    `json:"file_type"`              // The type of the child file.
	SortOrder       int    `json:"sortorder"`              // The order in which the child file appears.
}

// KVTags represents a key-value tag associated with a file.
type KVTags struct {
	Key   string `json:"key"`   // The key of the tag.
	Value string `json:"value"` // The value of the tag.
}

// Previews represents preview data for a file, such as images or videos.
type Previews struct {
	URL               string `json:"url,omitempty"`                // The URL of the preview.
	Filename          string `json:"filename,omitempty"`           // The filename of the preview.
	ExternalReference string `json:"external_reference,omitempty"` // External reference link for the preview.
	YoutubeVideoID    string `json:"youtubevideoid,omitempty"`     // YouTube video ID for the preview (if applicable).
	PreviewID         uint64 `json:"previewid,string,omitempty"`   // The unique ID of the preview.
	Size              int    `json:"size,omitempty"`               // The size of the preview file in bytes.
	PreviewType       int    `json:"preview_type,omitempty"`       // The type of preview (e.g., image, video).
	SortOrder         int    `json:"sortorder,omitempty"`          // The order in which the preview appears.
}

// Reactions represents user reactions to a file.
type Reactions struct {
	Count      int `json:"count"`      // Number of reactions of this type.
	ReactionID int `json:"reactionid"` // The ID of the reaction type.
}

// Tags represents a tag applied to a file.
type Tags struct {
	DisplayName string `json:"display_name"` // The display name of the tag.
	Tag         string `json:"tag"`          // The internal value of the tag.
}

// VoteData represents voting data for a file.
type VoteData struct {
	Score     float64 `json:"score"`      // The average score of the votes.
	VotesDown int     `json:"votes_down"` // Number of downvotes for the file.
	VotesUp   int     `json:"votes_up"`   // Number of upvotes for the file.
}
