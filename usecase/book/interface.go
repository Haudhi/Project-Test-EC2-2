package book

import (
	_entities "be7/layered/entities"
)

type BookUseCaseInterface interface {
	GetAll() ([]_entities.Book, error)
	CreateBook(request _entities.Book) (_entities.Book, error)
	UpdateBook(id int, request _entities.Book) (_entities.Book, error)
	DeleteBook(id int) (error)
	GetBookById(id int) (_entities.Book, error)
}