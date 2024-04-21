package telegram

import (
	"context"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/iamvikaskumar/telegram-gpt/pkg/gpt"
)

const (
	topicPrefix  = "/topic"
	phrasePrefix = "/phrase"
	gptPrefix    = "/ask_gpt"
)

type Client struct {
	*tgbotapi.BotAPI
	GptClient *gpt.GPT
}

// NewClient creates a client that uses the given RPC client.

func NewClient(token string, gptClient *gpt.GPT) *Client {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Client{bot, gptClient}
}

func (c *Client) Listen() {
	ctx := context.Background()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.GetUpdatesChan(u)
	var prompt string
	for update := range updates {

		if update.Message != nil {
			var msg tgbotapi.MessageConfig
			// check if message has '/topic' or '/phrase' as prefix
			if !strings.HasPrefix(update.Message.Text, gptPrefix) {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command. Please use command /ask_gpt followed by your question")
			} else {
				prompt = strings.TrimPrefix(update.Message.Text, gptPrefix)
				if prompt == "" {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "please type a question")
				} else {
					res, err := c.GptClient.GetReponse(ctx, prompt)
					if err != nil {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error generating reponse...")
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, res)
					}
				}
			}
			msg.ReplyToMessageID = update.Message.MessageID
			c.Send(msg)
		}
	}
}
