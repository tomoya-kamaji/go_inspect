package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/tomoya_kamaji/go_inspect/book/domain"
	"github.com/tomoya_kamaji/go_inspect/book/domain/mocks"
	"github.com/tomoya_kamaji/go_inspect/book/usecase"
)

const bookID = "1"

func TestBorrowBookUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mockを準備
	book := domain.NewBook(bookID, "Test Book", "Test Author", "1234567890")

	mockRepo := mocks.NewMockBookRepository(ctrl)
	mockRepo.EXPECT().Find(bookID).Return(book, nil).Times(1)
	mockRepo.EXPECT().Save(&isBorrowedMatcher{isBorrowed: true}).Return(nil).Times(1)

	u := usecase.NewBorrowBookUsecase(mockRepo)

	err := u.BorrowBook(bookID)
	assert.NoError(t, err)
}

type isBorrowedMatcher struct {
	isBorrowed bool
}

func (m *isBorrowedMatcher) Matches(x interface{}) bool {
	book, ok := x.(*domain.Book)
	if !ok {
		return false
	}
	return book.IsBorrowed
}
func (m *isBorrowedMatcher) String() string {
	return "isBorrowedMatcherが失敗したよ"
}
