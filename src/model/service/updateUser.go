package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) UpdateUser(ID string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	user, err := ud.FindUserById(ID)

	if user == nil {
		return rest_err.NewBadRequestError("There is no user with this ID")
	}

	if err != nil {
		return err
	}

	return ud.userRepository.UpdateUser(ID, userDomain)
}
