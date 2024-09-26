package model

import "encoding/json"

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(email, password, name string,
	age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (ud *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
