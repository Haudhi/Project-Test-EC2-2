package user

import (
	_entities "be7/layered/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetUserById(id int) (_entities.User, error) {
	var users _entities.User
	tx := ur.DB.Find(&users, id)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) CreateUser(request _entities.User) (_entities.User, error) {
	

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request , yx.Error
	}

	return request, nil
}

func (ur *UserRepository) UpdateUser(id int, request _entities.User) (_entities.User, error) {
	err := ur.DB.Model(&_entities.User{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		return request , err
	}

	return request, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	
	err := ur.DB.Unscoped().Delete(&_entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}





