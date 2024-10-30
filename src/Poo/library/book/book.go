package book

import (
	"fmt"
)

type Printable interface {
	ToString()
}

func Print(t Printable) {
	t.ToString()
}

// Modificador por convencion se usa mayusculas de inicio para atributos publicos
// Modificadores con minusculas son privados, por convención
type Book struct {
	title  string
	author string
	pages  int
}

type TextBook struct {
	Book
	editorrial string
	level      string
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func NewTextBook(title string, author string, pages int, editorrial string, level string) *TextBook {
	return &TextBook{
		Book:       Book{title, author, pages},
		editorrial: editorrial,
		level:      level,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) ToString() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.title, b.author, b.pages)
}

func (t *TextBook) ToString() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d Editorial: %s Level: %s\n", t.title, t.author, t.pages, t.editorrial, t.level)
}
