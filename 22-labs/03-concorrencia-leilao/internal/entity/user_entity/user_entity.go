package user_entity

import (
	"auction_go/internal/internal_error"
	"context"
)

type User struct {
	Id   string
	Name string
	//	Username string
	//	Password string
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, userId string) (*User, *internal_error.InternalError)
}
