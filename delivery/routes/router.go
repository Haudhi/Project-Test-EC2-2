package routes

import (
	_authHandler "be7/layered/delivery/handler/auth"
	_bookHandler "be7/layered/delivery/handler/book"
	_rentHandler "be7/layered/delivery/handler/rent"
	_userHandler "be7/layered/delivery/handler/user"
	_middlewares "be7/layered/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.GET("/users/:id", uh.GetUserByIdHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterPathh(e *echo.Echo, uh _bookHandler.BookHandler) {
	e.GET("/books", uh.GetAllHandler())
	e.GET("/books/:id", uh.GetBookByIdHandler())
	e.POST("/books", uh.CreateBookHandler(), _middlewares.JWTMiddleware())
	e.PUT("/books/:id", uh.UpdateBookHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/books/:id", uh.DeleteBookHandler(), _middlewares.JWTMiddleware())
}

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RentEndPoint(e *echo.Echo, uh _rentHandler.RentHandler) {
	e.GET("/rent", uh.GetAllHandler())
	e.GET("/rent/:id", uh.GetRentByIdHandler())
	e.POST("/rent", uh.CreateRentHandler())
	e.POST("/return", uh.ReturnHandler())
	
}

