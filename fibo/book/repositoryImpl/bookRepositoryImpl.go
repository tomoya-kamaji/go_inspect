package repositoryimpl

import "github.com/tomoya_kamaji/go_inspect/book/domain"

type BookRepositoryImpl struct{}

func NewInMemoryBookRepository() domain.BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) Save(book *domain.Book) error {
	// 本来はDBに保存する処理
	return nil
}

func (r *BookRepositoryImpl) Find(id string) (*domain.Book, error) {
	// 本来はDBから取得する処理
	testBook := domain.NewBook("1", "エンジニアリング組織論への招待", "佐々木 紀章", "978-4-7741-7568-3")
	return testBook, nil
}
