package coinspaid

import "testing"

func Test_getSignature(t *testing.T) {
	const want = "03c25fcf7cd35e7d995e402cd5d51edd72d48e1471e865907967809a0c189ba55b90815f20e2bb10f82c7a9e9d86554" +
		"6fda58989c2ae9e8e2ff7bc29195fa1ec"

	sig := genSignature("AbCdEfG123456", []byte(`{"currency":"BTC","foreign_id":"123456"}`))
	if sig != want {
		t.Errorf("invalid signature: %s, want %s", sig, want)
	}
}
