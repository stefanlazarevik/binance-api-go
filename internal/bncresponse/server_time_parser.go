package bncresponse

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"time"
)

func GetServerTime(response *http.Response) (time.Time, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return time.Time{}, err
	}

	timeI, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return time.Time{}, errors.New("[bncresponse] -> error when casting bodyI to timeI")
	}

	serverTime, isOkay := timeI[pnames.ServerTime].(float64)
	if !isOkay {
		return time.Time{}, errors.New("[bncresponse] -> error when parsing server time to float64")
	}
	return time.Unix(0, int64(serverTime)*int64(time.Millisecond)), nil
}
