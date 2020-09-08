package coinspaid

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type AccountsService service

type Account struct {
	Currency string          `json:"currency"`
	Type     CurrencyType    `json:"type"`
	Balance  decimal.Decimal `json:"balance"`
}

// List of all the balances (including zero balances).
func (s *AccountsService) List(ctx context.Context) ([]*Account, error) {
	req, err := s.client.newRequest(http.MethodPost, "accounts/list", nil)
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	err = s.client.do(ctx, req, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
