package repository

import (
	"context"
	"os"

	"github.com/marciojr/go-project/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(
	ID string,
) *rest_err.RestErr {

	collection_name := os.Getenv(MONGODB_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.D{{Key: "_id", Value: objID}}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.DeletedCount == 0 {
		return rest_err.NewNotFoundError("There is no user with this id")
	}

	return nil
}
