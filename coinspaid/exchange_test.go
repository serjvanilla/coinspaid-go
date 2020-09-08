package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/shopspring/decimal"
)

func TestExchangeService_Calculate(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/exchange/calculate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"sender_currency":"BTC","receiver_currency":"EUR"}`)
		fmt.Fprint(w, `{"data":{"sender_amount":1,"sender_currency":"BTC","receiver_amount":"8549.81680000","receiver_currency":"EUR","fee_amount":"85.49816800","fee_currency":"BTC","price":"8549.81680000","ts_fixed":1564293159,"ts_release":1564293219,"fix_period":60}}`)
	})

	want := &ExchangeRate{
		SenderAmount:     decimal.NewFromInt(1),
		SenderCurrency:   "BTC",
		ReceiverAmount:   decimal.RequireFromString("8549.81680000"),
		ReceiverCurrency: "EUR",
		FeeAmount:        decimal.RequireFromString("85.49816800"),
		FeeCurrency:      "BTC",
		Price:            decimal.RequireFromString("8549.81680000"),
		FixedAt:          Time(time.Unix(1564293159, 0)),
		ReleaseAt:        Time(time.Unix(1564293219, 0)),
		FixPeriod:        Duration(60 * time.Second),
	}

	got, err := client.Exchange.Calculate(context.Background(), ExchangeCalculateOptions{SenderCurrency: "BTC", ReceiverCurrency: "EUR"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestExchangeService_Fixed(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/exchange/fixed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"sender_currency":"LTC","receiver_currency":"USD","sender_amount":"6.5","foreign_id":"134453","price":"89.75202"}`)
		fmt.Fprint(w, `{"data":{"id":2687667,"foreign_id":"knwi24op9","type":"exchange","sender_amount":"0.01","sender_currency":"BTC","receiver_amount":"63.52069015","receiver_currency":"EUR","fee_amount":"6.98727592","fee_currency":"EUR","price":"6352.06901520","status":"processing"}}`)
	})

	want := &Exchange{
		ID:               2687667,
		ForeignID:        "knwi24op9",
		Type:             "exchange",
		SenderAmount:     decimal.NewFromFloat(0.01),
		SenderCurrency:   "BTC",
		ReceiverAmount:   decimal.RequireFromString("63.52069015"),
		ReceiverCurrency: "EUR",
		FeeAmount:        decimal.RequireFromString("6.98727592"),
		FeeCurrency:      "EUR",
		Price:            decimal.RequireFromString("6352.06901520"),
		Status:           "processing",
	}

	got, err := client.Exchange.Fixed(
		context.Background(),
		ExchangeOpts{
			SenderCurrency:   "LTC",
			ReceiverCurrency: "USD",
			SenderAmount:     decimal.NewFromFloat(6.5),
			ForeignID:        "134453",
			Price:            Decimal(decimal.NewFromFloat(89.75202)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestExchangeService_Now(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/exchange/now", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"sender_currency":"EUR","receiver_currency":"BTC","sender_amount":"2","foreign_id":"124876"}`)
		fmt.Fprint(w, `{"data":{"id":2687669,"foreign_id":"wph27bmsp81","type":"exchange","sender_amount":"0.001","sender_currency":"BTC","receiver_amount":"6.33984642","receiver_currency":"EUR","fee_amount":"0.69738311","fee_currency":"EUR","price":"6339.84642127","status":"processing"}}`)
	})

	want := &Exchange{
		ID:               2687669,
		ForeignID:        "wph27bmsp81",
		Type:             "exchange",
		SenderAmount:     decimal.NewFromFloat(0.001),
		SenderCurrency:   "BTC",
		ReceiverAmount:   decimal.RequireFromString("6.33984642"),
		ReceiverCurrency: "EUR",
		FeeAmount:        decimal.RequireFromString("0.69738311"),
		FeeCurrency:      "EUR",
		Price:            decimal.RequireFromString("6339.84642127"),
		Status:           "processing",
	}

	got, err := client.Exchange.Now(
		context.Background(),
		ExchangeOpts{
			SenderCurrency:   "EUR",
			ReceiverCurrency: "BTC",
			SenderAmount:     decimal.NewFromInt(2),
			ForeignID:        "124876",
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
