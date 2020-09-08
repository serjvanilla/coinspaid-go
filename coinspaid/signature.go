package coinspaid

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func genSignature(secret string, data []byte) string {
	h := hmac.New(sha512.New, []byte(secret))
	_, _ = h.Write(data)

	return hex.EncodeToString(h.Sum(nil))
}
