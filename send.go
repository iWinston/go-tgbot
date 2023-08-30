package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Send(target int64, message string) error {
	msg := tgbotapi.NewMessage(target, message)
	_, err := TgBot.Send(msg)
	return err
}
