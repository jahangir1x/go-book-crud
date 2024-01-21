package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &UserRepo{
		db: d,
	}
}

func (repo *UserRepo) GetUser(username string) (models.UserDetail, error) {
	var user models.UserDetail
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(user *models.UserDetail) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
