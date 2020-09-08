package coinspaid

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type CurrencyType string

const (
	CurrencyTypeCrypto CurrencyType = `crypto`
	CurrencyTypeFiat   CurrencyType = `fiat`
)

type CurrenciesService service

type Currency struct {
	ID                   int64           `json:"id"`
	Type                 CurrencyType    `json:"type"`
	Currency             string          `json:"currency"`
	MinimumAmount        decimal.Decimal `json:"minimum_amount"`
	DepositFeePercent    decimal.Decimal `json:"deposit_fee_percent"`
	WithdrawalFeePercent decimal.Decimal `json:"withdrawal_fee_percent"`
	Precision            int             `json:"precision"`
}

type CurrenciesPair struct {
	CurrencyFrom *Currency       `json:"currency_from"`
	CurrencyTo   *Currency       `json:"currency_to"`
	RateFrom     decimal.Decimal `json:"rate_from"`
	RateTo       decimal.Decimal `json:"rate_to"`
}

type CurrenciesPairsOptions struct {
	CurrencyFrom *string `json:"currency_from,omitempty"` // Filter by currency ISO that exchanges from, example: BTC
	CurrencyTo   *string `json:"currency_to,omitempty"`   // Filter by currency ISO that can be converted to, example: EUR
}

// List gets all supported currencies.
func (s *CurrenciesService) List(ctx context.Context) ([]*Currency, error) {
	req, err := s.client.newRequest(http.MethodPost, "currencies/list", nil)
	if err != nil {
		return nil, err
	}

	var currencies []*Currency
	err = s.client.do(ctx, req, &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

// Pairs returns list of currency pairs if no parameters passed. Returns particular pair and its price if currency parameters are passed.
func (s *CurrenciesService) Pairs(ctx context.Context, opts *CurrenciesPairsOptions) ([]*CurrenciesPair, error) {
	req, err := s.client.newRequest(http.MethodPost, "currencies/pairs", opts)
	if err != nil {
		return nil, err
	}

	var pairs []*CurrenciesPair
	err = s.client.do(ctx, req, &pairs)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}
