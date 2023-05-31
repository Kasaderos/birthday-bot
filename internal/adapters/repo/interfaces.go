package repo

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

type Repo interface {
	// city
	CityGet(ctx context.Context, id string) (*entities.CitySt, error)
	CityList(ctx context.Context, pars *entities.CityListParsSt) ([]*entities.CitySt, error)
	CityIdExists(ctx context.Context, id string) (bool, error)
	CityCreate(ctx context.Context, obj *entities.CityCUSt) (string, error)
	CityUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error
	CityDelete(ctx context.Context, id string) error

	// // user
	// UserGet(ctx context.Context, id string) (*entities.CitySt, error)
	// UserCreate(ctx context.Context, obj *entities.CityCUSt) (string, error)
	// UserUpdate(ctx context.Context, id string, obj *entities.CityCUSt) error
	// UserDelete(ctx context.Context, id string) error
}
