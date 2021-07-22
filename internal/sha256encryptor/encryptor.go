package sha256encryptor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptMessage(message, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}
