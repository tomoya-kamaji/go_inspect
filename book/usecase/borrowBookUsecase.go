package usecase

import (
	"github.com/tomoya_kamaji/go_inspect/book/domain"
)

type borrowBookUsecase struct {
	bookRepository domain.BookRepository
}

func NewBorrowBookUsecase(bookRepository domain.BookRepository) borrowBookUsecase {
	return borrowBookUsecase{
		bookRepository: bookRepository,
	}
}

func (u *borrowBookUsecase) BorrowBook(bookID string) error {
	book, err := u.bookRepository.Find(bookID)
	if err != nil {
		return err
	}

	err = book.Borrow()
	if err != nil {
		return err
	}

	err = u.bookRepository.Save(book)
	if err != nil {
		return err
	}

	return nil
}
