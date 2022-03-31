package rent

import (
	_entities "be7/layered/entities"
)

type RentRepositoryInterface interface {
	GetAll() ([]_entities.Rent, error)
	CreateRent(request _entities.Rent) (_entities.Rent, error)
	GetRentById(id int) (_entities.Rent, error)
	Return(id _entities.Rent) error
}