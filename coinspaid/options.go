package coinspaid

import (
	"net/http"
	"net/url"
)

type option func(*Client)

// WithSandbox is NewClient option for interacting with CoinsPaid's test environment.
func WithSandbox() option {
	return func(c *Client) {
		c.baseURL, _ = url.Parse(sandboxBaseURL)
	}
}

// WithHTTPClient is NewClient option to pass your custom http.Client.
//
// Strongly recommended, otherwise http.DefaultClient will be used.
func WithHTTPClient(client *http.Client) option {
	return func(c *Client) {
		c.client = client
	}
}
