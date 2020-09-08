package coinspaid

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

func (c *Client) CallbackHandler(f func(ctx context.Context, data *CallbackData) (ok bool)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(headerProcessingKey) != c.key {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		rawData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if genSignature(c.secret, rawData) != r.Header.Get(headerProcessingSignature) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var data *CallbackData
		err = json.Unmarshal(rawData, &data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}

		if f(r.Context(), data) {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
	})
}

type CallbackType string

const (
	CallbackTypeDeposit    CallbackType = `deposit`
	CallbackTypeWithdrawal CallbackType = `withdrawal`
	CallbackTypeExchange   CallbackType = `exchange`
	CallbackTypeFutures    CallbackType = `deposit_exchange`
	CallbackTypeInvoice    CallbackType = `invoice`
)

type CallbackStatus string

const (
	TransactionsStatusConfirmed    CallbackStatus = `confirmed`
	TransactionsStatusNotConfirmed CallbackStatus = `not_confirmed`
	TransactionsStatusCancelled    CallbackStatus = `cancelled`
)

type CallbackData struct {
	ID               int64           `json:"id"`
	ForeignID        string          `json:"foreign_id,omitempty"`
	Type             CallbackType    `json:"type"`
	CryptoAddress    *Address        `json:"crypto_address,omitempty"`
	CurrencySent     *CurrencyAmount `json:"currency_sent,omitempty"`
	CurrencyReceived *CurrencyAmount `json:"currency_received,omitempty"`
	Transactions     []Transaction   `json:"transactions"`
	Fees             []Fee           `json:"fees"`
	Error            string          `json:"error"`
	Status           CallbackStatus  `json:"status"`
	FuturesID        int64           `json:"futures_id,omitempty"`
	TransactionID    int64           `json:"transaction_id,omitempty"`
	FixedAt          *Time           `json:"fixed_at,omitempty"`
	ExpiresAt        *Time           `json:"expires_at,omitempty"`
}

type CurrencyAmount struct {
	Currency        string           `json:"currency"`
	Amount          decimal.Decimal  `json:"amount"`
	AmountMinusFee  *decimal.Decimal `json:"amount_minus_fee,omitempty"`
	RemainingAmount *decimal.Decimal `json:"remaining_amount,omitempty"`
}

type Transaction struct {
	ID              int64            `json:"id"`
	Currency        string           `json:"currency"`
	CurrencyTo      string           `json:"currency_to,omitempty"`
	TransactionType string           `json:"transaction_type"`
	Type            string           `json:"type"`
	Address         string           `json:"address,omitempty"`
	Tag             *string          `json:"tag"`
	Amount          decimal.Decimal  `json:"amount"`
	AmountTo        *decimal.Decimal `json:"amount_to,omitempty"`
	TxID            *string          `json:"txid"`
	RiskScore       *decimal.Decimal `json:"riskscore,omitempty"`
	Confirmations   int              `json:"confirmations,string"`
}

type Fee struct {
	Type     string          `json:"type"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
}
