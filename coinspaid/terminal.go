package coinspaid

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

const terminalURL = `https://terminal.coinspaid.com/`

type TerminalOptions struct {
	ClientID  int64   `json:"client_id"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount,omitempty"`
	ConvertTo *string `json:"convert_to,omitempty"`
	IsIFrame  bool    `json:"is_iframe,omitempty"`
	ForeignID string  `json:"foreign_id"`
	URLBack   string  `json:"url_back,omitempty"`
}

func (c *Client) GetTerminalLink(opts TerminalOptions) string {
	data, _ := json.Marshal(opts)
	dataStr := base64.URLEncoding.EncodeToString(data)

	u, _ := url.Parse(terminalURL)
	q := u.Query()
	q.Add("data", dataStr)
	q.Add("signature", genSignature(c.secret, data))
	u.RawQuery = q.Encode()

	return u.String()
}
