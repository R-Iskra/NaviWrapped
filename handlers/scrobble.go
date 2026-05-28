package handlers

type additionalInfo struct {
	DurationMs int `json:"duration_ms,omitempty"`
}

type trackMetadata struct {
	ArtistName     string         `json:"artist_name,omitempty"`
	ReleaseName    string         `json:"release_name,omitempty"`
	TrackName      string         `json:"track_name,omitempty"`
	AdditionalInfo additionalInfo `json:"additional_info"`
}

type listenData struct {
	ListenedAt    int           `json:"listened_at,omitempty"`
	TrackMetadata trackMetadata `json:"track_metadata"`
}

type submitPayload struct {
	ListenType string       `json:"listen_type,omitempty"`
	Payload    []listenData `json:"payload,omitempty"`
}
