package repository

import (
	"go_gin_boilerplate/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := ur.DB.Find(&users).Error
	return users, err
}

func (ur *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := ur.DB.First(&user, id).Error
	return user, err
}

func (ur *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	return ur.DB.Delete(user).Error
}

// Example of additional functionality based on new fields
func (ur *UserRepository) FindActiveUsers() ([]models.User, error) {
	var users []models.User
	err := ur.DB.Where("is_active = ?", true).Find(&users).Error
	return users, err
}
