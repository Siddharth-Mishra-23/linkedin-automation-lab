package auth

import "context"

type Authenticator interface {
	Login(ctx context.Context) error
	IsAuthenticated(ctx context.Context) (bool, error)
	HandleCheckpoint(ctx context.Context) error
	LoadSession() error
	SaveSession() error
}
