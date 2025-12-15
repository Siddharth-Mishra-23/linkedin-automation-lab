package connect

import "context"

type Connector interface {
	SendConnection(ctx context.Context, profileURL string, note string) error
	AlreadySent(profileURL string) bool
}
