package user

import (
	_entities "be7/layered/entities"
)

type UserRepositoryInterface interface {
	GetAll() ([]_entities.User, error)
	GetUserById(id int) (_entities.User, error)
	CreateUser(request _entities.User) (_entities.User, error)
	UpdateUser(id int, request _entities.User) (_entities.User, error)
	DeleteUser(id int) (error)
}
