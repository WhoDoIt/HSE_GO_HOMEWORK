package shelf

import (
	"strconv"
)

type Book struct {
	Name   string
	Author string
	Pages  []string
	empty  bool
}

func (b Book) String() string {
	if b.empty {
		return "empty book"
	}
	return "\"" + b.Name + "\" written by " + b.Author + " with " + strconv.Itoa(len(b.Pages)) + " pages"
}

type Searcher interface {
	Search(id string) (Book, bool)
	AddBook(id string, book Book) (Searcher, bool)
	RemoveBook(id string) (Searcher, bool)
	GetAllBooks() []Book
}

type MapStorage map[string]Book

type SliceStorage []Book

func (storage MapStorage) Search(id string) (Book, bool) {
	book, ok := (storage)[id]
	if !ok {
		return Book{empty: true}, false
	}
	return book, ok
}

func (storage MapStorage) AddBook(id string, book Book) (Searcher, bool) {
	storage[id] = book
	return storage, true
}

func (storage MapStorage) RemoveBook(id string) (Searcher, bool) {
	delete(storage, id)
	return storage, true
}

func (storage MapStorage) GetAllBooks() []Book {
	var res []Book
	for _, v := range storage {
		if !v.empty {
			res = append(res, v)
		}
	}
	return res
}

func (storage SliceStorage) Search(id string) (Book, bool) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return Book{empty: true}, false
	}

	if intId < 0 || intId >= len(storage) {
		return Book{empty: true}, false
	}

	return (storage)[intId], true
}

func (storage SliceStorage) AddBook(id string, book Book) (Searcher, bool) {
	intId, err := strconv.Atoi(id)
	if err != nil || intId < 0 {
		return storage, false
	}
	for intId >= len(storage) {
		storage = append(storage, Book{empty: true})
	}
	storage[intId] = book
	return storage, true
}

func (storage SliceStorage) RemoveBook(id string) (Searcher, bool) {
	intId, err := strconv.Atoi(id)
	if err != nil || intId < 0 || intId >= len(storage) {
		return storage, false
	}
	storage[intId] = Book{empty: true}
	return storage, true
}

func (storage SliceStorage) GetAllBooks() []Book {
	var res []Book
	for _, v := range storage {
		if !v.empty {
			res = append(res, v)
		}
	}
	return res
}
