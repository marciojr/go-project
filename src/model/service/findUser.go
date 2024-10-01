package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) FindUserById(ID string) (model.UserDomainInterface, *rest_err.RestErr) {

	user, err := ud.userRepository.FindUserById(ID)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	user, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
