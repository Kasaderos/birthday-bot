package core

import (
	"birthday-bot/internal/adapters/notifier"
	"birthday-bot/internal/domain/entities"
	"birthday-bot/internal/domain/errs"
	"context"
	"fmt"
	"time"
)

type User struct {
	r *St
}

func NewUser(r *St) *User {
	return &User{r: r}
}

func (c *User) Validate(ctx context.Context, obj *entities.UserCUSt) error {
	return nil
}

func (c *User) Get(ctx context.Context, id int64, errNE bool) (*entities.UserSt, error) {
	result, err := c.r.repo.UserGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, errs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *User) Create(ctx context.Context, obj *entities.UserCUSt) (int64, error) {
	var err error

	err = c.Validate(ctx, obj)
	if err != nil {
		return -1, err
	}

	// create
	result, err := c.r.repo.UserCreate(ctx, obj)
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (c *User) Update(ctx context.Context, id int64, obj *entities.UserCUSt) error {
	var err error

	err = c.Validate(ctx, obj)
	if err != nil {
		return err
	}

	err = c.r.repo.UserUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) Delete(ctx context.Context, id int64) error {
	return c.r.repo.UserDelete(ctx, id)
}

func (c *User) NotifyBirthday(ctx context.Context) {
	users, err := c.r.repo.BirthdayUsersList(ctx, time.Now())
	if err != nil {
		c.r.lg.Errorw("birthday users list", err)
	}
	for _, user := range users {
		select {
		case <-ctx.Done():
			return
		default:
		}

		err = c.r.notifier.Send(notifier.Message{
			ChatID:  user.TelegramChatID,
			Payload: fmt.Sprintf("Happy birthday!"),
		})
		if err != nil {
			c.r.lg.Errorw("send congrats", err, user.ID)
		}
	}
}
