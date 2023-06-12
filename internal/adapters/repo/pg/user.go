package pg

import (
	"birthday-bot/internal/domain/entities"
	"context"
	"time"
)

func (d *St) UserGet(ctx context.Context, id int64) (*entities.UserSt, error) {
	var result entities.UserSt

	var date time.Time
	err := d.db.QueryRow(ctx, `
		select
			id,
  			first_name,
  			last_name,
  			birthday,
  			telegram_chat_id
		from users
		where id = $1
	`, id).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&date,
		&result.TelegramChatID,
	)
	// comments: date isn't related with timezone
	// so postgres formats the date
	// automatically to the current timestamp
	// you can also expilicitly convert birthday to
	// timestamp like: birthday::timestamp
	result.Birthday = date.Format(time.DateOnly)

	return &result, err
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
		result["birthday"] = *obj.Birthday
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

func (d *St) BirthdayUsersList(ctx context.Context, t time.Time, offsetID, limit int64) ([]*entities.UserSt, error) {

	args := map[string]any{
		"offsetID": offsetID,
		"limit":    limit,
	}

	conds := []string{
		"DATE_PART('month', birthday) = DATE_PART('month', CURRENT_DATE)",
		"DATE_PART('day', birthday) = DATE_PART('day', CURRENT_DATE)",
		"id > ${offsetID}",
	}

	rows, err := d.db.QueryM(ctx, `
		select
			t.id,
			t.first_name,
			t.last_name,
			t.telegram_chat_id
		from users t
		`+d.tOptionalWhere(conds)+`
		limit ${limit}`,
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
