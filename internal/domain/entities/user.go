package entities

type UserSt struct {
	ID             int64  `db:"id" json:"id"`
	FirstName      string `db:"first_name" json:"first_name"`
	LastName       string `db:"last_name" json:"last_name"`
	Birthday       string `db:"birthday" json:"birthday"`
	TelegramChatID int64  `db:"telegram_chat_id" json:"telegram_chat_id"`
}
