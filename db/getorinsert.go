package db

func GetOrInsertArtist(name string) (int64, error) {
	id, err := GetArtist(name)
	if err == nil {
		return id, nil
	}
	return InsertArtist(name)
}
