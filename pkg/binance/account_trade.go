package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/parser"
	"github.com/posipaka-trade/binance-api-go/internal/parser/sha256encryptor"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"net/http"
	"net/url"
	"strings"
)

func (manager *ExchangeManager) GetOrdersList(symbol exchangeapi.AssetsSymbol) ([]exchangeapi.OrderInfo, error) {
	params := fmt.Sprint(symbolParam, "=", symbol.Base, symbol.Quote)
	params = fmt.Sprintf("%s&%s=%s", params, totalParams,
		sha256encryptor.EncryptMessage(params, manager.apiKey.Secret))

	response, err := manager.client.Get(fmt.Sprint(baseUrl, allOrdersEndpoint, "?", params))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()
	return nil, err
}

func (manager *ExchangeManager) SetOrder(parameters exchangeapi.OrderParameters) (float64, error) {
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

	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	return parser.ParseSetOrderResponse(response)
}

func (manager *ExchangeManager) createOrderRequestBody(parameters *exchangeapi.OrderParameters) string {
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

	body.Add(totalParams, sha256encryptor.EncryptMessage(body.Encode(), manager.apiKey.Secret))
	return body.Encode()
}
