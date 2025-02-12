package user

import (
	"auction_go/configuration/logger"
	"auction_go/internal/entity/user_entity"
	"auction_go/internal/internal_error"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) FindUserById(ctx context.Context, userId string) (*user_entity.User, *internal_error.InternalError) {
	filter := bson.M{"_id": userId}

	var userEntityMongo UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("User not found with this id = %s. Error: %s", userId, err), err)
			return nil, internal_error.NewInternalServerError(fmt.Sprintf("User not found with this id = %s. Error: %s", userId, err))
		}

		logger.Error("Error trying to find user by userId. Error: %s", err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find user by userId. Error: %s", err))
	}

	userEntity := &user_entity.User{
		Id:   userEntityMongo.Id,
		Name: userEntityMongo.Name,
	}

	return userEntity, nil

}
