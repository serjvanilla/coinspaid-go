package coinspaid

import (
	"bytes"
	"context"
	"net/http"
)

// Ping tests if API is up and running and your authorization is working.
func (c *Client) Ping(ctx context.Context) (string, error) {
	req, err := c.newRequest(http.MethodGet, "ping", struct{}{})
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = c.do(ctx, req, buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
