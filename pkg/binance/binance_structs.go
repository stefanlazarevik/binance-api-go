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
	newOrderEndpoint       = "/api/v3/order"
	allOrdersEndpoint      = "/api/v3/allOrders"
	getPriceEndpoint       = "/api/v3/ticker/price"
	getCandlestickEndpoint = "/api/v3/klines"
)

const goodTilCanceled = "GTC"

type Candlesticks struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrade            int64  `json:"numberOfTrade"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	Ignore                   string `json:"ignore"`
}
