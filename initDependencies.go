package main

import (
	"github.com/marciojr/go-project/src/controller"
	"github.com/marciojr/go-project/src/model/repository"
	"github.com/marciojr/go-project/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
