// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ashikask2002/ecomerce.git/pkg/api"
	"github.com/ashikask2002/ecomerce.git/pkg/api/handler"
	"github.com/ashikask2002/ecomerce.git/pkg/config"
	"github.com/ashikask2002/ecomerce.git/pkg/db"
	"github.com/ashikask2002/ecomerce.git/pkg/repository"
	"github.com/ashikask2002/ecomerce.git/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}