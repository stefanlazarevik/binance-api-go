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

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(BaseUrl, openOrdersEndpoint, "?", queryStr), nil)
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
func (manager *ExchangeManager) SetOrder(parameters order.Parameters) (order.Info, error) {
	requestBody := manager.createOrderRequestBody(parameters)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprint(BaseUrl, newOrderEndpoint), strings.NewReader(requestBody))
	if err != nil {
		return order.Info{}, err
	}

	bncrequest.SetHeader(request, manager.apiKey.Key)

	response, err := manager.client.Do(request)
	if err != nil {
		return order.Info{}, err
	}

	defer bncresponse.CloseBody(response)
	orderInfo, err := acctrade.ParseOrderInfoResponse(response)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Assets = parameters.Assets
	if orderInfo.Type == order.Market {
		orderInfo.QuoteQuantity = parameters.Quantity
	} else if orderInfo.Type == order.Limit {
		orderInfo.Price = parameters.Price
	}

	return orderInfo, nil
}

func (manager *ExchangeManager) createOrderRequestBody(params order.Parameters) string {
	params = manager.applyFilter(params)
	body := url.Values{}
	body.Set(pnames.Symbol, fmt.Sprint(params.Assets.Base, params.Assets.Quote))
	body.Set(pnames.Side, orderSideAlias[params.Side])
	body.Set(pnames.Type, orderTypeAlias[params.Type])

	if params.Type == order.Limit {
		body.Set(pnames.TimeInForce, "GTC")
		body.Set(pnames.Price, strconv.FormatFloat(params.Price, 'f', -1, 64))
		body.Set(pnames.Quantity, strconv.FormatFloat(params.Quantity, 'f', -1, 64))
	} else if params.Type == order.Market {
		if params.Side == order.Buy {
			body.Add(pnames.QuoteOrderQty, strconv.FormatFloat(params.Quantity, 'f', -1, 64))
		} else {
			body.Add(pnames.Quantity, fmt.Sprint(params.Quantity))
		}
	}

	return bncrequest.Sign(body, manager.apiKey.Secret)
}

func (manager *ExchangeManager) GetAssetBalance(asset string) (float64, error) {
	urk := make(url.Values, 0)
	signature := bncrequest.Sign(urk, manager.apiKey.Secret)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprint(BaseUrl, accountInfoEndpoint, "?", signature), nil)
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
