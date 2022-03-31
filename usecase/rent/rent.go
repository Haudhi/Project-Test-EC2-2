package rent

import (
	_entities "be7/layered/entities"
	_bookRepository "be7/layered/repository/book"
	_rentRepository "be7/layered/repository/rent"
	"errors"
	"time"
)

type RentUseCase struct {
	rentRepository _rentRepository.RentRepositoryInterface
	bookRepository _bookRepository.BookRepositoryInterface
}

func NewRentUseCase(rentRepo _rentRepository.RentRepositoryInterface, bookRepo _bookRepository.BookRepositoryInterface) RentUseCaseInterface {
	return &RentUseCase{
		rentRepository: rentRepo,
		bookRepository: bookRepo,
	}
}

func (uuc *RentUseCase) GetAll() ([]_entities.Rent, error)  {
	rents, err := uuc.rentRepository.GetAll()
	return rents, err
}

func (uuc *RentUseCase) CreateRent(request _entities.Rent) (_entities.Rent, error) {
	book, err := uuc.bookRepository.GetBookById(int(request.BookID))
	if book.Status != "available" {
		return _entities.Rent{}, errors.New("Cannot rent this book")
	}

	if book.UserID == int(request.UserID) {
		return _entities.Rent{}, errors.New("you cant rent your own book, this book is already yours")
	}


	rents, err := uuc.rentRepository.CreateRent(request)
	return rents, err
}

func (uuc *RentUseCase) GetRentById(id int) (_entities.Rent, error) {
	rents, err := uuc.rentRepository.GetRentById(id)
	return rents, err
}

func (uuc *RentUseCase) Return(id _entities.Rent) (error) {
	var book _entities.Book

	rentbook, err := uuc.rentRepository.GetRentById(int(id.ID))

	if rentbook.Status == "returned" {
		return errors.New("cannot return this book")
	}

	rentbook.Status = "returned"
	rentbook.ReturnDate = time.Now()
	uuc.rentRepository.Return(rentbook)

	book.Status = "available"
	uuc.bookRepository.UpdateBook(int(rentbook.BookID), book)

	return err
}