package repository

import "backend/core/domain/models"

type UserRepository interface {
	FindById(id string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id string) error
}

type userRepository struct {
}

func (repository *userRepository) FindById(id string) (*models.User, error) {
	return nil, nil
}

func (repository *userRepository) Create(user models.User) (*models.User, error) {
	return nil, nil
}

func (repository *userRepository) Update(user models.User) (*models.User, error) {
	return nil, nil
}

func (repository *userRepository) Delete(id string) error {
	return nil
}
