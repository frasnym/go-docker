package main

import (
	"fmt"
	"os"

	_userHttpDeliver "github.com/frasnym/go-docker/user/delivery/http"
	_userRepo "github.com/frasnym/go-docker/user/repository"
	_userUcase "github.com/frasnym/go-docker/user/usecase"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	userRepo := _userRepo.NewPgSqlUserRepository()
	userUcase := _userUcase.NewUserUsecase(userRepo)
	_userHttpDeliver.NewUserHandler(app, userUcase)

	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8000"
	}
	app.Run(iris.Addr(fmt.Sprintf(":%s", port)))
}
