package browser

import (
	"context"
	"fmt"
	"time"
)

type MockBrowser struct{}

func (m *MockBrowser) Start(ctx context.Context) error {
	fmt.Println("[Browser] Starting mock browser")
	return nil
}

func (m *MockBrowser) Close() error {
	fmt.Println("[Browser] Closing mock browser")
	return nil
}

func (m *MockBrowser) NewPage(ctx context.Context) (Page, error) {
	fmt.Println("[Browser] Opening new mock page")
	return &MockPage{}, nil
}

// MockPage simulates page actions.
type MockPage struct{}

func (p *MockPage) Goto(url string) error {
	fmt.Println("[Page] Navigating to:", url)
	return nil
}

func (p *MockPage) Click(selector string) error {
	fmt.Println("[Page] Clicking element:", selector)
	return nil
}

func (p *MockPage) Type(selector, text string) error {
	fmt.Printf("[Page] Typing into %s: %s\n", selector, text)
	return nil
}

func (p *MockPage) Scroll(x, y int) error {
	fmt.Printf("[Page] Scrolling to (%d, %d)\n", x, y)
	return nil
}

func (p *MockPage) Wait(ms int) {
	fmt.Printf("[Page] Waiting for %d ms\n", ms)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (p *MockPage) HTML() (string, error) {
	return "<html><body>Mock Page</body></html>", nil
}
