package book

import (
	"be7/layered/delivery/helper"
	_middlewares "be7/layered/delivery/middlewares"
	_bookUseCase "be7/layered/usecase/book"
	"fmt"
	"net/http"
	"strconv"

	_entities "be7/layered/entities"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(u _bookUseCase.BookUseCaseInterface) BookHandler {
	return BookHandler{
		bookUseCase: u,
	}
}

func (uh *BookHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := uh.bookUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all users", books))
	}
}

func (uh *BookHandler) CreateBookHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		var param _entities.Book

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		books, err := uh.bookUseCase.CreateBook(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create book", books))
	}
}

func (uh *BookHandler) UpdateBookHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		var param _entities.Book
		id, _ := strconv.Atoi(c.Param("id"))

		getid, err := uh.bookUseCase.GetBookById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}


		fmt.Println("id param user id", getid.UserID)

		if idToken != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized2"))
		}

		
		err = c.Bind(&param)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		books, err := uh.bookUseCase.UpdateBook(id, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data", books))
	}
}

func (uh *BookHandler) DeleteBookHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		
		id, _ := strconv.Atoi(c.Param("id"))

		getid, err := uh.bookUseCase.GetBookById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}

		fmt.Println("id param user id", getid.UserID)

		if idToken != getid.UserID {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		
		err = uh.bookUseCase.DeleteBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success delete book", err))
	}
}

func (uh *BookHandler) GetBookByIdHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		

		books, err := uh.bookUseCase.GetBookById(id)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get book by id", books))
	}
}
