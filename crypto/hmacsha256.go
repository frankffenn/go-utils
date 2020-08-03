package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func HmacSha256(s, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(s))

	return fmt.Sprintf("%x", h.Sum(nil))
}
