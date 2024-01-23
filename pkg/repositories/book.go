package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
)

// BookRepo defines the methods of the domain.IBookRepo interface.
type BookRepo struct {
	db *gorm.DB
}

// BookDBInstance returns a new instance of the BookRepo struct.
func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &BookRepo{
		db: d,
	}
}

// GetFilteredBooks returns a list of books filtered by the request.
func (repo *BookRepo) GetFilteredBooks(request map[string]string) ([]models.BookDetail, error) {
	// get all books
	var bookDetails []models.BookDetail
	if err := repo.db.Find(&bookDetails).Error; err != nil {
		return nil, err
	}

	// parse the schema
	parsedSchema, err := schema.Parse(&models.BookDetail{}, &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		return nil, err
	}

	// filter the authors for each field in the request
	for key, value := range request {
		mappedFieldInDB := parsedSchema.FieldsByName[key].DBName
		err = repo.db.Where(mappedFieldInDB+" = ?", value).Find(&bookDetails).Error
		if err != nil {
			return nil, err
		}
	}

	return bookDetails, nil
}

// GetBook returns a book by the bookID.
func (repo *BookRepo) GetBook(bookID uint) (*models.BookDetail, error) {
	bookDetail := &models.BookDetail{}
	if err := repo.db.Where("id = ?", bookID).First(bookDetail).Error; err != nil {
		return nil, err
	}
	return bookDetail, nil
}

// CreateBook creates a new book with given book details.
func (repo *BookRepo) CreateBook(book *models.BookDetail) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBook updates a book with given book details.
func (repo *BookRepo) UpdateBook(book *models.BookDetail) error {
	if err := repo.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBook deletes a book with the given bookID
func (repo *BookRepo) DeleteBook(bookID uint) error {
	bookDetail := &models.BookDetail{}
	if err := repo.db.Where("id = ?", bookID).Delete(bookDetail).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBooksByAuthorID deletes books by authorID.
func (repo *BookRepo) DeleteBooksByAuthorID(authorID uint) error {
	bookDetail := &models.BookDetail{}
	if err := repo.db.Where("author_id = ?", authorID).Delete(bookDetail).Error; err != nil {
		return err
	}
	return nil
}
