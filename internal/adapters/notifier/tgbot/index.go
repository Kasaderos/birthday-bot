package tgbot

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/notifier"
)

type St struct {
	lg logger.Lite
}

func New(lg logger.Lite) *St {
	return &St{
		lg: lg,
	}
}

func (o *St) Send(msg notifier.Message) error {
	return nil
}
