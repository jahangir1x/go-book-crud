package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type bookService struct {
	repo domain.IBookRepo
}

// interface binding
func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &bookService{
		repo: bookRepo,
	}
}

// all methods of interface are implemented
func (service *bookService) GetAllBooks() ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.repo.GetAllBooks()
	if len(book) == 0 {
		return nil, errors.New("No book found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			ID:          val.ID,
			BookName:    val.BookName,
			AuthorID:    val.AuthorID,
			Publication: val.Publication,
		})
	}
	return allBooks, nil
}
func (service *bookService) GetBook(bookID uint) (types.BookRequest, error) {
	bookDetail, err := service.repo.GetBook(bookID)
	book := types.BookRequest{
		ID:          bookDetail.ID,
		BookName:    bookDetail.BookName,
		AuthorID:    bookDetail.AuthorID,
		Publication: bookDetail.Publication,
	}
	if err != nil {
		return book, errors.New("No book found")
	}
	return book, nil
}
func (service *bookService) CreateBook(book *models.BookDetail) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("BookDetail was not created")
	}
	return nil
}

func (service *bookService) UpdateBook(updatedBook *models.BookDetail) error {

	existingBook, err := service.GetBook(uint(updatedBook.ID))
	if err != nil {
		return errors.New("No book found")
	}
	if updatedBook.BookName == "" {
		updatedBook.BookName = existingBook.BookName
	}
	if updatedBook.AuthorID == 0 {
		updatedBook.AuthorID = existingBook.AuthorID
	}
	if updatedBook.Publication == "" {
		updatedBook.Publication = existingBook.Publication
	}
	if err := service.repo.UpdateBook(updatedBook); err != nil {
		return errors.New("BookDetail update was unsuccessful")
	}
	return nil
}
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.repo.DeleteBook(bookID); err != nil {
		return errors.New("BookDetail deletion was unsuccessful")
	}
	return nil
}
