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

	accounts, err := cli.Accounts.List(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, account := range accounts {
		fmt.Printf("%s: %s\n", account.Currency, account.Balance)
	}
}
