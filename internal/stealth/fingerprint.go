package stealth

import (
	"context"
	"fmt"
)

type FingerprintPlugin struct{}

func (f *FingerprintPlugin) Name() string {
	return "Browser Fingerprint Masking"
}

func (f *FingerprintPlugin) Apply(ctx context.Context) error {
	fmt.Println("[Stealth][Fingerprint] Randomizing user agent and viewport")
	return nil
}
