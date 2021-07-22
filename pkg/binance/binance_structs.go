package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
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

var orderSideAlias = map[exchangeapi.OrderSide]string{
	exchangeapi.Buy:  "BUY",
	exchangeapi.Sell: "SELL",
}

var orderTypeAlias = map[exchangeapi.OrderType]string{
	exchangeapi.Limit:  "LIMIT",
	exchangeapi.Market: "MARKET",
}

// binance api endpoints
const (
	newOrderEndpoint   = "/api/v3/order"
	openOrdersEndpoint = "/api/v3/openOrders"
)

const goodTilCanceled = "GTC"
