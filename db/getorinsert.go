package db

func GetOrInsertArtist(name string) (int64, error) {
	id, err := GetArtist(name)
	if err == nil {
		return id, nil
	}
	return InsertArtist(name)
}

func GetOrInsertAlbum(name string, artistName string) (int64, error) {
	id, err := GetAlbum(name)
	if err == nil {
		return id, nil
	}

	artistID, err := GetOrInsertArtist(artistName)
	if err != nil {
		return 0, err
	}

	return InsertAlbum(name, artistID)
}

func GetOrInsertSong(name string, albumName string, artistName string) (int64, error) {
	id, err := GetSong(name)
	if err == nil {
		return id, nil
	}

	albumID, err := GetOrInsertAlbum(albumName, artistName)
	if err != nil {
		return 0, err
	}

	return InsertSong(name, albumID)
}

func GetOrInsertGenre(name string) (int64, error) {
	id, err := GetGenre(name)
	if err == nil {
		return id, nil
	}
	return InsertGenre(name)
}

func GetOrInsertUser(name string) (int64, error) {
	id, err := GetUser(name)
	if err == nil {
		return id, nil
	}
	return InsertUser(name)
}
