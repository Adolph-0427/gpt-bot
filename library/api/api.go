package api

import "context"

type OpenApi interface {
	Chat(ctx context.Context, content string) (*ChatResult, error)
}
