// +build integration

package integration

import (
	"context"
	"testing"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestCurrenciesService_List(t *testing.T) {
	want := []*coinspaid.Currency{
		{
			ID:                   1,
			Type:                 coinspaid.CurrencyTypeCrypto,
			Currency:             "BTC",
			MinimumAmount:        decimal.NewFromFloat(0.0001),
			DepositFeePercent:    decimal.NewFromFloat(0.008),
			WithdrawalFeePercent: decimal.Zero,
			Precision:            8,
		},
		{
			ID:                   2,
			Type:                 coinspaid.CurrencyTypeCrypto,
			Currency:             "ETH",
			MinimumAmount:        decimal.NewFromFloat(0.01),
			DepositFeePercent:    decimal.NewFromFloat(0.008),
			WithdrawalFeePercent: decimal.Zero,
			Precision:            8,
		},
		{
			ID:                   3,
			Type:                 coinspaid.CurrencyTypeFiat,
			Currency:             "EUR",
			MinimumAmount:        decimal.Zero,
			DepositFeePercent:    decimal.Zero,
			WithdrawalFeePercent: decimal.Zero,
			Precision:            8,
		},
		{
			ID:                   4,
			Type:                 coinspaid.CurrencyTypeFiat,
			Currency:             "USD",
			MinimumAmount:        decimal.Zero,
			DepositFeePercent:    decimal.Zero,
			WithdrawalFeePercent: decimal.Zero,
			Precision:            8,
		},
	}

	c := getSandboxClient(t)

	currencies, err := c.Currencies.List(context.Background())
	assert.NoError(t, err, "should complete without error")
	assert.True(t, cmp.Equal(currencies, want), "should return exact result")
}

func TestCurrenciesService_Pairs(t *testing.T) {
	c := getSandboxClient(t)

	pairs, err := c.Currencies.Pairs(context.Background(), nil)

	assert.NoError(t, err, "should complete without error")
	assert.Len(t, pairs, 10, "should return 10 currencies pairs")
	for _, pair := range pairs {
		assert.NotEmpty(t, pair.CurrencyFrom.Currency, "from currency name should be present")
		assert.NotEmpty(t, pair.CurrencyFrom.Type, "from currency type should be present")
		assert.NotEmpty(t, pair.CurrencyTo.Currency, "to currency name should be present")
		assert.NotEmpty(t, pair.CurrencyTo.Type, "to currency type should be present")
		assert.True(t, pair.RateFrom.IsPositive(), "rate from should be greater than zero")
		assert.True(t, pair.RateTo.IsPositive(), "rate from should be greater than zero")
	}
}

func TestCurrenciesService_PairsFiltered(t *testing.T) {
	c := getSandboxClient(t)
	opts := &coinspaid.CurrenciesPairsOptions{CurrencyFrom: coinspaid.String("USD")}
	pairs, err := c.Currencies.Pairs(context.Background(), opts)

	assert.NoError(t, err, "should complete without error")
	assert.Len(t, pairs, 2, "should return 2 currencies pairs")
	for _, pair := range pairs {
		assert.Equal(t, pair.CurrencyFrom.Currency, "USD", "from currency name should be USD")
		assert.Equal(t, pair.CurrencyFrom.Type, coinspaid.CurrencyTypeFiat, "from currency type should be present")
		assert.NotEmpty(t, pair.CurrencyTo.Currency, "to currency name should be present")
		assert.NotEmpty(t, pair.CurrencyTo.Type, "to currency type should be present")
		assert.True(t, pair.RateFrom.IsPositive(), "rate from should be greater than zero")
		assert.True(t, pair.RateTo.IsPositive(), "rate from should be greater than zero")
	}
}
