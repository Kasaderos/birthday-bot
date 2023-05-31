package usecases

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

func (u *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	return u.cr.User.Get(ctx, id, true)
}
