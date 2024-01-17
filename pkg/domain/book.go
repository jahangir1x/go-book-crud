package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

// for database repository operation (call from service)
type IBookRepo interface {
	GetAllBooks() []models.BookDetail
	GetBook(bookID uint) (models.BookDetail, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(book *models.BookDetail) error
	DeleteBook(bookID uint) error
}

// for service operation (response to controller | call from controller)
type IBookService interface {
	GetAllBooks() ([]types.BookRequest, error)
	GetBook(bookID uint) (types.BookRequest, error)
	CreateBook(book *models.BookDetail) error
	UpdateBook(bookRequest types.BookRequest, book *models.BookDetail) error
	DeleteBook(bookID uint) error
}
