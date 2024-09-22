package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marciojr/go-project/src/controller"
)

func InitRoutes(r *gin.RouterGroup, uc controller.UserControllerInterface) {

	r.GET("/users", uc.FindAllUsers)
	r.GET("/users/:id", uc.FindUserById)
	r.GET("/users/email/:email", uc.FindUserByEmail)
	r.POST("/users", uc.CreateUser)
	r.PUT("/users", uc.UpdateUser)
	r.DELETE("/users", uc.DeleteUser)
}
