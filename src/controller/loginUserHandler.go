package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciojr/go-project/src/controller/model/request"
	"github.com/marciojr/go-project/src/view"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {

	var loginUserRequest request.LoginUserRequest

	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Email or password are wrong": err.Error()})
		return
	}

	domainResult, err := uc.service.LoginUser(loginUserRequest.Email, loginUserRequest.Password)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
