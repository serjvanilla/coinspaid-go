// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/google/go-cmp/cmp"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"
)

func TestAddressesService_List(t *testing.T) {
	want := []*coinspaid.Account{
		{Currency: "USD", Type: "fiat", Balance: decimal.RequireFromString("10.81135268")},
		{Currency: "BTC", Type: "crypto", Balance: decimal.RequireFromString("0.00994876")},
		{Currency: "ETH", Type: "crypto", Balance: decimal.Zero},
		{Currency: "EUR", Type: "fiat", Balance: decimal.Zero},
		{Currency: "USDTE", Type: "crypto", Balance: decimal.Zero},
		{Currency: "XRP", Type: "crypto", Balance: decimal.Zero},
		{Currency: "CNY", Type: "fiat", Balance: decimal.Zero},
		{Currency: "EURS", Type: "crypto", Balance: decimal.Zero},
		{Currency: "AUD", Type: "fiat", Balance: decimal.Zero},
		{Currency: "GBP", Type: "fiat", Balance: decimal.Zero},
		{Currency: "CAD", Type: "fiat", Balance: decimal.Zero},
		{Currency: "SEK", Type: "fiat", Balance: decimal.Zero},
		{Currency: "NOK", Type: "fiat", Balance: decimal.Zero},
		{Currency: "CHF", Type: "fiat", Balance: decimal.Zero},
		{Currency: "RUB", Type: "fiat", Balance: decimal.Zero},
		{Currency: "JPY", Type: "fiat", Balance: decimal.Zero},
		{Currency: "NZD", Type: "fiat", Balance: decimal.Zero},
		{Currency: "MXN", Type: "fiat", Balance: decimal.Zero},
		{Currency: "ARS", Type: "fiat", Balance: decimal.Zero},
		{Currency: "BRL", Type: "fiat", Balance: decimal.Zero},
		{Currency: "INR", Type: "fiat", Balance: decimal.Zero},
		{Currency: "KRW", Type: "fiat", Balance: decimal.Zero},
		{Currency: "MYR", Type: "fiat", Balance: decimal.Zero},
		{Currency: "THB", Type: "fiat", Balance: decimal.Zero},
		{Currency: "IDR", Type: "fiat", Balance: decimal.Zero},
		{Currency: "VND", Type: "fiat", Balance: decimal.Zero},
		{Currency: "PEN", Type: "fiat", Balance: decimal.Zero},
		{Currency: "CLP", Type: "fiat", Balance: decimal.Zero},
		{Currency: "KZT", Type: "fiat", Balance: decimal.Zero},
		{Currency: "UAH", Type: "fiat", Balance: decimal.Zero},
		{Currency: "CZK", Type: "fiat", Balance: decimal.Zero},
		{Currency: "PLN", Type: "fiat", Balance: decimal.Zero},
		{Currency: "ZAR", Type: "fiat", Balance: decimal.Zero},
	}

	c := getSandboxClient(t)

	accounts, err := c.Accounts.List(context.Background())
	assert.NoError(t, err, "should complete without error")
	assert.True(t, cmp.Equal(accounts, want), "should return exact result")
}
