package main

import (
	"fmt"
	"strconv"

	"github.com/WhoDoIt/HSE_GO_HOMEWORK/task1/internal/shelf"
)

func main() {
	b1 := shelf.Book{Name: "Kto Ya?", Author: "Ya"}
	b2 := shelf.Book{Name: "Kto Ya? 2", Author: "Ne Ya"}
	b3 := shelf.Book{Name: "Harry Potter", Author: "Donald Trump"}
	library := shelf.NewMyLibrary(func(s string) string {
		return strconv.Itoa(len(s))
	})
	library, _ = library.AddBook(b1)
	library, _ = library.AddBook(b2)
	library, _ = library.AddBook(b3)

	for _, v := range library.GetAllBooks() {
		fmt.Println(v)
	}

	library = library.ReplaceStorage(shelf.MapStorage{})
	fmt.Println()
	for _, v := range library.GetAllBooks() {
		fmt.Println(v)
	}

	b4 := shelf.Book{Name: "Kak stat millionerom", Author: "Andrew LTake"}
	b5 := shelf.Book{Name: "Top 5 oshibok chelevochestva", Author: "Anonymous"}
	b6 := shelf.Book{Name: "Harry Potter 2", Author: "My Grandma"}

	library, _ = library.AddBook(b4)
	library, _ = library.AddBook(b5)
	library, _ = library.AddBook(b6)

	fmt.Println()
	for _, v := range library.GetAllBooks() {
		fmt.Println(v)
	}
}
