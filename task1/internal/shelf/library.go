package shelf

type Library interface {
	Search(name string) (Book, bool)
	AddBook(book Book) (Library, bool)
	RemoveBook(book Book) (Library, bool)
	GetAllBooks() []Book
	ReplaceStorage(searcher Searcher) Library
	ReplaceIdGenerator(gen func(string) string) Library
}

type MyLibrary struct {
	storage Searcher
	gen     func(string) string
}

func NewMyLibrary(gen func(string) string) Library {
	return MyLibrary{storage: SliceStorage{}, gen: gen}
}
func (l MyLibrary) Search(name string) (Book, bool) {
	return l.storage.Search(l.gen(name))
}

func (l MyLibrary) AddBook(book Book) (Library, bool) {
	storage, ok := l.storage.AddBook(l.gen(book.Name), book)
	l.storage = storage
	return l, ok
}

func (l MyLibrary) RemoveBook(book Book) (Library, bool) {
	storage, ok := l.storage.RemoveBook(l.gen(book.Name))
	l.storage = storage
	return l, ok
}

func (l MyLibrary) GetAllBooks() []Book {
	return l.storage.GetAllBooks()
}

func (l MyLibrary) ReplaceStorage(searcher Searcher) Library {
	books := l.storage.GetAllBooks()
	l.storage = searcher
	for _, v := range books {
		l.storage, _ = l.storage.AddBook(l.gen(v.Name), v)
	}
	return l
}

func (l MyLibrary) ReplaceIdGenerator(gen func(string) string) Library {
	books := l.storage.GetAllBooks()
	for _, v := range books {
		l.storage, _ = l.storage.RemoveBook(l.gen(v.Name))
	}
	l.gen = gen
	for _, v := range books {
		l.storage, _ = l.storage.AddBook(l.gen(v.Name), v)
	}
	return l
}
