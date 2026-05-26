package db

import (
	"fmt"
)

func InsertArtist(name string) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO artists (name) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert artist: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, fmt.Errorf("insert artist: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get artist id: %w", err)
	}

	return id, nil
}

func InsertAlbum(name string, artistID int64) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO albums (name, artist_id) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert album: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, artistID)
	if err != nil {
		return 0, fmt.Errorf("insert album: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get album id: %w", err)
	}

	return id, nil
}

func InsertSong(name string, albumID int64) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO songs (name, album_id) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert song: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, albumID)
	if err != nil {
		return 0, fmt.Errorf("insert song: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get song id: %w", err)
	}

	return id, nil
}
