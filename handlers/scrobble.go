package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/R-Iskra/NaviWrapped/db"
)

type additionalInfo struct {
	DurationMs int64 `json:"duration_ms,omitempty"`
}

type trackMetadata struct {
	ArtistName     string         `json:"artist_name,omitempty"`
	ReleaseName    string         `json:"release_name,omitempty"`
	TrackName      string         `json:"track_name,omitempty"`
	AdditionalInfo additionalInfo `json:"additional_info"`
}

type listenData struct {
	ListenedAt    int64         `json:"listened_at,omitempty"`
	TrackMetadata trackMetadata `json:"track_metadata"`
}

type submitPayload struct {
	ListenType string       `json:"listen_type,omitempty"`
	Payload    []listenData `json:"payload,omitempty"`
}

func SubmitListens(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Token ")

	userID, err := db.GetOrInsertUser(token)
	if err != nil {
		http.Error(w, "error resolving user", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "error reading request body", http.StatusBadRequest)
		return
	}

	var payload submitPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if payload.ListenType == "playing_now" {
		w.WriteHeader(http.StatusOK)
		return
	}

	for _, listen := range payload.Payload {
		if listen.ListenedAt == 0 || listen.TrackMetadata.ArtistName == "" ||
			listen.TrackMetadata.ReleaseName == "" || listen.TrackMetadata.TrackName == "" ||
			listen.TrackMetadata.AdditionalInfo.DurationMs == 0 {
			http.Error(w, "missing required fields", http.StatusBadRequest)
			return
		}

		songID, err := db.GetOrInsertSong(
			listen.TrackMetadata.TrackName,
			listen.TrackMetadata.ReleaseName,
			listen.TrackMetadata.ArtistName,
		)
		if err != nil {
			http.Error(w, "error saving listen", http.StatusInternalServerError)
			return
		}

		err = db.InsertScrobble(userID, songID, time.Unix(listen.ListenedAt, 0), listen.TrackMetadata.AdditionalInfo.DurationMs)
		if err != nil {
			http.Error(w, "error saving scrobble", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
