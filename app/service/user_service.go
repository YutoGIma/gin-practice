package service

import (
	"myapp/app/model"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s UserService) GetUsers() ([]model.User, error) {
	var users []model.User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s UserService) CreateUser(user model.User) error {
	return s.DB.Create(user).Error
}

func (s UserService) GetUser(id uint) (model.User, error) {
	var user model.User
	err := s.DB.First(user, id).Error
	return user, err
}

func (s UserService) UpdateUser(user model.User) error {
	return s.DB.Save(user).Error
}

func (s UserService) DeleteUser(id uint) error {
	return s.DB.Delete(model.User{}, id).Error
}
