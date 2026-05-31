package db

import (
	"fmt"
	"time"
)

func InsertArtist(name string) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO artists (name) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert artist %s: %w", name, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, fmt.Errorf("insert artist %s: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get artist id %s: %w", name, err)
	}

	return id, nil
}

func InsertAlbum(name string, artistID int64) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO albums (name, artist_id) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert album %s: %w", name, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, artistID)
	if err != nil {
		return 0, fmt.Errorf("insert album %s: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get album id %s: %w", name, err)
	}

	return id, nil
}

func InsertSong(name string, albumID int64) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO songs (name, album_id) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert song %s: %w", name, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, albumID)
	if err != nil {
		return 0, fmt.Errorf("insert song %s: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get song id %s: %w", name, err)
	}

	return id, nil
}

func InsertGenre(name string) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO genres (name) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert genre %s: %w", name, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, fmt.Errorf("insert genre %s: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get genre id %s: %w", name, err)
	}

	return id, nil
}

func InsertSongGenre(songID int64, genreID int64) error {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO song_genres (song_id, genre_id) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("prepare insert song %d genre %d: %w", songID, genreID, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(songID, genreID)
	if err != nil {
		return fmt.Errorf("insert song %d genre %d: %w", songID, genreID, err)
	}

	return nil
}

func InsertUser(name string) (int64, error) {
	stmt, err := DB.Prepare("INSERT OR IGNORE INTO users (name) VALUES (?)")
	if err != nil {
		return 0, fmt.Errorf("prepare insert user %s: %w", name, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, fmt.Errorf("insert user %s: %w", name, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get user id %s: %w", name, err)
	}

	return id, nil
}

func InsertScrobble(userID int64, songID int64, submissionTime time.Time, duration int64) error {
	stmt, err := DB.Prepare(`INSERT OR IGNORE INTO scrobbles (user_id, song_id, submission_time, duration)
							 VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("prepare insert scrobble (%d, %d, %v, %d): %w", userID, songID, submissionTime, duration, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, songID, submissionTime, duration)
	if err != nil {
		return fmt.Errorf("insert scrobble (%d, %d, %v, %d): %w", userID, songID, submissionTime, duration, err)
	}

	return nil
}
