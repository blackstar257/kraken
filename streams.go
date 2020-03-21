package kraken

import "time"

// Stream ...
type Stream struct {
	ID                int           `json:"_id"`
	Game              string        `json:"game"`
	BroadcastPlatform string        `json:"broadcast_platform"`
	CommunityID       string        `json:"community_id"`
	CommunityIds      []interface{} `json:"community_ids"`
	Viewers           int           `json:"viewers"`
	VideoHeight       int           `json:"video_height"`
	AverageFps        int           `json:"average_fps"`
	Delay             int           `json:"delay"`
	CreatedAt         time.Time     `json:"created_at"`
	IsPlaylist        bool          `json:"is_playlist"`
	StreamType        string        `json:"stream_type"`
	Preview           StreamPreview `json:"preview"`
	Channel           StreamChannel `json:"channel"`
}

// StreamPreview ...
type StreamPreview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}

// StreamChannel ...
type StreamChannel struct {
	Mature                       bool        `json:"mature"`
	Status                       string      `json:"status"`
	BroadcasterLanguage          string      `json:"broadcaster_language"`
	BroadcasterSoftware          string      `json:"broadcaster_software"`
	DisplayName                  string      `json:"display_name"`
	Game                         string      `json:"game"`
	Language                     string      `json:"language"`
	ID                           int         `json:"_id"`
	Name                         string      `json:"name"`
	CreatedAt                    time.Time   `json:"created_at"`
	UpdatedAt                    time.Time   `json:"updated_at"`
	Partner                      bool        `json:"partner"`
	Logo                         string      `json:"logo"`
	VideoBanner                  interface{} `json:"video_banner"`
	ProfileBanner                interface{} `json:"profile_banner"`
	ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
	URL                          string      `json:"url"`
	Views                        int         `json:"views"`
	Followers                    int         `json:"followers"`
	BroadcasterType              string      `json:"broadcaster_type"`
	Description                  string      `json:"description"`
	PrivateVideo                 bool        `json:"private_video"`
	PrivacyOptionsEnabled        bool        `json:"privacy_options_enabled"`
}

// StreamResponse ...
type StreamResponse struct {
	ResponseCommon
	Data Stream
}

// GetStream gets the stream referenced by the slug.
func (c *Client) GetStream(slug string) (*StreamResponse, error) {
	resp, err := c.get("/streams/"+slug, &Stream{}, nil)
	if err != nil {
		return nil, err
	}

	stream := &StreamResponse{}
	stream.Error = resp.Error
	stream.ErrorStatus = resp.ErrorStatus
	stream.ErrorMessage = resp.ErrorMessage
	stream.StatusCode = resp.StatusCode
	stream.Data = *resp.Data.(*Stream)

	return stream, nil
}


