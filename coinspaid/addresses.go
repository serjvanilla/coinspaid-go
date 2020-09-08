package coinspaid

import (
	"context"
	"net/http"
)

type AddressesService service

type Address struct {
	ID        int64   `json:"id"`
	Currency  string  `json:"currency"`
	ConvertTo *string `json:"convert_to,omitempty"`
	Address   string  `json:"address"`
	Tag       *string `json:"tag"`
	ForeignID string  `json:"foreign_id,omitempty"`
}

type AddressTakeOptions struct {
	ForeignID string  `json:"foreign_id"`           // Your info for this address, will returned as reference in Address responses, example: user-id:2048
	Currency  string  `json:"currency"`             // ISO of currency to receive funds in, example: BTC
	ConvertTo *string `json:"convert_to,omitempty"` // If you need auto exchange all incoming funds for example to EUR, specify this param as EUR or any other supported currency ISO, to see list of pairs see previous method.
}

// Take address for depositing crypto and (it depends on specified params) exchange from crypto to fiat on-the-fly.
func (s *AddressesService) Take(ctx context.Context, opts AddressTakeOptions) (*Address, error) {
	req, err := s.client.newRequest(http.MethodPost, "addresses/take", opts)
	if err != nil {
		return nil, err
	}

	var address *Address
	err = s.client.do(ctx, req, &address)
	if err != nil {
		return nil, err
	}

	return address, nil
}
