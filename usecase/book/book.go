package book

import (
	_entities "be7/layered/entities"
	_bookRepository "be7/layered/repository/book"
)

type BookUseCase struct {
	bookRepository _bookRepository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo _bookRepository.BookRepositoryInterface) BookUseCaseInterface {
	return &BookUseCase{
		bookRepository: bookRepo,
	}
}

func (uuc *BookUseCase) GetAll() ([]_entities.Book, error) {
	books, err := uuc.bookRepository.GetAll()
	return books, err
}

func (uuc *BookUseCase) CreateBook(request _entities.Book) (_entities.Book, error) {
	books, err := uuc.bookRepository.CreateBook(request)
	return books, err
}

func (uuc *BookUseCase) UpdateBook(id int, request _entities.Book) (_entities.Book, error) {
	books, err := uuc.bookRepository.UpdateBook(id, request)
	return books, err
}

func (uuc *BookUseCase) DeleteBook(id int) error {
	err := uuc.bookRepository.DeleteBook(id)
	return err
}



func (uuc *BookUseCase) GetBookById(id int) (_entities.Book, error) {
	books, err := uuc.bookRepository.GetBookById(id)
	return books, err
}
