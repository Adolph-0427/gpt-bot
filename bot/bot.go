package bot

import (
	"context"
	"gpt-bot/library/api"
)

type Bot struct {
	openAPi api.OpenApi
}

func NewBot() *Bot {
	openApi := &api.Api{}
	return &Bot{
		openAPi: openApi,
	}
}

func (b *Bot) Chat(ctx context.Context, content string) (string, error) {
	result, err := b.openAPi.Chat(ctx, content)
	if err != nil {
		return "", err
	}
	return result.Choices[0].Message.Content, nil
}
