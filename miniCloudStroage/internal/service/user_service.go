package service

import (
	"miniCloudStroage/internal/models"
	"miniCloudStroage/internal/repository"
)

type UserService struct {
	ur *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{ur: ur}
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.ur.Create(user)
}

func (us *UserService) UpdateUser(user *models.User) error {
	return us.ur.Update(user)
}

func (us *UserService) GetById(userId uint64) (*models.User, error) {
	return us.ur.GetById(userId)
}

func (us *UserService) DeleteUser(userId uint64) error {
	return us.ur.Delete(userId)
}
