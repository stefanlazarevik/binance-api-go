package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"time"
)

const baseUrl = "https://api.binance.com"

//const baseUrl = "https://testnet.binance.vision"

type ExchangeManager struct {
	nextRequestTime time.Time
	symbolsLimits   []symbol.Limits
	apiKey          exchangeapi.ApiKey

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

// errors keys
const (
	RetryAfter = "Retry-After"
	UsedWeight = "X-MBX-USED-WEIGHT"
)

// binance api endpoints
const (
	newOrderEndpoint       = "/api/v3/order"
	openOrdersEndpoint     = "/api/v3/openOrders"
	getPriceEndpoint       = "/api/v3/ticker/price"
	getCandlestickEndpoint = "/api/v3/klines"
	getServerTimeEndpoint  = "/api/v3/time"
	exchangeInfoEndpoint   = "/api/v3/exchangeInfo"
	accountInfoEndpoint    = "/api/v3/account"
	getAllCoinsEndpoint    = "/sapi/v1/capital/config/getall"
)
