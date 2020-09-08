# coinspaid-go #
Go client library for accessing [CoinsPaid API v2](https://docs.coinspaid.com/docs/api-documentation/api-reference).

[![PkgGoDev](https://pkg.go.dev/badge/github.com/serjvanilla/coinspaid-go)](https://pkg.go.dev/github.com/serjvanilla/coinspaid-go)

## Usage ##
```go
package main

import (
    "context"
    "fmt"
    "net/http"

    "github.com/serjvanilla/coinspaid-go/coinspaid"
)

func main() {
	cli := coinspaid.NewClient(
		"your_key", "your_secret",
		coinspaid.WithSandbox(),                  // using sandbox for testing purpose
		coinspaid.WithHTTPClient(&http.Client{}), // using custom http client, otherwise http.DefaultClient will be used (not recommended)
	)

	address, err := cli.Addresses.Take(
		context.Background(),
		coinspaid.AddressTakeOptions{
			ForeignID: "foreign-id-12345",
			Currency:  "BTC",
			ConvertTo: coinspaid.String("USD"),
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New address[%d]: %s %s\n", address.ID, address.Currency, address.Address)
}
```

