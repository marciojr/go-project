package service

import (
	"fmt"

	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}
