package bncrequest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"net/url"
	"time"
)

func SetHeader(req *http.Request, key string) {
	req.Header.Set("X-MBX-APIKEY", key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

func Sing(data url.Values, key string) string {
	timestampMs := time.Now().UnixNano() / int64(time.Millisecond)
	data.Set(pnames.Timestamp, fmt.Sprint(timestampMs))

	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(data.Encode()))

	return fmt.Sprintf("%s&%s=%s", data.Encode(), pnames.Signature, hex.EncodeToString(hash.Sum(nil)))
}
