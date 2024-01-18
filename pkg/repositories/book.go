package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

// parent struct to implement interface binding
type bookRepo struct {
	db *gorm.DB
}

// interface binding
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

// all methods of interface are implemented
func (repo *bookRepo) GetAllBooks(request map[string]string) []models.BookDetail {
	var book []models.BookDetail
	parsedSchema, err := schema.Parse(&models.BookDetail{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		panic("Error in parsing schema")
	}
	err = repo.db.Find(&book).Error
	for key, value := range request {
		if value != "" {
			mappedName := parsedSchema.FieldsByName[key].DBName
			err = repo.db.Where(mappedName+" = ?", value).Find(&book).Error
		}
	}
	if err != nil {
		return []models.BookDetail{}
	}
	return book
}
func (repo *bookRepo) GetBook(bookID uint) (models.BookDetail, error) {
	var book models.BookDetail
	if err := repo.db.Where("id = ?", bookID).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}
func (repo *bookRepo) CreateBook(book *models.BookDetail) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) UpdateBook(book *models.BookDetail) error {
	if err := repo.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) DeleteBook(bookID uint) error {
	var Book models.BookDetail
	if err := repo.db.Where("id = ?", bookID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
func (repo *bookRepo) DeleteBooksByAuthorID(authorID uint) error {
	var Book models.BookDetail
	if err := repo.db.Where("author_id = ?", authorID).Delete(&Book).Error; err != nil {
		return err
	}
	return nil
}
