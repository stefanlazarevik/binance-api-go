package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncrequest"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"net/url"
	"strings"
)

func (manager *ExchangeManager) GetOrdersList(assets symbol.Assets) ([]order.Info, error) {
	params := url.Values{}
	params.Set(pnames.Symbol, fmt.Sprint(assets.Base, assets.Quote))
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

func (manager *ExchangeManager) SetOrder(parameters order.Parameters) (float64, error) {
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

func (manager *ExchangeManager) createOrderRequestBody(params *order.Parameters) string {
	body := url.Values{}
	body.Set(pnames.Symbol, fmt.Sprint(params.Assets.Base, params.Assets.Quote))
	body.Set(pnames.Side, orderSideAlias[params.Side])
	body.Set(pnames.Type, orderTypeAlias[params.Type])

	if params.Type == order.Limit {
		body.Set(pnames.TimeInForce, "GTC")
		body.Set(pnames.Price, fmt.Sprint(params.Price))
		body.Set(pnames.Quantity, fmt.Sprint(params.Quantity))
	} else if params.Type == order.Market {
		body.Add(pnames.QuoteOrderQty, fmt.Sprint(params.Quantity))
	}

	return bncrequest.Sing(body, manager.apiKey.Secret)
}
func (manager *ExchangeManager) GetCurrentPrice(symbol symbol.Assets) (float64, error) {
	params := fmt.Sprintf("symbol=%s%s", symbol.Base, symbol.Quote)

	response, err := manager.client.Get(fmt.Sprint(baseUrl, getPriceEndpoint, "?", params))
	if err != nil {
		return 0, err
	}

	return bncresponse.GetCurrentPrice(response)
}
func (manager *ExchangeManager) GetCandlestick(symbol symbol.Assets, interval string, limit int) (exchangeapi.Candlesticks, error) {
	params := fmt.Sprintf("symbol=%s%s&interval=%s&limit=%d", symbol.Base, symbol.Quote, interval, limit)

	response, err := manager.client.Get(fmt.Sprint(baseUrl, getCandlestickEndpoint, "?", params))
	if err != nil {
		return exchangeapi.Candlesticks{}, err
	}

	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return nil, err
	//}
	//var candleData []Candlesticks
	//
	//err = json.Unmarshal(body, &candleData)
	//if err != nil {
	//	return nil, err
	//}

	return bncresponse.GetCandlestick(response)
}
