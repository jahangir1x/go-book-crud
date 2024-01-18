package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type bookService struct {
	bookRepo   domain.IBookRepo
	authorRepo domain.IAuthorRepo
}

// interface binding
func BookServiceInstance(bookRepo domain.IBookRepo, authorRepo domain.IAuthorRepo) domain.IBookService {
	return &bookService{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

// all methods of interface are implemented
func (service *bookService) GetAllBooks(request map[string]string) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.bookRepo.GetAllBooks(request)
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
	bookDetail, err := service.bookRepo.GetBook(bookID)
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
	_, err := service.authorRepo.GetAuthor(book.AuthorID)
	if err != nil {
		return errors.New("Author ID not found")
	}
	if err := service.bookRepo.CreateBook(book); err != nil {
		return errors.New("BookDetail was not created")
	}
	return nil
}

func (service *bookService) UpdateBook(updatedBook *models.BookDetail) error {

	existingBook, err := service.GetBook(updatedBook.ID)
	if err != nil {
		return errors.New("No book found")
	}
	if updatedBook.BookName == "" {
		updatedBook.BookName = existingBook.BookName
	}
	if err != nil {
		return errors.New("Author ID not found")
	}
	if updatedBook.AuthorID == 0 {
		updatedBook.AuthorID = existingBook.AuthorID
	}
	if updatedBook.Publication == "" {
		updatedBook.Publication = existingBook.Publication
	}
	if err := service.bookRepo.UpdateBook(updatedBook); err != nil {
		return errors.New("BookDetail update was unsuccessful")
	}
	return nil
}
func (service *bookService) DeleteBook(bookID uint) error {
	if err := service.bookRepo.DeleteBook(bookID); err != nil {
		return errors.New("BookDetail deletion was unsuccessful")
	}
	return nil
}
