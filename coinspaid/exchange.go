package coinspaid

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type ExchangeService service

type ExchangeCalculateOptions struct {
	ReceiverAmount   *decimal.Decimal `json:"receiver_amount,omitempty"` // Amount you want to calculate for getting. The parameter is required when the "sender_amount" parameter is absent
	SenderCurrency   string           `json:"sender_currency"`           // Currency ISO for which you want to calculate the exchange rate, example: "BTC"
	ReceiverCurrency string           `json:"receiver_currency"`         // Currency ISO to be exchanged, example: "EUR"
	SenderAmount     *decimal.Decimal `json:"sender_amount,omitempty"`   // Amount you want to calculate. The parameter is required when the "receiver_amount" parameter is absent
}

type ExchangeOpts struct {
	SenderCurrency   string           `json:"sender_currency"`   // Currency ISO which you want to exchange, example: "LTC"
	ReceiverCurrency string           `json:"receiver_currency"` // Currency ISO to be exchanged, example: "USD"
	SenderAmount     decimal.Decimal  `json:"sender_amount"`     // Amount you want to exchange
	ForeignID        string           `json:"foreign_id"`        // Unique foreign ID in your system, example: "134453"
	Price            *decimal.Decimal `json:"price,omitempty"`   // Exchange rate price on which exchange will be placed
}

type Exchange struct {
	ID               int64           `json:"id"`
	ForeignID        string          `json:"foreign_id"`
	Type             string          `json:"type"`
	SenderAmount     decimal.Decimal `json:"sender_amount"`
	SenderCurrency   string          `json:"sender_currency"`
	ReceiverAmount   decimal.Decimal `json:"receiver_amount"`
	ReceiverCurrency string          `json:"receiver_currency"`
	FeeAmount        decimal.Decimal `json:"fee_amount"`
	FeeCurrency      string          `json:"fee_currency"`
	Price            decimal.Decimal `json:"price"`
	Status           string          `json:"status"`
}

type ExchangeRate struct {
	SenderAmount     decimal.Decimal `json:"sender_amount"`
	SenderCurrency   string          `json:"sender_currency"`
	ReceiverAmount   decimal.Decimal `json:"receiver_amount"`
	ReceiverCurrency string          `json:"receiver_currency"`
	FeeAmount        decimal.Decimal `json:"fee_amount"`
	FeeCurrency      string          `json:"fee_currency"`
	Price            decimal.Decimal `json:"price"`
	FixedAt          Time            `json:"ts_fixed"`
	ReleaseAt        Time            `json:"ts_release"`
	FixPeriod        Duration        `json:"fix_period"`
}

// Calculate gets info about exchange rates.
func (s *ExchangeService) Calculate(ctx context.Context, opts ExchangeCalculateOptions) (*ExchangeRate, error) {
	req, err := s.client.newRequest(http.MethodPost, "exchange/calculate", opts)
	if err != nil {
		return nil, err
	}

	var rate *ExchangeRate
	err = s.client.do(ctx, req, &rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}

// Fixed makes exchange on a given fixed exchange rate.
func (s *ExchangeService) Fixed(ctx context.Context, opts ExchangeOpts) (*Exchange, error) {
	req, err := s.client.newRequest(http.MethodPost, "exchange/fixed", opts)
	if err != nil {
		return nil, err
	}

	var ex *Exchange
	err = s.client.do(ctx, req, &ex)
	if err != nil {
		return nil, err
	}

	return ex, nil
}

// Now makes exchange without mentioning the price.
func (s *ExchangeService) Now(ctx context.Context, opts ExchangeOpts) (*Exchange, error) {
	req, err := s.client.newRequest(http.MethodPost, "exchange/now", opts)
	if err != nil {
		return nil, err
	}

	var ex *Exchange
	err = s.client.do(ctx, req, &ex)
	if err != nil {
		return nil, err
	}

	return ex, nil
}
