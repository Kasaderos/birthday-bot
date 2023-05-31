package usecases

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

func (u *St) UserGet(ctx context.Context, id string) (*entities.UserSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.User.Get(ctx, id, true)
}
