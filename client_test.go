package tgbot

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewClient(t *testing.T) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("无法加载 .env 文件：", err)
	}

	tgBot, err := NewClient(os.Getenv("ProxyUrl"), os.Getenv("TgBotToken"))
	if err != nil {
		t.Error(err)
	}
	if tgBot == nil {
		t.Error("tgBot is nil")
	}
}
