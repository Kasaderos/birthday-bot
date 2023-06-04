package usecases

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

func (u *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	return u.cr.User.Get(ctx, id, true)
}

func (u *St) UserUpdate(ctx context.Context,
	id int64, obj *entities.UserCUSt) error {

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.User.Update(ctx, id, obj)
	})
}

func (u *St) UserCreate(ctx context.Context,
	obj *entities.UserCUSt) (int64, error) {
	var err error
	var result int64

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.User.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) UserDelete(ctx context.Context,
	id int64) error {

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.User.Delete(ctx, id)
	})
}
