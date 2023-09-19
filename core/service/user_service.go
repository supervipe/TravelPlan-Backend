package service

import (
	"backend/core/domain/dto"
	"backend/core/domain/models"
	"backend/infra/repository"
	"strconv"
)

type UserService interface {
	GetUser(userId string) (*dto.UserDTO, error)
	CreateUser(user *dto.CreateUserDTO) (*dto.UserDTO, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func UserServiceConstructor(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) GetUser(userId string) (*dto.UserDTO, error) {
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return nil, err
	}
	user, err := service.userRepository.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *userService) CreateUser(user *dto.CreateUserDTO) (*dto.UserDTO, error) {
	userModel := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	userModel, err := service.userRepository.Create(*userModel)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Id:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}, nil
}
