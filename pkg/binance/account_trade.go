package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncrequest"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"net/http"
	"net/url"
	"strings"
)

func (manager *ExchangeManager) GetOrdersList(symbol exchangeapi.AssetsSymbol) ([]exchangeapi.OrderInfo, error) {
	params := url.Values{}
	params.Set(pnames.Symbol, fmt.Sprint(symbol.Base, symbol.Quote))
	queryStr := bncrequest.Sing(params, manager.apiKey.Secret)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, openOrdersEndpoint, "?", queryStr), nil)
	if err != nil {
		return nil, err
	}

	bncrequest.SetHeader(req, manager.apiKey.Key)

	resp, err := manager.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(resp)
	return bncresponse.ParseGetOrderList(resp)
}

func (manager *ExchangeManager) SetOrder(parameters exchangeapi.OrderParameters) (float64, error) {
	requestBody := manager.createOrderRequestBody(&parameters)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprint(baseUrl, newOrderEndpoint), strings.NewReader(requestBody))
	if err != nil {
		return 0, err
	}

	bncrequest.SetHeader(request, manager.apiKey.Key)

	response, err := manager.client.Do(request)
	if err != nil {
		return 0, err
	}

	defer bncresponse.CloseBody(response)
	return bncresponse.ParseSetOrder(response)
}

func (manager *ExchangeManager) createOrderRequestBody(params *exchangeapi.OrderParameters) string {
	body := url.Values{}
	body.Set(pnames.Symbol, fmt.Sprint(params.Symbol.Base, params.Symbol.Quote))
	body.Set(pnames.Side, orderSideAlias[params.Side])
	body.Set(pnames.Type, orderTypeAlias[params.Type])

	if params.Type == exchangeapi.Limit {
		body.Set(pnames.TimeInForce, "GTC")
		body.Set(pnames.Price, fmt.Sprint(params.Price))
		body.Set(pnames.Quantity, fmt.Sprint(params.Quantity))
	} else if params.Type == exchangeapi.Market {
		body.Add(pnames.QuoteOrderQty, fmt.Sprint(params.Quantity))
	}

	return bncrequest.Sing(body, manager.apiKey.Secret)
}
