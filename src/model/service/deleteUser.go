package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUser(ID string) *rest_err.RestErr {

	err := ud.userRepository.DeleteUser(ID)
	if err != nil {
		return err
	}
	return nil
}
