package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) LoginUser(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	user, err := ud.userRepository.FindUserByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
