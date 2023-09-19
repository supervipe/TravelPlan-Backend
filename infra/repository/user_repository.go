package repository

import (
	"backend/application/config"
	"backend/core/domain/models"
)

type UserRepository interface {
	FindById(id uint) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id string) error
}

type userRepository struct {
}

func UserRepositoryConstructor() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) FindById(id uint) (*models.User, error) {
	db := config.GetDB()
	user := models.User{}
	user.ID = id
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Create(user models.User) (*models.User, error) {
	db := config.GetDB()
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Update(user models.User) (*models.User, error) {
	return nil, nil
}

func (repository *userRepository) Delete(id string) error {
	return nil
}
