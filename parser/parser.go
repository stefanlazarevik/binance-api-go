package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"posipaka-trade-cmn/exchangeapi"
)

// responses json keys
const (
	codeKey = "code" // error code
	msgKey = "msg"   // error message
)

func ParseSetOrderResponse(response *http.Response) (float64, error) {
	_, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func getResponseBody(response *http.Response) (map[string]interface{}, error) {
	if response.StatusCode / 100 != 2 && response.Body == nil {
		return nil, &exchangeapi.ExchangeError{
			Type:    exchangeapi.HttpErr,
			Code:    response.StatusCode,
			Message: response.Status,
		}
	}

	respondBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(respondBody, &body)
	if err != nil  {
		return nil, err
	}

	if response.StatusCode / 100 != 2 {
		return nil, parseBinanceError(body)
	}

	return body, nil
}

func parseBinanceError(body map[string]interface{}) error {
	code, isOkay := body[codeKey].(int)
	if !isOkay {
		return errors.New("[binance-api-go.parser] -> failed to parse binance error code")
	}

	message, isOkay := body[msgKey].(string)
	if !isOkay {
		return errors.New("[binance-api-go.parser] -> failed to parse binance error message")
	}

	return &exchangeapi.ExchangeError{
		Type:    exchangeapi.BinanceErr,
		Code:	 code,
		Message: message,
	}
}