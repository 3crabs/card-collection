package storage

func NewStorageMemory() Storage {
	return newStorageMemory()
}

func NewStorageDb() Storage {
	return newStorageDb()
}
