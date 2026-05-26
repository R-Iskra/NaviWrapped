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
