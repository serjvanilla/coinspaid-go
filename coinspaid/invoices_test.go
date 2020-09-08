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

func TestInvoicesService_Create(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/invoices/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"timer":true,"title":"Test invoice","description":"some description","currency":"EUR","sender_currency":"BTC","amount":"100","foreign_id":"test","url_success":"https://example.com/success","url_failed":"https://example.com/fail","email_user":"user@address.com"}`)
		fmt.Fprint(w, `{"data":{"id":79,"url":"https://app.coinspaid.com/invoice/RB9NZv","foreign_id":"164","name":"TESTNAME","status":"created","currency":"EUR","amount":"106.75","sender_currency":"BTC","sender_amount":null,"fixed_at":1581929889,"release_at":1581930789}}`)
	})

	want := &Invoice{
		ID:             79,
		URL:            "https://app.coinspaid.com/invoice/RB9NZv",
		ForeignID:      "164",
		Name:           "TESTNAME",
		Status:         "created",
		Currency:       "EUR",
		Amount:         decimal.NewFromFloat(106.75),
		SenderCurrency: "BTC",
		SenderAmount:   nil,
		FixedAt:        Time(time.Unix(1581929889, 0)),
		ReleaseAt:      Time(time.Unix(1581930789, 0)),
	}

	got, err := client.Invoices.Create(
		context.Background(),
		InvoicesCreateOptions{
			Timer:          true,
			Title:          "Test invoice",
			Description:    "some description",
			Currency:       "EUR",
			SenderCurrency: String("BTC"),
			Amount:         decimal.NewFromInt(100),
			ForeignID:      "test",
			URLSuccess:     "https://example.com/success",
			URLFailed:      "https://example.com/fail",
			EmailUser:      "user@address.com",
		},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
