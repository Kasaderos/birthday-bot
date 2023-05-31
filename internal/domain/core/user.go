package core

import (
	"birthday-bot/internal/domain/entities"
	"birthday-bot/internal/domain/errs"
	"context"
)

type User struct {
	r *St
}

func NewUser(r *St) *User {
	return &User{r: r}
}

// func (c *User) ValidateCU(ctx context.Context, obj *entities.UserCUSt, id string) error {
// 	// forCreate := id == ""

// 	return nil
// }

// func (c *User) List(ctx context.Context, pars *entities.UserListParsSt) ([]*entities.UserSt, error) {
// 	items, err := c.r.repo.UserList(ctx, pars)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return items, nil
// }

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

// func (c *User) IdExists(ctx context.Context, id string) (bool, error) {
// 	return c.r.repo.UserIdExists(ctx, id)
// }

// func (c *User) Create(ctx context.Context, obj *entities.UserCUSt) (string, error) {
// 	var err error

// 	err = c.ValidateCU(ctx, obj, "")
// 	if err != nil {
// 		return "", err
// 	}

// 	// create
// 	result, err := c.r.repo.UserCreate(ctx, obj)
// 	if err != nil {
// 		return "", err
// 	}

// 	return result, nil
// }

// func (c *User) Update(ctx context.Context, id string, obj *entities.UserCUSt) error {
// 	var err error

// 	err = c.ValidateCU(ctx, obj, id)
// 	if err != nil {
// 		return err
// 	}

// 	err = c.r.repo.UserUpdate(ctx, id, obj)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (c *User) Delete(ctx context.Context, id string) error {
// 	return c.r.repo.UserDelete(ctx, id)
// }

func (c *User) NotifyUserBirthday() {

}
