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

func UserRepositoryConstructor() UserRepository {
	return &userRepository{}
}

func (repository *userRepository) FindById(id string) (*models.User, error) {
	return &models.User{
		Id:       "1",
		Name:     "John Doe",
		Email:    "jonhdoe@email.com",
		Password: "123456",
	}, nil
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
