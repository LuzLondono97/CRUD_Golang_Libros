package repository

import (
	"context"
	"dojo_go/model"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetOne(ctx context.Context, id uint) (model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id uint, user model.User) error
	Delete(ctx context.Context, id uint) error
}