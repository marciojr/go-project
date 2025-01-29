package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	user, err := ud.FindUserByEmail(userDomain.GetEmail())

	if user != nil {
		return nil, rest_err.NewBadRequestError("Email already exists")
	}

	if err != nil {
		return nil, err
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
