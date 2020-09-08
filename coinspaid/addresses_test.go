package coinspaid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddressesService_Take(t *testing.T) {
	client, mux, _, cleanup := setupTest()

	mux.HandleFunc("/addresses/take", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testHeaders(t, r, client.key, client.secret)
		testBody(t, r, `{"foreign_id":"test","currency":"BTC","convert_to":"USD"}`)
		fmt.Fprint(w, `{"data":{"id":387552,"currency":"BTC","convert_to":"USD","address":"2MwToDP2jFAwQxFuUSFzTH9FM8kMDWQCpRt","tag":null,"foreign_id":"test"}}`)
	})

	want := &Address{
		ID:        387552,
		Currency:  "BTC",
		ConvertTo: String("USD"),
		Address:   "2MwToDP2jFAwQxFuUSFzTH9FM8kMDWQCpRt",
		Tag:       nil,
		ForeignID: "test",
	}

	got, err := client.Addresses.Take(context.Background(), AddressTakeOptions{
		ForeignID: "test",
		Currency:  "BTC",
		ConvertTo: String("USD"),
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("response mismatch (-want +got):\n%s", diff)
	}

	cleanup()
}
