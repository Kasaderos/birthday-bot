package pg

import (
	"birthday-bot/internal/adapters/db"
	"birthday-bot/internal/domain/entities"
	"context"
	"errors"
	"fmt"
	"time"
)

func (d *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	var result entities.UserSt

	err := d.db.QueryRow(ctx, `
		select
			id,
  			first_name,
  			last_name,
  			birthday,
  			telegram_chat_id
		from user
		where t.id = $1
	`, id).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Birthday,
		&result.TelegramChatID,
	)
	if errors.Is(err, db.ErrNoRows) {
		return nil, nil
	}

	return &result, nil
}

func (d *St) UserUpdate(ctx context.Context, id int64, obj *entities.UserCUSt) error {
	fields := d.getUserFields(obj)
	cols := d.tPrepareFieldsToUpdate(fields)

	fields["cond_id"] = id

	return d.db.ExecM(ctx, `
		update users 
		set `+cols+`
		where id = ${cond_id}
	`, fields)
}

func (d *St) getUserFields(obj *entities.UserCUSt) map[string]any {
	result := map[string]any{}

	if obj.FirstName != nil {
		result["first_name"] = *obj.FirstName
	}

	if obj.LastName != nil {
		result["last_name"] = *obj.LastName
	}

	if obj.Birthday != nil {
		result["birthday"] = fmt.Sprintf("TO_DATE(%s, YYYY-MM-DD)", *obj.Birthday)
	}

	if obj.TelegramChatID != nil {
		result["telegram_chat_id"] = *obj.TelegramChatID
	}

	return result
}

func (d *St) UserCreate(ctx context.Context, obj *entities.UserCUSt) (int64, error) {
	fields := d.getUserFields(obj)
	cols, values := d.tPrepareFieldsToCreate(fields)

	var newId int64

	err := d.db.QueryRowM(ctx, `
		insert into users (`+cols+`)
		values (`+values+`)
		returning id
	`, fields).Scan(&newId)

	return newId, err
}

func (d *St) UserDelete(ctx context.Context, id int64) error {
	return d.db.Exec(ctx, `
		delete from users 
		where id = $1
	`, id)
}

func (d *St) BirthdayUsersList(ctx context.Context, t time.Time) ([]*entities.UserSt, error) {
	birthDate := t.Format(time.DateOnly)

	args := map[string]any{
		"birth_date": birthDate,
	}

	conds := []string{
		"DATE_PART('month', ${birth_date}) = DATE_PART('month', CURRENT_DATE)",
		"DATE_PART('day', ${birth_date}) = DATE_PART('day', CURRENT_DATE)",
	}

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.first_name,
			t.last_name,
			t.telegram_chat_id
		from users t
		`+d.tOptionalWhere(conds),
		args,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entities.UserSt, 0)

	for rows.Next() {
		item := &entities.UserSt{}

		err = rows.Scan(
			&item.ID,
			&item.FirstName,
			&item.LastName,
			&item.TelegramChatID,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
