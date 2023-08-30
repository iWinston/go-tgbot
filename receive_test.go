package tgbot

import (
	"fmt"
	"log"
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func handle(update tgbotapi.Update) {
	fmt.Println("群组ID", update.Message.Chat.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Text {
	case "/start":
		msg.Text = "欢迎使用机器人！"
	case "/help":
		msg.Text = "这是帮助信息。"
	default:
		msg.Text = "我不明白你说的是什么。"
	}

	_, err := TgBot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func TestHandleUpdates(t *testing.T) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("无法加载 .env 文件：", err)
	}

	Init(os.Getenv("ProxyUrl"), os.Getenv("TgBotToken"))

	err = ListenUpdates(handle)
	if err != nil {
		t.Error(err)
	}
}
