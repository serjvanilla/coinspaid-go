// +build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestWithdrawalService_Crypto(t *testing.T) {
	c := getSandboxClient(t)

	want := &coinspaid.Withdrawal{
		ID:               0,
		ForeignID:        fmt.Sprintf("test-%d", time.Now().UnixNano()),
		Type:             "withdrawal",
		Status:           "processing",
		Amount:           decimal.NewFromFloat(0.001),
		SenderAmount:     decimal.NewFromFloat(0.001),
		SenderCurrency:   "BTC",
		ReceiverAmount:   decimal.NewFromFloat(0.001),
		ReceiverCurrency: "BTC",
	}

	opts := coinspaid.WithdrawalCryptoOptions{
		ForeignID: want.ForeignID,
		Amount:    decimal.NewFromFloat(0.001),
		Currency:  "BTC",
		Address:   "2MsbFE9tNtL7pL2iduoYTSenmcVZwcSntm5",
	}

	withdraw, err := c.Withdrawal.Crypto(context.Background(), opts)
	assert.NoError(t, err, "should complete without error")
	assert.True(t, cmp.Equal(withdraw, want, cmpopts.IgnoreFields(coinspaid.Withdrawal{}, "ID")), "should return exact result")
	assert.NotZero(t, withdraw.ID, "id should be not zero")
}
