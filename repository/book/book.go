package book

import (
	_entities "be7/layered/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (ur *BookRepository) GetAll() ([]_entities.Book, error) {
	var books []_entities.Book
	tx := ur.DB.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}

func (ur *BookRepository) GetBookById(id int) (_entities.Book, error) {
	var books _entities.Book
	tx := ur.DB.Find(&books, id)
	if tx.Error != nil {
		return books, tx.Error
	}

	return books, nil
}

func (ur *BookRepository) CreateBook(request _entities.Book) (_entities.Book, error) {
	

	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request , yx.Error
	}

	return request, nil
}

func (ur *BookRepository) UpdateBook(id int, request _entities.Book) (_entities.Book, error) {
	err := ur.DB.Model(&_entities.Book{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		return request , err
	}

	return request, nil
}

func (ur *BookRepository) DeleteBook(id int) error {
	
	err := ur.DB.Unscoped().Delete(&_entities.Book{}, id).Error
	if err != nil {
		return err
	}

	return nil
}






