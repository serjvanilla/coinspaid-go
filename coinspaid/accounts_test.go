package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/google/go-cmp/cmp"
)

func TestAccountsService_List(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/accounts/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		fmt.Fprint(w, `{"data":[{"currency":"RUB","type":"fiat","balance":"2699794.29053025"},{"currency":"DOGE","type":"crypto","balance":"1656013.69744328"}]}`)
	})

	want := []*Account{
		{Currency: "RUB", Type: CurrencyTypeFiat, Balance: decimal.RequireFromString("2699794.29053025")},
		{Currency: "DOGE", Type: CurrencyTypeCrypto, Balance: decimal.RequireFromString("1656013.69744328")},
	}

	got, err := client.Accounts.List(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
