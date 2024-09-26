package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marciojr/go-project/src/configuration/database/mongodb"
	"github.com/marciojr/go-project/src/controller/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env file")
	}

	ctx := context.Background()

	database, err := mongodb.NewMongoDbConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
