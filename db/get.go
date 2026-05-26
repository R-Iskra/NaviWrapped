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
