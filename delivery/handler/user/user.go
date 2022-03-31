package user

import (
	"be7/layered/delivery/helper"
	_middlewares "be7/layered/delivery/middlewares"
	_userUseCase "be7/layered/usecase/user"
	"fmt"
	"net/http"
	"strconv"

	_entities "be7/layered/entities"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(u _userUseCase.UserUseCaseInterface) UserHandler {
	return UserHandler{
		userUseCase: u,
	}
}

func (uh *UserHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all users", users))
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {
		var param _entities.User

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		users, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success create user", users))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)

		var param _entities.User
		id, _ := strconv.Atoi(c.Param("id"))

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
		users, err := uh.userUseCase.UpdateUser(id, param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update data", users))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		
		id, _ := strconv.Atoi(c.Param("id"))

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		

		err := uh.userUseCase.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success delete user", err))
	}
}

func (uh *UserHandler) GetUserByIdHandler() echo.HandlerFunc {
	
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		fmt.Println("id token", idToken)
		
		id, err := strconv.Atoi(c.Param("id"))

		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		
		users, err := uh.userUseCase.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get user by id", users))
	}
}
