package rent

import (
	_entities "be7/layered/entities"
	"fmt"

	"gorm.io/gorm"
)

type RentRepository struct {
	DB *gorm.DB
}

func NewRentRepository(db *gorm.DB) *RentRepository {
	return &RentRepository{
		DB: db,
	}
	
}

func (ur *RentRepository) GetAll() ([]_entities.Rent, error) {
	var rents []_entities.Rent
	tx := ur.DB.Find(&rents)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rents, nil
}

func (ur *RentRepository) GetRentById(id int) (_entities.Rent, error) {
	var rents _entities.Rent
	tx := ur.DB.Find(&rents, id)
	if tx.Error != nil {
		return rents, tx.Error
	}

	return rents, nil
}

func (ur *RentRepository) CreateRent(request _entities.Rent) (_entities.Rent, error) {
	

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request , yx.Error
	}

	tx := ur.DB.Model(&_entities.Book{}).Where("id = ?", request.BookID).Update("status", "unavailable")
	if tx.Error != nil {
		return request , tx.Error
	}

	return request, nil
}

func (ur *RentRepository) Return(id _entities.Rent) error {

	yx := ur.DB.Model(&_entities.Rent{}).Where("id = ?", id.ID).Updates(&id)
	fmt.Println(id.ID)
	if yx.Error != nil {
		return yx.Error
	}

	return nil
}








