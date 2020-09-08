// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"
)

func TestExchangeService_Rates(t *testing.T) {
	c := getSandboxClient(t)

	rate, err := c.Exchange.Calculate(
		context.Background(),
		coinspaid.ExchangeCalculateOptions{
			SenderCurrency:   "BTC",
			ReceiverCurrency: "USD",
			ReceiverAmount:   coinspaid.Decimal(decimal.NewFromInt(100)),
		},
	)

	assert.NoError(t, err, "should complete without error")
	assert.True(t, rate.SenderAmount.IsPositive(), "sender amount should be greater than zero")
	assert.Equal(t, rate.SenderCurrency, "BTC", "sender currency should be equal BTC")
	assert.True(t, rate.ReceiverAmount.Equal(decimal.NewFromInt(100)), "sender amount should be equal 100")
	assert.Equal(t, rate.ReceiverCurrency, "USD", "receiver currency should be equal USD")
	assert.True(t, rate.FeeAmount.IsPositive(), "fee amount should be greater than zero")
	assert.Equal(t, rate.FeeCurrency, "USD", "fee currency should be equal USD")
	assert.True(t, rate.Price.IsPositive(), "price should be greater than zero")
	assert.WithinDuration(t, rate.FixedAt.Time(), time.Now(), time.Minute, "should be fixed within a minute")
	assert.Equal(t, rate.FixPeriod.Duration(), 60*time.Second, "should be fixed for 1 minute")
	assert.Equal(t, rate.ReleaseAt.Time(), rate.FixedAt.Time().Add(rate.FixPeriod.Duration()), "should be released after 1 minute")
}
