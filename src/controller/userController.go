package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marciojr/go-project/src/model/service"
)

type UserControllerInterface interface {
	FindAllUsers(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}
