package pg

import (
	"birthday-bot/internal/domain/entities"
	"context"
)

func (d *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	var result entities.UserSt

	// err := d.db.QueryRow(ctx, `
	// 	select
	// 		t.id,
	// 		t.code,
	// 		t.name,
	// 		t.org_bin
	// 	from city t
	// 	where t.id = $1
	// `, id).Scan(
	// 	&result.Id,
	// 	&result.Code,
	// 	&result.Name,
	// 	&result.OrgBin,
	// )
	// if errors.Is(err, db.ErrNoRows) {
	// 	return nil, nil
	// }

	return &result, nil
}
