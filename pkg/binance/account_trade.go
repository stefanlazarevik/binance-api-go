package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncrequest"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse/acctrade"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (manager *ExchangeManager) GetOrdersList(assets symbol.Assets) ([]order.Info, error) {
	params := url.Values{}
	params.Set(pnames.Symbol, fmt.Sprint(assets.Base, assets.Quote))
	queryStr := bncrequest.Sign(params, manager.apiKey.Secret)

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
	return acctrade.ParseGetOrderList(resp)
}

// SetOrder Set LIMIT or MARKET order on exchange.
// LIMIT - price is mandatory and quantity must be a value in Base equivalent regardless of side.
// MARKET - when buying quantity must in Quote; when selling quantity must be in Base.
func (manager *ExchangeManager) SetOrder(parameters order.Parameters) (order.OrderInfo, error) {
	requestBody := manager.createOrderRequestBody(parameters)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprint(baseUrl, newOrderEndpoint), strings.NewReader(requestBody))
	if err != nil {
		return order.OrderInfo{}, err
	}

	bncrequest.SetHeader(request, manager.apiKey.Key)

	response, err := manager.client.Do(request)
	if err != nil {
		return order.OrderInfo{}, err
	}

	defer bncresponse.CloseBody(response)
	return acctrade.ParseSetOrder(response)
}

func (manager *ExchangeManager) createOrderRequestBody(params order.Parameters) string {
	params = manager.applyFilter(params)
	body := url.Values{}
	body.Set(pnames.Symbol, fmt.Sprint(params.Assets.Base, params.Assets.Quote))
	body.Set(pnames.Side, orderSideAlias[params.Side])
	body.Set(pnames.Type, orderTypeAlias[params.Type])

	if params.Type == order.Limit {
		body.Set(pnames.TimeInForce, "GTC")
		if params.Assets.Quote == "BTC" || params.Assets.Quote == "ETH" {
			body.Set(pnames.Price, strconv.FormatFloat(params.Price, 'f', -1, 64))
		} else {
			body.Set(pnames.Price, fmt.Sprint(params.Price))
		}
		body.Set(pnames.Quantity, fmt.Sprint(params.Quantity))
	} else if params.Type == order.Market {
		if params.Side == order.Buy {
			body.Add(pnames.QuoteOrderQty, fmt.Sprint(params.Quantity))
		} else {
			body.Add(pnames.Quantity, fmt.Sprint(params.Quantity))
		}
	}

	return bncrequest.Sign(body, manager.apiKey.Secret)
}

func (manager *ExchangeManager) GetAssetBalance(asset string) (float64, error) {
	urk := make(url.Values, 0)
	signature := bncrequest.Sign(urk, manager.apiKey.Secret)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, accountInfoEndpoint, "?", signature), nil)
	if err != nil {
		return 0, err
	}

	bncrequest.SetHeader(request, manager.apiKey.Key)

	response, err := manager.client.Do(request)
	if err != nil {
		return 0, err
	}

	defer bncresponse.CloseBody(response)
	return acctrade.ParseBalancesInfo(response, asset)
}

func (manager *ExchangeManager) GetAllCoinsInfo() ([]string, error) {
	urk := make(url.Values, 0)
	signature := bncrequest.Sign(urk, manager.apiKey.Secret)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, getAllCoinsEndpoint, "?", signature), nil)
	if err != nil {
		return nil, err
	}
	bncrequest.SetHeader(request, manager.apiKey.Key)

	response, err := manager.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return acctrade.ParseAllCoinsResponse(response)
}
