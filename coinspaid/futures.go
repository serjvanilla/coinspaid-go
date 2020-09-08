package coinspaid

import (
	"context"
	"errors"
	"net/http"
	"sort"

	"github.com/shopspring/decimal"
)

var (
	ErrUnknownPair      = errors.New("unknown pair")
	ErrAmountIsTooLarge = errors.New("amount is too large")
)

type FuturesService service

type FuturesRatesOptions struct {
	Address string `json:"address"` // Exchange address for which you want to calculate futures' rates
}

type FuturesConfirmOptions struct {
	Address          string          `json:"address"`           // Exchange address for which you want to confirm futures
	SenderCurrency   string          `json:"sender_currency"`   // Currency ISO which you want to exchange, example: "BTC"
	ReceiverCurrency string          `json:"receiver_currency"` // Currency ISO to be exchanged, example: "EUR"
	ReceiverAmount   decimal.Decimal `json:"receiver_amount"`   // Amount you want to receive
}

type FuturesRates struct {
	Addresses []*Address                            `json:"addresses"`
	FixedAt   Time                                  `json:"ts_fixed"`
	ReleaseAt Time                                  `json:"ts_release"`
	Rates     map[string]map[string]decimal.Decimal `json:"rates"`
}

// RateFor is helper method for calculating rate for queried amount.
func (r *FuturesRates) RateFor(pair string, amount decimal.Decimal) (decimal.Decimal, error) {
	rawRates, ok := r.Rates[pair]
	if !ok {
		return decimal.Zero, ErrUnknownPair
	}

	type rate struct {
		amount decimal.Decimal
		rate   decimal.Decimal
	}
	rates := make([]rate, 0, len(rawRates))

	for k, v := range rawRates {
		a, err := decimal.NewFromString(k)
		if err != nil {
			return decimal.Zero, err
		}
		rates = append(rates, rate{amount: a, rate: v})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].amount.LessThan(rates[j].amount)
	})

	for _, rate := range rates {
		if amount.LessThanOrEqual(rate.amount) {
			return rate.rate, nil
		}
	}

	return decimal.Zero, ErrAmountIsTooLarge
}

type Futures struct {
	ID               int64           `json:"futures_id"`
	SenderCurrency   string          `json:"sender_currency"`
	ReceiverCurrency string          `json:"receiver_currency"`
	FeeCurrency      string          `json:"fee_currency"`
	Price            decimal.Decimal `json:"price"`
	Address          Address         `json:"address"`
	SenderAmount     decimal.Decimal `json:"sender_amount"`
	ReceiverAmount   decimal.Decimal `json:"receiver_amount"`
	FeeAmount        decimal.Decimal `json:"fee_amount"`
	FixedAt          Time            `json:"ts_fixed"`
	ReleaseAt        Time            `json:"ts_release"`
}

// Rates gets info about rates for futures.
func (s *FuturesService) Rates(ctx context.Context, opts FuturesRatesOptions) (*FuturesRates, error) {
	req, err := s.client.newRequest(http.MethodPost, "futures/rates", opts)
	if err != nil {
		return nil, err
	}

	var rate *FuturesRates
	err = s.client.do(ctx, req, &rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}

// Confirm futures transaction.
func (s *FuturesService) Confirm(ctx context.Context, opts FuturesConfirmOptions) (*Futures, error) {
	req, err := s.client.newRequest(http.MethodPost, "futures/confirm", opts)
	if err != nil {
		return nil, err
	}

	var futures *Futures
	err = s.client.do(ctx, req, &futures)
	if err != nil {
		return nil, err
	}

	return futures, nil
}
