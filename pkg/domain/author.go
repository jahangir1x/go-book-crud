package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

type IAuthorRepo interface {
	GetAllAuthors() []models.AuthorDetail
	GetAuthor(authorID uint) (models.AuthorDetail, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(author *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}

type IAuthorService interface {
	GetAllAuthors() ([]types.AuthorRequest, error)
	GetAuthor(authorID uint) (types.AuthorRequest, error)
	CreateAuthor(author *models.AuthorDetail) error
	UpdateAuthor(updatedAuthor *models.AuthorDetail) error
	DeleteAuthor(authorID uint) error
}
