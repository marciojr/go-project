package service

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
)

func (ud *userDomainService) LoginUser(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	user, err := ud.userRepository.FindUserByEmail(email)

	if user == nil {
		return nil, rest_err.NewForbiddenError("Your email or password is incorrect")
	}

	if err != nil {
		return nil, err
	}

	if !user.ComparePassword(password) {
		return nil, rest_err.NewForbiddenError("Your email or password is incorrect")
	}

	return user, nil
}
