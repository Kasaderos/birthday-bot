package core

import (
	"birthday-bot/internal/adapters/notifier"
	"birthday-bot/internal/domain/entities"
	"birthday-bot/internal/domain/errs"
	"context"
	"fmt"
	"time"
)

const DefaultListSize = int64(100)

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

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (c *User) NotifyBirthday(ctx context.Context) {
	offset := int64(0)
	limit := DefaultListSize
	for {
		users, err := c.r.repo.BirthdayUsersList(ctx, time.Now(), offset, limit)
		if err != nil {
			c.r.lg.Errorw("birthday users list", err)
			return
		}
		for _, user := range users {
			offset = max(offset, user.ID)

			select {
			case <-ctx.Done():
				return
			default:
			}

			err = c.r.notifier.Send(notifier.Message{
				ChatID:  user.TelegramChatID,
				Payload: fmt.Sprintf("Happy birthday, %s!", user.FirstName),
			})
			if err != nil {
				c.r.lg.Errorw("[notifier] send congrats", err, user.ID)
			}
		}

		// don't to next select
		if int64(len(users)) < limit {
			return
		}
	}
}
