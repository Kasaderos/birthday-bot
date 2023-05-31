package repo

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

type Repo interface {
	// user
	UserGet(ctx context.Context, id int64) (*entities.UserSt, error)
	// UserCreate(ctx context.Context, obj *entities.CityCUSt) (string, error)
	// UserUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error
	// UserDelete(ctx context.Context, id string) error
}
