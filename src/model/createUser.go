package model

import (
	"fmt"

	"github.com/marciojr/go-project/src/configuration/rest_err"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	ud.EncryptPassword()

	fmt.Print()

	return nil
}
