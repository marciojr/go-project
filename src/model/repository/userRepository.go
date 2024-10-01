package repository

import (
	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_COLLECTION = "MONGODB_COLLECTION"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(
		ID string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(
		ID string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}
