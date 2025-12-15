package messaging

import "context"

type MessageTemplate struct {
	ID      string
	Content string
}

type Messenger interface {
	SendMessage(ctx context.Context, profileURL string, message string) error
}
