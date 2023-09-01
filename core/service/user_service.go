package service

import (
	"backend/core/domain/dto"
	"backend/infra/repository"
)

type UserService interface {
	GetUser(userId string) (*dto.UserDTO, error)
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
	user, err := service.userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
