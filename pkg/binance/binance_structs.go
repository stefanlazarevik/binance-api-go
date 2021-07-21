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

// binance parametersValue
const (
	symbolParam      = "symbol"
	sideParam        = "side"
	typeParam        = "type"
	quantityParam    = "quantity"
	priceParam       = "price"
	timeInForceParam = "timeInForce"
	totalParams      = "totalParams"
)

// binance api endpoints
const (
	newOrderEndpoint  = "/api/v3/order"
	allOrdersEndpoint = "/api/v3/allOrders"
)

const goodTilCanceled = "GTC"
