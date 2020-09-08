package coinspaid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetTerminalLink(t *testing.T) {
	want := `https://terminal.coinspaid.com/?data=eyJjbGllbnRfaWQiOjEsImN1cnJlbmN5IjoiQlRDIiwiYW1vdW50IjowLjAwMDUsImNvbnZlcnRfdG8iOiJFVVIiLCJmb3JlaWduX2lkIjoiMTU4ODE0NjQ4MDE1ODgxNDY0ODAiLCJ1cmxfYmFjayI6Imh0dHBzOi8vY29pbnNwYWlkLmNvbS8_MTU4ODE0NjQ4MCJ9&signature=dc6b63b63f4708953449286ea73ed8d0892675c4b5698cce02a4c3887114bb312eb496efdd38ae1457421cc637b9e07f651b32a6e7592a74c0e10b9b8e86b732`

	client, _, _, cleanup := setupTest()
	defer cleanup()

	opts := TerminalOptions{
		ClientID:  1,
		Currency:  "BTC",
		Amount:    0.0005,
		ConvertTo: String("EUR"),
		ForeignID: "15881464801588146480",
		URLBack:   "https://coinspaid.com/?1588146480",
	}

	link := client.GetTerminalLink(opts)

	assert.Equal(t, want, link)
}
