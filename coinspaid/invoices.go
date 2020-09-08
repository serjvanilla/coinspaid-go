package coinspaid

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
)

type InvoicesService service

type Invoice struct {
	ID             int64            `json:"id"`
	URL            string           `json:"url"`
	ForeignID      string           `json:"foreign_id"`
	Name           string           `json:"name"`
	Status         string           `json:"status"`
	Currency       string           `json:"currency"`
	Amount         decimal.Decimal  `json:"amount"`
	SenderCurrency string           `json:"sender_currency"`
	SenderAmount   *decimal.Decimal `json:"sender_amount"`
	FixedAt        Time             `json:"fixed_at"`
	ReleaseAt      Time             `json:"release_at"`
}

type InvoicesCreateOptions struct {
	Timer          bool            `json:"timer"`                 // Time on the rate is fixed for invoice payment (15 minutes). During this time the user has to pay an invoice.
	Title          string          `json:"title"`                 // Invoice title that will be displayed to the user
	Description    string          `json:"description,omitempty"` // Invoice description that will be displayed to the user
	Currency       string          `json:"currency"`              // ISO invoice currency that you want to receive from the user, for example: "EUR”
	SenderCurrency *string         `json:"sender_currency"`       // Currency of user invoice payment (3rd type invoice will be externalized at the time of sending this parameter with timer= true), example: "BTC"
	Amount         decimal.Decimal `json:"amount"`                // Invoice amount that you want to receive from the user
	ForeignID      string          `json:"foreign_id"`            // Unique foreign ID in your system, example: "164"
	URLSuccess     string          `json:"url_success"`           // URL on which we redirect the user in case of a successful invoice payment, example: "https://merchant.name.com/url_success"
	URLFailed      string          `json:"url_failed"`            // URL on which we redirect the user in case of an unsuccessful invoice payment, example: "https://merchant.name.com/url_failed"
	EmailUser      string          `json:"email_user"`            // In case the payment amount does not match the amount stated above, we will send an email to the stated address with instructions on funds recovery. In case of underpayment, the whole amount will be refunded. In case of overpayment, user will be able to recover the difference by following the instructions
}

// Create invoice for the client for a specified amount.
func (s *InvoicesService) Create(ctx context.Context, opts InvoicesCreateOptions) (*Invoice, error) {
	req, err := s.client.newRequest(http.MethodPost, "invoices/create", opts)
	if err != nil {
		return nil, err
	}

	var invoice *Invoice
	err = s.client.do(ctx, req, &invoice)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}
