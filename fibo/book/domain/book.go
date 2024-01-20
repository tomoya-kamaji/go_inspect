package domain

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	ISBN       string
	IsBorrowed bool
}

func NewBook(id, title, author, isbn string) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		ISBN:       isbn,
		IsBorrowed: false,
	}
}

// 借りる
func (b *Book) Borrow() error {
	if b.IsBorrowed {
		return fmt.Errorf("book %s is already borrowed", b.ID)
	}
	b.IsBorrowed = true
	return nil
}

// 返す
func (b *Book) Return() error {
	if !b.IsBorrowed {
		return fmt.Errorf("book %s is not borrowed", b.ID)
	}
	b.IsBorrowed = false
	return nil
}
