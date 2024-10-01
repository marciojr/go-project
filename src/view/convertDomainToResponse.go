package view

import (
	"github.com/marciojr/go-project/src/controller/model/response"
	"github.com/marciojr/go-project/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

func ConvertDomainListToResponse(
	userDomains []model.UserDomainInterface,
) []response.UserResponse {

	userResponses := make([]response.UserResponse, len(userDomains))

	for i, userDomain := range userDomains {
		userResponses[i] = ConvertDomainToResponse(userDomain)
	}

	return userResponses
}
