package tgbot

import (
	"log"
	"net/http"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var TgBot *tgbotapi.BotAPI

func Init(proxyUrl, botToken string) {
	var err error
	TgBot, err = NewClient(proxyUrl, botToken)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func NewClient(proxyUrl, botToken string) (*tgbotapi.BotAPI, error) {

	// 读取环境变量设置代理
	proxyURL, _ := url.Parse(proxyUrl)

	// 创建带有代理的 HTTP 客户端
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	// 创建 Bot API 客户端
	bot, err := tgbotapi.NewBotAPIWithClient(botToken, httpClient)

	if err != nil {
		return nil, err
	}

	log.Printf("连接到 Telegram 服务器，Bot 用户名为 %s", bot.Self.UserName)
	return bot, nil
}
