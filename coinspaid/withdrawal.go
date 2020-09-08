package coinspaid

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type WithdrawalService service

type WithdrawalCryptoOptions struct {
	ForeignID string          `json:"foreign_id"`           // Unique foreign ID in your system, example: "122929"
	Amount    decimal.Decimal `json:"amount"`               // Amount of funds to withdraw, example: "3500"
	Currency  string          `json:"currency"`             // Currency ISO to be withdrawn, example: "EUR"
	ConvertTo *string         `json:"convert_to,omitempty"` // If you want to auto convert for example EUR to BTC, specify this param as BTC or any other currency supported (see list of exchangeable pairs API method).
	Address   string          `json:"address"`              // Cryptocurrency address where you want to send funds.
	Tag       *string         `json:"tag"`                  // Tag (if it's Ripple or BNB) or memo (if it's Bitshares or EOS)
}

// Crypto creates withdraw in crypto to any specified address. You can send Cryptocurrency from your Fiat currency balance by using "convert_to" parameter.
func (s *WithdrawalService) Crypto(ctx context.Context, opts WithdrawalCryptoOptions) (*Withdrawal, error) {
	req, err := s.client.newRequest(http.MethodPost, "withdrawal/crypto", opts)
	if err != nil {
		return nil, err
	}

	var transaction *Withdrawal
	err = s.client.do(ctx, req, &transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

type Withdrawal struct {
	ID               int64           `json:"id"`
	ForeignID        string          `json:"foreign_id"`
	Type             string          `json:"type"`
	Status           string          `json:"status"`
	Amount           decimal.Decimal `json:"amount"`
	SenderAmount     decimal.Decimal `json:"sender_amount"`
	SenderCurrency   string          `json:"sender_currency"`
	ReceiverAmount   decimal.Decimal `json:"receiver_amount"`
	ReceiverCurrency string          `json:"receiver_currency"`
}
