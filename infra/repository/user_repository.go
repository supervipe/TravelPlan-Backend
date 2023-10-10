package repository

import (
	"backend/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	database *gorm.DB
}

func UserRepositoryConstructor(db *gorm.DB) UserRepository {
	return &userRepository{
		database: db,
	}
}

func (repository *userRepository) FindById(id uint) (*models.User, error) {
	user := models.User{}
	user.ID = id
	err := repository.database.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Create(user models.User) (*models.User, error) {
	err := repository.database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Update(user models.User) (*models.User, error) {
	err := repository.database.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) Delete(id uint) error {
	user := models.User{}
	user.ID = id
	err := repository.database.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
