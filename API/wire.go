//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"inititaryplanner/controllers"
	"inititaryplanner/dal"
	"inititaryplanner/dal/db"
	"inititaryplanner/service"
)

func InitializeMainController() (MainController, error) {
	wire.Build(
		db.GetMainMongoDatabase,

		dal.NewAttractionDal,

		service.NewAttractionService,

		controllers.NewAttractionController,

		NewMainController,
	)
	return &mainController{}, nil
}
