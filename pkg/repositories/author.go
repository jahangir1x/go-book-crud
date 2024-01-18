package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func AuthorDBInstance(d *gorm.DB) domain.IAuthorRepo {
	return &AuthorRepo{
		db: d,
	}
}

func (repo *AuthorRepo) GetAllAuthors() []models.AuthorDetail {
	var author []models.AuthorDetail
	err := repo.db.Find(&author).Error
	if err != nil {
		return []models.AuthorDetail{}
	}
	return author
}

func (repo *AuthorRepo) GetAuthor(authorID uint) (models.AuthorDetail, error) {
	var author models.AuthorDetail
	if err := repo.db.Where("id = ?", authorID).First(&author).Error; err != nil {
		return author, err
	}
	return author, nil
}

func (repo *AuthorRepo) CreateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Create(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AuthorRepo) UpdateAuthor(author *models.AuthorDetail) error {
	if err := repo.db.Save(author).Error; err != nil {
		return err
	}
	return nil
}

func (repo *AuthorRepo) DeleteAuthor(authorID uint) error {
	var author models.AuthorDetail
	if err := repo.db.Where("id = ?", authorID).Delete(&author).Error; err != nil {
		return err
	}
	return nil
}
