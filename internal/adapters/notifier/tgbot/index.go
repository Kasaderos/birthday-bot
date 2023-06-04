package tgbot

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/notifier"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type St struct {
	lg  logger.Lite
	bot *tgbotapi.BotAPI
}

func New(lg logger.Lite, token string) (*St, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &St{
		lg:  lg,
		bot: bot,
	}, nil
}

func (o *St) Send(msg notifier.Message) error {
	// todo write message templates
	_, err := o.bot.Send(
		tgbotapi.NewMessage(
			msg.ChatID,
			msg.Payload,
		))
	return err
}
