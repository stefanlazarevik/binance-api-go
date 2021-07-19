package binance_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/parser"
	"net/http"
	"net/url"
	"posipaka-trade-cmn/exchangeapi"
	"strings"
)

func (manager *BinanceExchangeManager) SetOrder(parameters exchangeapi.OrderParameters) (float64, error) {
	requestBody := manager.createOrderRequestBody(&parameters)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprint(baseUrl, newOrderEndpoint), strings.NewReader(requestBody))
	if err != nil {
		return 0, err
	}

	request.Header.Set("X-MBX-APIKEY", manager.apiKey.Key)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := manager.client.Do(request)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	return parser.ParseSetOrderResponse(response)
}

func (manager *BinanceExchangeManager) createOrderRequestBody(parameters *exchangeapi.OrderParameters) string {
	body := url.Values{}
	body.Add(symbolParam, fmt.Sprint(parameters.Symbol.Base, parameters.Symbol.Quote))
	body.Add(sideParam, orderSideAlias[parameters.Side])
	body.Add(typeParam, orderTypeAlias[parameters.Type])

	if parameters.Type == exchangeapi.Limit {
		body.Add(timeInForceParam, "GTC")
		body.Add(priceParam, fmt.Sprint(parameters.Price))
		body.Add(quantityParam, fmt.Sprint(parameters.Quantity))
	} else if parameters.Type == exchangeapi.Market {
		body.Add(quantityParam, fmt.Sprint(parameters.Quantity))
	}

	hash := hmac.New(sha256.New, []byte(manager.apiKey.Secret))
	hash.Write([]byte(body.Encode()))

	body.Add(totalParams, hex.EncodeToString(hash.Sum(nil)))
	return body.Encode()
}
