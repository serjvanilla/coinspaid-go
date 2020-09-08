package coinspaid

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/shopspring/decimal"
)

const (
	sandboxBaseURL    = "https://app.sandbox.cryptoprocessing.com/api/v2/"
	productionBaseURL = "https://app.coinspaid.com/api/v2/"

	headerProcessingKey       = "X-Processing-Key"
	headerProcessingSignature = "X-Processing-Signature"
)

var ErrNilContext = errors.New("context must be non-nil")

type Client struct {
	client  *http.Client
	baseURL *url.URL

	key, secret string

	common service

	Accounts   *AccountsService
	Addresses  *AddressesService
	Currencies *CurrenciesService
	Exchange   *ExchangeService
	Futures    *FuturesService
	Invoices   *InvoicesService
	Withdrawal *WithdrawalService
}

type service struct {
	client *Client
}

// NewClient returns a new CoinsPaid API client.
func NewClient(key, secret string, opts ...option) *Client {
	c := &Client{
		client: http.DefaultClient,
		key:    key,
		secret: secret,
	}
	c.baseURL, _ = url.Parse(productionBaseURL)

	for _, opt := range opts {
		opt(c)
	}

	c.common.client = c

	c.Accounts = (*AccountsService)(&c.common)
	c.Addresses = (*AddressesService)(&c.common)
	c.Currencies = (*CurrenciesService)(&c.common)
	c.Exchange = (*ExchangeService)(&c.common)
	c.Futures = (*FuturesService)(&c.common)
	c.Invoices = (*InvoicesService)(&c.common)
	c.Withdrawal = (*WithdrawalService)(&c.common)

	return c
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Decimal is a helper routine that allocates a new decimal.Decimal value
// to store v and returns a pointer to it.
func Decimal(v decimal.Decimal) *decimal.Decimal { return &v }

// newRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(headerProcessingKey, c.key)
	req.Header.Set(headerProcessingSignature, genSignature(c.secret, data))

	return req, nil
}

// do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	if ctx == nil {
		return ErrNilContext
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
	case http.StatusForbidden:
		var authErr *AuthorizationError
		err = decodeJSONResponse(resp.Body, &authErr)
		if err != nil {
			return err
		}

		return authErr
	default:
		var apiErr *APIError
		err = decodeJSONResponse(resp.Body, &apiErr)
		if err != nil {
			return err
		}

		apiErr.StatusCode = resp.StatusCode

		return apiErr
	}

	type r struct {
		Data interface{} `json:"data"`
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = decodeJSONResponse(resp.Body, &r{Data: v})
		}
	}

	return err
}

func decodeJSONResponse(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if errors.Is(err, io.EOF) {
		err = nil
	}

	return err
}
