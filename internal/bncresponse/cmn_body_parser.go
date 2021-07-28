package bncresponse

import (
	"encoding/json"
	"errors"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"io/ioutil"
	"net/http"
)

func getResponseBody(response *http.Response) (interface{}, error) {
	if response.StatusCode/100 != 2 && response.Body == nil {
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

	if respondBody[0] == '[' && respondBody[len(respondBody)-1] == ']' {
		var body []map[string]interface{}
		err = json.Unmarshal(respondBody, &body)
		if err != nil {
			return nil, err
		}

		return body, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(respondBody, &body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode/100 != 2 {
		return nil, parseBinanceError(body)
	}

	return body, nil
}

func parseBinanceError(body map[string]interface{}) error {
	code, isOkay := body[codeKey].(float64)
	if !isOkay {
		return errors.New("[bncresponse] -> error code key not found")
	}

	message, isOkay := body[msgKey].(string)
	if !isOkay {
		return errors.New("[bncresponse] -> failed to parse binance error message")
	}

	return &exchangeapi.ExchangeError{
		Type:    exchangeapi.BinanceErr,
		Code:    int(code),
		Message: message,
	}
}
