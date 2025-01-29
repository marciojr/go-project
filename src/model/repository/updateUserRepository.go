package repository

import (
	"context"
	"os"

	"github.com/marciojr/go-project/src/configuration/rest_err"
	"github.com/marciojr/go-project/src/model"
	"github.com/marciojr/go-project/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(
	ID string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.D{{Key: "_id", Value: objID}}

	value := converter.ConvertDomainToEntity(userDomain)

	update := bson.D{{Key: "$set", Value: value}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.ModifiedCount == 0 {
		return rest_err.NewNotFoundError("Some problem occurred while updating the user")
	}

	return nil
}
