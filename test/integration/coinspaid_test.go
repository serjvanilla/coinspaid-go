// +build integration

package integration

import (
	"context"
	"os"
	"testing"

	"github.com/serjvanilla/coinspaid-go/coinspaid"
)

const (
	msgEnvMissing         = "Skipping test because the required environment variable (%v) is not present."
	envKeyCoinsPaidKey    = `COINSPAID_KEY`
	envKeyCoinsPaidSecret = `COINSPAID_SECRET`
)

func getSandboxClient(t *testing.T) *coinspaid.Client {
	key, ok := os.LookupEnv(envKeyCoinsPaidKey)
	if !ok {
		t.Skipf(msgEnvMissing, envKeyCoinsPaidKey)
	}

	secret, ok := os.LookupEnv(envKeyCoinsPaidSecret)
	if !ok {
		t.Skipf(msgEnvMissing, envKeyCoinsPaidSecret)
	}

	return coinspaid.NewClient(key, secret, coinspaid.WithSandbox())
}

func TestClient_Ping(t *testing.T) {
	c := getSandboxClient(t)

	if pong, err := c.Ping(context.Background()); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if pong != "OK" {
		t.Errorf("ping is not ok: %v", pong)
	}
}
