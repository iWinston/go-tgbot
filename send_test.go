package tgbot

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("无法加载 .env 文件：", err)
	}

	Init(os.Getenv("ProxyUrl"), os.Getenv("TgBotToken"))

	botId, err := strconv.ParseInt(os.Getenv("TgChatId"), 10, 64)
	if err != nil {
		panic(err)
	}

	err = Send(botId, "test")
	assert.NoError(t, err)
}
