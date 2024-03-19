package main

import (
	"log"

	authUsers "github.com/NIKFED/go-auth-users"
	"github.com/NIKFED/go-auth-users/pkg/handler"
	"github.com/NIKFED/go-auth-users/pkg/repository"
	"github.com/NIKFED/go-auth-users/pkg/service"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)

	srv := new(authUsers.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
