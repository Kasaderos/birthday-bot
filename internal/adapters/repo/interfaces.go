package repo

import (
	"birthday-bot/internal/domain/entities"
	"context"
	"time"
)

type Repo interface {
	// user
	UserGet(ctx context.Context, id int64) (*entities.UserSt, error)
	UserCreate(ctx context.Context, obj *entities.UserCUSt) (int64, error)
	UserUpdate(ctx context.Context, id int64, obj *entities.UserCUSt) error
	UserDelete(ctx context.Context, id int64) error
	BirthdayUsersList(ctx context.Context, t time.Time) ([]*entities.UserSt, error)
}
