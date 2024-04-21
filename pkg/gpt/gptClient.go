package gpt

import (
	"context"
	"errors"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

type GPT struct {
	*openai.Client
}

func GetClient(token string) *GPT {
	client := openai.NewClient(token)
	return &GPT{client}
}

func (c *GPT) GetReponse(ctx context.Context, text string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: text,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return "", err
	}
	defer stream.Close()

	var res string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			//exit out of the loop when the steam ends
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return "", err
		}
		res += response.Choices[0].Delta.Content
	}
	return res, nil
}
