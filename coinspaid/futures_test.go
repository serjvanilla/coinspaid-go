package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestFuturesService_Rates(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/futures/rates", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"address":"test"}`)
		fmt.Fprint(w, `{"data":{"addresses":[{"id":384620,"currency":"BTC","convert_to":"USD","address":"3GQwSBQErsQ863RcuvCkub6SZBPHKuwSV7","tag":null,"foreign_id":"ds23fgk"}],"ts_fixed":1581507311,"ts_release":1581513311,"rates":{"BTCUSD":{"5.00000000":"10312.67153000","10.00000000":"10312.67153000","20.00000000":"10312.67153000","50.00000000":"10312.67153000","100.00000000":"10312.67153000","200.00000000":"10312.67153000","300.00000000":"10312.67153000"}}}}`)
	})

	want := &FuturesRates{
		Addresses: []*Address{
			{
				ID:        384620,
				Currency:  "BTC",
				ConvertTo: String("USD"),
				Address:   "3GQwSBQErsQ863RcuvCkub6SZBPHKuwSV7",
				ForeignID: "ds23fgk",
			},
		},
		FixedAt:   Time(time.Unix(1581507311, 0)),
		ReleaseAt: Time(time.Unix(1581513311, 0)),
		Rates: map[string]map[string]decimal.Decimal{
			"BTCUSD": {
				"5.00000000":   decimal.NewFromFloat(10312.67153),
				"10.00000000":  decimal.NewFromFloat(10312.67153),
				"20.00000000":  decimal.NewFromFloat(10312.67153),
				"50.00000000":  decimal.NewFromFloat(10312.67153),
				"100.00000000": decimal.NewFromFloat(10312.67153),
				"200.00000000": decimal.NewFromFloat(10312.67153),
				"300.00000000": decimal.NewFromFloat(10312.67153),
			},
		},
	}

	got, err := client.Futures.Rates(
		context.Background(),
		FuturesRatesOptions{
			Address: "test",
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestFuturesService_Confirm(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/futures/confirm", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"address":"test","sender_currency":"BTC","receiver_currency":"EUR","receiver_amount":"100"}`)
		fmt.Fprint(w, `{"data":{"futures_id":92,"sender_currency":"BTC","receiver_currency":"USD","fee_currency":"USD","price":"10299.94946000","address":{"id":384620,"currency":"BTC","convert_to":"USD","address":"3GQwSBQErsQ863RcuvCkub6SZBPHKuwSV7","tag":null,"foreign_id":"ds23fgk"},"sender_amount":"0.00292","receiver_amount":"30.00000000","fee_amount":"1.50000000","ts_fixed":1581506566,"ts_release":1581512566}}`)
	})

	want := &Futures{
		ID:               92,
		SenderCurrency:   "BTC",
		ReceiverCurrency: "USD",
		FeeCurrency:      "USD",
		Price:            decimal.NewFromFloat(10299.94946),
		Address: Address{
			ID:        384620,
			Currency:  "BTC",
			ConvertTo: String("USD"),
			Address:   "3GQwSBQErsQ863RcuvCkub6SZBPHKuwSV7",
			Tag:       nil,
			ForeignID: "ds23fgk",
		},
		SenderAmount:   decimal.NewFromFloat(0.00292),
		ReceiverAmount: decimal.NewFromInt(30),
		FeeAmount:      decimal.NewFromFloat(1.5),
		FixedAt:        Time(time.Unix(1581506566, 0)),
		ReleaseAt:      Time(time.Unix(1581512566, 0)),
	}

	got, err := client.Futures.Confirm(
		context.Background(),
		FuturesConfirmOptions{
			Address:          "test",
			SenderCurrency:   "BTC",
			ReceiverCurrency: "EUR",
			ReceiverAmount:   decimal.NewFromInt(100),
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestFuturesRates_RateFor(t *testing.T) {
	tests := []struct {
		amount decimal.Decimal
		price  decimal.Decimal
		error  error
	}{
		{
			amount: decimal.NewFromFloat(5),
			price:  decimal.NewFromFloat(1),
			error:  nil,
		},
		{
			amount: decimal.NewFromFloat(5.00000001),
			price:  decimal.NewFromFloat(2),
			error:  nil,
		},
		{
			amount: decimal.NewFromFloat(300.00000001),
			price:  decimal.Zero,
			error:  ErrAmountIsTooLarge,
		},
	}

	rates := &FuturesRates{
		Rates: map[string]map[string]decimal.Decimal{
			"BTCUSD": {
				"5.00000000":   decimal.NewFromFloat(1.),
				"10.00000000":  decimal.NewFromFloat(2.),
				"20.00000000":  decimal.NewFromFloat(3.),
				"50.00000000":  decimal.NewFromFloat(4.),
				"100.00000000": decimal.NewFromFloat(5.),
				"200.00000000": decimal.NewFromFloat(6.),
				"300.00000000": decimal.NewFromFloat(8.),
			},
		},
	}

	for _, test := range tests {
		rate, err := rates.RateFor("BTCUSD", test.amount)
		assert.Equal(t, err, test.error, "unexpected error")
		assert.True(t, rate.Equal(test.price), "wrong price")
	}
}
