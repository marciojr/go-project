package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciojr/go-project/src/view"
)

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {}

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	ID := c.Param("id")

	domainResult, err := uc.service.FindUserById(ID)
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	email := c.Param("email")

	domainResult, err := uc.service.FindUserByEmail(email)
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
