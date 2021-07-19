package binance_api

import (
	"net/http"
	"posipaka-trade-cmn/exchangeapi"
)

const baseUrl = "https://api.binance.com"

type BinanceExchangeManager struct {
	apiKey exchangeapi.ApiKey

	client *http.Client
}

func New(key exchangeapi.ApiKey) *BinanceExchangeManager {
	return &BinanceExchangeManager{
		apiKey: key,
		client: &http.Client{},
	}
}

var orderSideAlias = map[exchangeapi.OrderSide]string{
	exchangeapi.Buy: "BUY",
	exchangeapi.Sell: "SELL",
}

var orderTypeAlias = map[exchangeapi.OrderType]string{
	exchangeapi.Limit: "LIMIT",
	exchangeapi.Market: "MARKET",
}

// binance parametersValue
const (
	symbolParam = "symbol"
	sideParam = "side"
	typeParam = "type"
	quantityParam = "quantity"
	priceParam = "price"
	timeInForceParam = "timeInForce"
	totalParams = "totalParams"
)

// binance api endpoints
const (
	newOrderEndpoint = "/api/v3/order"
)

const goodTilCanceled = "GTC"