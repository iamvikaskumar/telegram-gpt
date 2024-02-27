package telegram

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iamvikaskumar/telegram-gpt/config"
)

const (
	topicPrefix  = "/topic"
	phrasePrefix = "/phrase"
)

type Client struct {
	*tgbotapi.BotAPI
}

// NewClient creates a client that uses the given RPC client.

// func NewClient(c *rpc.Client) *Client {
func NewClient(token string) *Client {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Client{bot}
}

func GetClient() {
	token := config.GetConfig().TelegramToken
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// check if message has '/topic' or '/phrase' as prefix
			if !strings.HasPrefix(update.Message.Text, topicPrefix) && !strings.HasPrefix(update.Message.Text, phrasePrefix) {
				continue
			}
			var msg tgbotapi.MessageConfig
			if strings.HasPrefix(update.Message.Text, topicPrefix) {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "you asked a topic...")
				msg.ReplyToMessageID = update.Message.MessageID
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "you asked a phrase...")
				msg.ReplyToMessageID = update.Message.MessageID
			}

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func (c *Client) Listen() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// check if message has '/topic' or '/phrase' as prefix
			if !strings.HasPrefix(update.Message.Text, topicPrefix) && !strings.HasPrefix(update.Message.Text, phrasePrefix) {
				continue
			}
			var msg tgbotapi.MessageConfig
			if strings.HasPrefix(update.Message.Text, topicPrefix) {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "you asked a topic...")
				msg.ReplyToMessageID = update.Message.MessageID
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "you asked a phrase...")
				msg.ReplyToMessageID = update.Message.MessageID
			}
			c.Send(msg)
		}
	}
}

func (c *Client) SendMesage(messageID int, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	//msg.ReplyToMessageID = messageID
	c.Send(msg)
}
