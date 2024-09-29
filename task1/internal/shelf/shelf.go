package shelf

import (
	"strconv"
)

type Book struct {
	Name   string
	Author string
}

type Searcher interface {
	Search(id string) (Book, bool)
}

type Library interface {
	Search(name string) (Book, bool)
}

type MapStorage map[string]Book

type SliceStorage []Book

func (storage MapStorage) Search(id string) (Book, bool) {
	book, ok := storage[id]
	if !ok {
		return Book{}, false
	}
	return book, ok
}

func (storage SliceStorage) Search(id string) (Book, bool) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return Book{}, false
	}

	if intId < 0 || intId >= len(storage) {
		return Book{}, false
	}

	return storage[intId], true
}

type MyLibrary struct {
	Storage Searcher
}

// func (self MyLibrary) Search(name string) (Book, bool) {

// }
