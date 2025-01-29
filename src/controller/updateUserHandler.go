package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciojr/go-project/src/configuration/validation"
	"github.com/marciojr/go-project/src/controller/model/request"
	"github.com/marciojr/go-project/src/model"
	"github.com/marciojr/go-project/src/view"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	var userRequest request.UserRequest

	ID := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(ID, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	domain.SetID(ID)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
