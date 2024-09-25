package model

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) SetID(ID string) {
	ud.ID = ID
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) SetEmail(email string) {
	ud.Email = email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) SetPassword(password string) {
	ud.Password = password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) SetName(name string) {
	ud.Name = name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) SetAge(age int8) {
	ud.Age = age
}
