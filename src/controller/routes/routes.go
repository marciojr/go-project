package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marciojr/go-project/src/controller/handlers"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", controller.FindAllUsers)
	r.GET("/users/:id", controller.FindUserById)
	r.GET("/users/email/:email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users", controller.UpdateUser)
	r.DELETE("/users", controller.DeleteUser)
}
