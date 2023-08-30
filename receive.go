package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ListenUpdates(handle func(update tgbotapi.Update)) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := TgBot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		handle(update)
	}

	return nil
}
