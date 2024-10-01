package repository

import (
	"context"
	"os"

	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
	"github.com/marciojr/go-project/src/model/repository/entity"
	"github.com/marciojr/go-project/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserById(
	ID string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.D{{Key: "_id", Value: objID}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("There is no user with this id")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {

	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("There is no user with this email")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}
