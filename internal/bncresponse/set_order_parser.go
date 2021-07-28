package bncresponse

import (
	"errors"
	"net/http"
)

func ParseSetOrder(response *http.Response) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}

	_, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return 0, errors.New("[bncresponse] -> set order response is not key/value pair array")
	}
	return 0, nil
}
