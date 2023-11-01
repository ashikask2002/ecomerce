//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/ashikask2002/ecomerce.git/pkg/api"
	"github.com/ashikask2002/ecomerce.git/pkg/api/handler"
	"github.com/ashikask2002/ecomerce.git/pkg/config"
	"github.com/ashikask2002/ecomerce.git/pkg/db"
	"github.com/ashikask2002/ecomerce.git/pkg/repository"
	"github.com/ashikask2002/ecomerce.git/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP)
	return &http.ServerHTTP{}, nil
}
