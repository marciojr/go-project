package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marciojr/go-project/src/controller"
	"github.com/marciojr/go-project/src/controller/routes"
	"github.com/marciojr/go-project/src/model/service"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
