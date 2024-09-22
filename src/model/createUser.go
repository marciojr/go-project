package model

import (
	"fmt"

	"github.com/marciojr/go-project/src/configuration/rest_err"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {

	ud.EncryptPassword()

	fmt.Print()

	return nil
}
