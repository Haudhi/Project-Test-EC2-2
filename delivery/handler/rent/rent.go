package rent

import (
	"be7/layered/delivery/helper"
	_rentUseCase "be7/layered/usecase/rent"
	"net/http"
	"strconv"

	_entities "be7/layered/entities"

	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	rentUseCase _rentUseCase.RentUseCaseInterface
}

func NewRentHandler(u _rentUseCase.RentUseCaseInterface) RentHandler {
	return RentHandler{
		rentUseCase: u,
	}
}

func (uh *RentHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		rents, err := uh.rentUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all rented book", rents))
	}
}

func (uh *RentHandler) CreateRentHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		var param _entities.Rent

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		rents, err := uh.rentUseCase.CreateRent(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success rent a book", rents))
	}
}


func (uh *RentHandler) GetRentByIdHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		
		rents, err := uh.rentUseCase.GetRentById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get rented book by id", rents))
	}
}

func (uh *RentHandler) ReturnHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		var param _entities.Rent
		

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		err = uh.rentUseCase.Return(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success return a book", param))
		
	}
}
