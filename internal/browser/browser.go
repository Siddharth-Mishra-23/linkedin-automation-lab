package browser

import "context"

type Browser interface {
	Start(ctx context.Context) error
	Close() error

	NewPage(ctx context.Context) (Page, error)
}

type Page interface {
	Goto(url string) error
	Click(selector string) error
	Type(selector, text string) error
	Scroll(x, y int) error
	Wait(ms int)
	HTML() (string, error)
}
