package repository

import (
	"gorm.io/gorm"
	"miniCloudStroage/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepository) Delete(userId uint64) error {
	return ur.db.Delete(&models.User{}, userId).Error
}

func (ur *UserRepository) GetById(userId uint64) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
