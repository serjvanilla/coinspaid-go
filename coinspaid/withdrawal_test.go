package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestWithdrawalService_Crypto(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/withdrawal/crypto", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"foreign_id":"test","amount":"25","currency":"EUR","convert_to":"BTC","address":"moK7athCdyQBFB6SB53xPLNvUQeg1nEzzg","tag":null}`)
		fmt.Fprint(w, `{"data":{"id":2964068,"foreign_id":"test","type":"withdrawal_exchange","status":"processing","amount":"25.00000000","sender_amount":"25.00000000","sender_currency":"EUR","receiver_amount":"0.00261850","receiver_currency":"BTC"}}`)
	})

	want := &Withdrawal{
		ID:               2964068,
		ForeignID:        "test",
		Type:             "withdrawal_exchange",
		Status:           "processing",
		Amount:           decimal.NewFromInt(25),
		SenderAmount:     decimal.NewFromInt(25),
		SenderCurrency:   "EUR",
		ReceiverAmount:   decimal.NewFromFloat(0.00261850),
		ReceiverCurrency: "BTC",
	}

	got, err := client.Withdrawal.Crypto(context.Background(), WithdrawalCryptoOptions{
		ForeignID: "test",
		Amount:    decimal.NewFromInt(25),
		Currency:  "EUR",
		ConvertTo: String("BTC"),
		Address:   "moK7athCdyQBFB6SB53xPLNvUQeg1nEzzg",
		Tag:       nil,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}

func TestWithdrawalService_Crypto_InvalidTypes(t *testing.T) {
	t.Name()
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/withdrawal/crypto", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"foreign_id":"test","amount":"1","currency":"test","address":"test","tag":null}`)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"errors":{"currency":"The selected currency is invalid."}}`)
	})

	wantErr := &APIError{
		StatusCode: http.StatusBadRequest,
		Errors: map[string]string{
			"currency": "The selected currency is invalid.",
		},
	}

	resp, err := client.Withdrawal.Crypto(context.Background(), WithdrawalCryptoOptions{
		ForeignID: "test",
		Amount:    decimal.NewFromInt(1),
		Currency:  "test",
		Address:   "test",
	})
	if resp != nil {
		t.Errorf("unexpected response: %v", resp)
	} else if diff := cmp.Diff(wantErr, err); diff != "" {
		t.Errorf("error mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
