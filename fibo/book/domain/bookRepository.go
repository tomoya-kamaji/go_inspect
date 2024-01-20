package domain

type BookRepository interface {
	Save(book *Book) error
	Find(id string) (*Book, error)
}
