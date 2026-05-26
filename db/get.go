package db

import (
	"fmt"
)

func GetArtist(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM artists WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get artist: %w", err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get artist: %w", err)
	}

	return id, nil
}

func GetAlbum(name string) (int64, error) {
	stmt, err := DB.Prepare("SELECT id FROM albums WHERE name = ?")
	if err != nil {
		return 0, fmt.Errorf("prepare get album: %w", err)
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("get album: %w", err)
	}

	return id, nil
}
