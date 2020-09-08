package coinspaid_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/serjvanilla/coinspaid-go/coinspaid"
)

const (
	testKey    = `test_key`
	testSecret = `test_secret`
)

func TestClient_CallbackHandler(t *testing.T) {
	c := coinspaid.NewClient(testKey, testSecret)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader("{\"id\":3576,\"type\":\"deposit_exchange\",\"crypto_address\":{\"id\":7419,\"currency\":\"BTC\",\"convert_to\":\"USD\",\"address\":\"2N1QTWoPZ6Z7rQue3KS5YAMUQhKabrN4nHn\",\"tag\":null,\"foreign_id\":\"test\"},\"currency_sent\":{\"currency\":\"BTC\",\"amount\":\"0.00100000\"},\"currency_received\":{\"currency\":\"USD\",\"amount\":\"11.40406820\",\"amount_minus_fee\":\"10.81135268\"},\"transactions\":[{\"id\":4418,\"currency\":\"BTC\",\"transaction_type\":\"blockchain\",\"type\":\"deposit\",\"address\":\"2N1QTWoPZ6Z7rQue3KS5YAMUQhKabrN4nHn\",\"tag\":null,\"amount\":\"0.00100000\",\"txid\":\"daeb0cfb14a2c6d425906f75f14fb4b6b55221a0586a0c890436dd1a9bd9005d\",\"riskscore\":null,\"confirmations\":\"1\"},{\"id\":4419,\"currency\":\"BTC\",\"currency_to\":\"USD\",\"transaction_type\":\"exchange\",\"type\":\"exchange\",\"amount\":\"0.00100000\",\"amount_to\":\"11.40406820\"}],\"fees\":[{\"type\":\"fee_crypto_deposit_to_fiat\",\"currency\":\"USD\",\"amount\":\"0.59271552\"}],\"error\":\"\",\"status\":\"confirmed\"}"))
	req.Header.Add("X-Processing-Key", testKey)
	req.Header.Add("X-Processing-Signature", "3a031e16649f5120e4f08d462578681f8b6842815dbf5fc1e209445585cf91957b0555bc68ef48d95cee4349441404511198a0a5accd8aac396c029b4665c88b")

	handler := c.CallbackHandler(func(ctx context.Context, data *coinspaid.CallbackData) bool {
		if data == nil {
			t.Error("nil data passed to handler")
			return false
		}

		const wantID = 3576

		if data.ID != wantID {
			t.Errorf("wrong callback id: got %v want %v", data.ID, wantID)
		}

		return true
	})

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
