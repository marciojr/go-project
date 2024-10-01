package controller

import "github.com/gin-gonic/gin"

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	ID := c.Param("id")

	err := uc.service.DeleteUser(ID)
	if err != nil {
		c.JSON(err.Code, err.Error())
	}
}
