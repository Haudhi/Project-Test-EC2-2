package main

import (
	"be7/layered/configs"
	_userHandler "be7/layered/delivery/handler/user"
	_userRepository "be7/layered/repository/user"
	_userUseCase "be7/layered/usecase/user"

	_bookHandler "be7/layered/delivery/handler/book"
	_bookRepository "be7/layered/repository/book"
	_bookUseCase "be7/layered/usecase/book"

	_authHandler "be7/layered/delivery/handler/auth"
	_authRepository "be7/layered/repository/auth"
	_authUseCase "be7/layered/usecase/auth"

	_rentHandler "be7/layered/delivery/handler/rent"
	_rentRepository "be7/layered/repository/rent"
	_rentUseCase "be7/layered/usecase/rent"

	"fmt"
	"log"

	_middlewares "be7/layered/delivery/middlewares"
	_routes "be7/layered/delivery/routes"
	_utils "be7/layered/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	bookRepo := _bookRepository.NewBookRepository(db)
	bookUseCase := _bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := _bookHandler.NewBookHandler(bookUseCase)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	rentRepo := _rentRepository.NewRentRepository(db)
	rentUseCase := _rentUseCase.NewRentUseCase(rentRepo, bookRepo)
	rentHandler := _rentHandler.NewRentHandler(rentUseCase)



	
	e := echo.New()
	// e.Use(_middlewares.JWTMiddleware())
	_routes.RegisterPath(e, userHandler)
	_routes.RegisterPathh(e, bookHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.RentEndPoint(e, rentHandler)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
	// log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))

}


