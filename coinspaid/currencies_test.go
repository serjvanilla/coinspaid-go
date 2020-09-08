package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestCurrenciesService_List(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/currencies/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		fmt.Fprint(w, `{"data":[{"id":1,"type":"crypto","currency":"BTC","minimum_amount":"0.00030000","deposit_fee_percent":"0.010000","withdrawal_fee_percent":"0.010000","precision":8},{"id":20,"type":"fiat","currency":"USD","minimum_amount":"0.00000000","deposit_fee_percent":"0.000000","withdrawal_fee_percent":"0.000000","precision":8}]}`)
	})

	want := []*Currency{
		{
			ID:                   1,
			Type:                 CurrencyTypeCrypto,
			Currency:             "BTC",
			MinimumAmount:        decimal.NewFromFloat(0.0003),
			DepositFeePercent:    decimal.NewFromFloat(0.01),
			WithdrawalFeePercent: decimal.NewFromFloat(0.01),
			Precision:            8,
		},
		{
			ID:                   20,
			Type:                 CurrencyTypeFiat,
			Currency:             "USD",
			MinimumAmount:        decimal.Zero,
			DepositFeePercent:    decimal.Zero,
			WithdrawalFeePercent: decimal.Zero,
			Precision:            8,
		},
	}

	got, err := client.Currencies.List(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestCurrenciesService_Pairs(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/currencies/pairs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"currency_from":"BTC","currency_to":"EUR"}`)
		fmt.Fprint(w, `{"data":[{"currency_from":{"currency":"BTC","type":"crypto","min_amount":"0.00300000","min_amount_deposit_with_exchange":"0.00030000"},"currency_to":{"currency":"EUR","type":"fiat"},"rate_from":"1","rate_to":"9484.70016880"}]}`)
	})

	want := []*CurrenciesPair{
		{
			CurrencyFrom: &Currency{Type: CurrencyTypeCrypto, Currency: "BTC"},
			CurrencyTo:   &Currency{Type: CurrencyTypeFiat, Currency: "EUR"},
			RateFrom:     decimal.NewFromInt(1),
			RateTo:       decimal.RequireFromString("9484.70016880"),
		},
	}

	got, err := client.Currencies.Pairs(context.Background(), &CurrenciesPairsOptions{CurrencyFrom: String("BTC"), CurrencyTo: String("EUR")})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
