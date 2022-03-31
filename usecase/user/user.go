package user

import (
	_entities "be7/layered/entities"
	_userRepository "be7/layered/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.CreateUser(request)
	return users, err
}

func (uuc *UserUseCase) UpdateUser(id int, request _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.UpdateUser(id, request)
	return users, err
}

func (uuc *UserUseCase) DeleteUser(id int) error {
	err := uuc.userRepository.DeleteUser(id)
	return err
}



func (uuc *UserUseCase) GetUserById(id int) (_entities.User, error) {
	users, err := uuc.userRepository.GetUserById(id)
	return users, err
}