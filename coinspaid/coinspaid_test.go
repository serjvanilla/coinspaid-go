package coinspaid

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const baseURLPath = `/api/v2`

func setupTest() (client *Client, mux *http.ServeMux, server *httptest.Server, cleanup func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	server = httptest.NewServer(apiHandler)

	client = NewClient("key", "secret")
	u, _ := url.Parse(server.URL + baseURLPath + "/")
	client.baseURL = u

	return client, mux, server, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("request method: %v, want %v", got, want)
	}
}

func testHeaders(t *testing.T, r *http.Request, key, secret string) {
	t.Helper()

	if h := r.Header.Get(headerProcessingKey); h != key {
		t.Errorf("%s header: %v, want %v", headerProcessingKey, h, key)
	}

	data, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))

	want := genSignature(secret, data)
	if h := r.Header.Get(headerProcessingSignature); h != want {
		if h := r.Header.Get(headerProcessingSignature); h != want {
			t.Errorf("%s header: %v, want %v", headerProcessingSignature, h, want)
		}
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	t.Helper()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}
