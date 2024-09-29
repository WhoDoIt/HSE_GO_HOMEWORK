package main

import (
	"github.com/WhoDoIt/HSE_GO_HOMEWORK/task1/internal/shelf"
)

func main() {
	b1 := shelf.Book{Name: "Kto Ya?", Author: "Ya"}
	b2 := shelf.Book{Name: "Kto Ya? 2", Author: "Ne Ya"}
	b3 := shelf.Book{Name: "Harry Potter", Author: "Donald Trump"}
	slice := []shelf.Book{b1, b2, b3}
	library := shelf.MyLibrary{Storage: }
}
