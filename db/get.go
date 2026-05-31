package db

import (
	"fmt"
)

func GetArtist(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM artists WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get artist %s: %w", name, err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get artist %s: %w", name, err)
	}

	return id, nil
}

func GetAlbum(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM albums WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get album %s: %w", name, err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get album %s: %w", name, err)
	}

	return id, nil
}

func GetSong(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM songs WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get song %s: %w", name, err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get song %s: %w", name, err)
	}

	return id, nil
}

func GetGenre(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM genres WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get genre %s: %w", name, err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get genre %s: %w", name, err)
	}

	return id, nil
}

func GetUser(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM users WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get user %s: %w", name, err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get user %s: %w", name, err)
	}

	return id, nil
}

func GetSongGenres(songID int64) ([]int64, error) {
	stmt, err := DB.Prepare("SELECT genre_id FROM song_genres WHERE song_id = ?")
	if err != nil {
		return nil, fmt.Errorf("prepare get song genres %d: %w", songID, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(songID)
	if err != nil {
		return nil, fmt.Errorf("get song genres %d: %w", songID, err)
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan genre id %d: %w", songID, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func GetGenreSongs(genreID int64) ([]int64, error) {
	stmt, err := DB.Prepare("SELECT song_id FROM song_genres WHERE genre_id = ?")
	if err != nil {
		return nil, fmt.Errorf("prepare get genre songs %d: %w", genreID, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(genreID)
	if err != nil {
		return nil, fmt.Errorf("get genre songs %d: %w", genreID, err)
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan song id %d: %w", genreID, err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}
