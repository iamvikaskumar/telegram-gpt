package main

import (
	"github.com/iamvikaskumar/telegram-gpt/config"
	"github.com/iamvikaskumar/telegram-gpt/pkg/telegram"
)

func main() {
	token := config.GetConfig().TelegramToken
	telegramClient := telegram.NewClient(token)
	telegramClient.SendMesage(17, 6880161489, "kite baje sona hai ?")
}
