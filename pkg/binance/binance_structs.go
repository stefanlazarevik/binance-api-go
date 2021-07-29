package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
)

const baseUrl = "https://api.binance.com"

type ExchangeManager struct {
	apiKey exchangeapi.ApiKey

	client *http.Client
}

func New(key exchangeapi.ApiKey) *ExchangeManager {
	return &ExchangeManager{
		apiKey: key,
		client: &http.Client{},
	}
}

var orderSideAlias = map[order.Side]string{
	order.Buy:  "BUY",
	order.Sell: "SELL",
}

var orderTypeAlias = map[order.Type]string{
	order.Limit:  "LIMIT",
	order.Market: "MARKET",
}

// binance api endpoints
const (
	newOrderEndpoint       = "/api/v3/order"
	openOrdersEndpoint     = "/api/v3/openOrders"
	getPriceEndpoint       = "/api/v3/ticker/price"
	getCandlestickEndpoint = "/api/v3/klines"
)

const goodTilCanceled = "GTC"
