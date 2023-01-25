package repository

import (
	"github.com/gabrielmvas/user-api-golang/model"

	"context"
)


type Repository interface {
	GetUser(ctx context.Context, email string) (model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
