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
