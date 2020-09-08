package coinspaid_test

import (
	"context"
	"log"
	"net/http"

	"github.com/serjvanilla/coinspaid-go/coinspaid"
)

func ExampleClient_CallbackHandler() {
	coinspaidClient := coinspaid.NewClient("your_key", "your_secret", coinspaid.WithSandbox())

	handler := func(ctx context.Context, data *coinspaid.CallbackData) (ok bool) {
		log.Printf("callback received: id %d, type %s\n", data.ID, data.Type)

		return true
	}

	mux := http.NewServeMux()
	mux.Handle("/callback", coinspaidClient.CallbackHandler(handler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
