// +build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/serjvanilla/coinspaid-go/coinspaid"

	"github.com/stretchr/testify/assert"
)

func TestAddressesService_Take(t *testing.T) {
	c := getSandboxClient(t)

	opts := coinspaid.AddressTakeOptions{
		ForeignID: fmt.Sprintf("test-%d", time.Now().UnixNano()),
		Currency:  "BTC",
		ConvertTo: coinspaid.String("USD"),
	}

	address, err := c.Addresses.Take(context.Background(), opts)
	assert.NoError(t, err, "should complete without error")
	assert.Equal(t, address.Currency, "BTC", "currency should be equal BTC")
	assert.Equal(t, address.ConvertTo, coinspaid.String("USD"), "convert to should be equal USD")
	assert.Equal(t, address.ForeignID, opts.ForeignID, "foreign id should match opts")
	assert.Len(t, address.Address, 35, "address should be 35 characters long")
}
