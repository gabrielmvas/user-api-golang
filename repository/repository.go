package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gabrielmvas/user-api-golang/model"
)


var (ErrUserNotFound = errors.New("User not found"))

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (rep repository) GetUser(ctx context.Context, email string) (model.User, error) {
	var out user
	err := rep.db.
		Collection("users").
		FindOne(ctx, bson.M{"email": email}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return toOutputModel(out), nil
}

func (rep repository) GetUsers(ctx context.Context) ([]model.User, error) {
	cursor, err := rep.db.Collection("users").Find(ctx, bson.D{{}})

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var users []model.User

	err = cursor.All(ctx, &users)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return users, nil
}

func (rep repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	out, err := rep.db.Collection("users").InsertOne(ctx, fromInputModel(user))
	
	if err != nil {
		return model.User{}, err
	}
	
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (rep repository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	in := bson.M{}
	
	if user.FirstName != "" {
		in["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		in["last_name"] = user.LastName
	}
	if user.Password != "" {
		in["password"] = user.Password
	}
	
	out, err := rep.db.
		Collection("users").
		UpdateOne(ctx, bson.M{"email": user.Email}, bson.M{"$set": in})
	
		if err != nil {
		return model.User{}, err
	}
	if out.MatchedCount == 0 {
		return model.User{}, ErrUserNotFound
	}
	return user, nil
}

func (rep repository) DeleteUser(ctx context.Context, email string) error {
	out, err := rep.db.
		Collection("users").
		DeleteOne(ctx, bson.M{"email": email})
	
		if err != nil {
		return err
	}
	
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

type user struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName string `bson:"last_name,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func fromInputModel(userModel model.User) user {
	return user{
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
		Email: userModel.Email,
		Password: userModel.Password,
	}
}

func toOutputModel(userModel user) model.User {
	return model.User{
		ID: userModel.ID.String(),
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
		Email: userModel.Email,
		Password: userModel.Password,
	}
}
