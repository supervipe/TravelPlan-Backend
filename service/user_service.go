package service

import (
	"backend/domain/dto"
	"backend/domain/models"
	"backend/infra/repository"
	"errors"
	"strconv"
)

type UserService interface {
	GetUser(userId string) (*dto.UserDTO, error)
	CreateUser(userCreateDTO *dto.CreateUserDTO) (*dto.UserDTO, error)
	UpdateUserPassword(userUpdateDTO *dto.UpdateUserPasswordDTO, userId string) (*dto.UserDTO, error)
	DeleteUser(userId string) error
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

func (service *userService) CreateUser(userCreateDTO *dto.CreateUserDTO) (*dto.UserDTO, error) {
	userModel := &models.User{
		Name:     userCreateDTO.Name,
		Email:    userCreateDTO.Email,
		Password: userCreateDTO.Password,
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

func (service *userService) UpdateUserPassword(userUpdateDTO *dto.UpdateUserPasswordDTO, userId string) (*dto.UserDTO, error) {
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return nil, err
	}
	user, err := service.userRepository.FindById(uint(id))
	if err != nil {
		return nil, err
	}

	if user.Password != userUpdateDTO.OldPassword {
		return nil, errors.New("old password is wrong")
	}

	user.Password = userUpdateDTO.NewPassword
	user, err = service.userRepository.Update(*user)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *userService) DeleteUser(userId string) error {
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return err
	}
	err = service.userRepository.Delete(uint(id))
	if err != nil {
		return err
	}

	return nil
}
