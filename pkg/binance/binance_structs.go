package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"sync"
	"time"
)

const BaseUrl = "https://api.binance.com"
const BaseMarginUrl = "https://fapi.binance.com"

//const BaseUrl = "https://testnet.binance.vision"

type ExchangeManager struct {
	symbolsLimits []symbol.Limits
	apiKey        exchangeapi.ApiKey

	client *http.Client

	isWorking bool
	wg        sync.WaitGroup
}

func New(key exchangeapi.ApiKey) *ExchangeManager {
	mgr := &ExchangeManager{
		apiKey: key,
		client: &http.Client{
			Transport: http.DefaultTransport,
		},
		isWorking: true,
	}

	mgr.wg.Add(1)
	go func() {
		defer mgr.wg.Done()
		lastConnectionTime := time.Time{}
		for mgr.isWorking {
			if time.Now().Sub(lastConnectionTime) >= 75*time.Second {
				_, _ = mgr.client.Get(BaseUrl)
				lastConnectionTime = time.Now()
			}
			time.Sleep(time.Second)
		}
	}()

	return mgr
}

// Finish completes inner goroutines
func (manager *ExchangeManager) Finish() {
	manager.isWorking = false
	manager.wg.Wait()
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
	MarginGetPriceEndpoint = "/fapi/v1/ticker/price"
	newOrderEndpoint       = "/api/v3/order"
	openOrdersEndpoint     = "/api/v3/openOrders"
	GetPriceEndpoint       = "/api/v3/ticker/price"
	getCandlestickEndpoint = "/api/v3/klines"
	getServerTimeEndpoint  = "/api/v3/time"
	exchangeInfoEndpoint   = "/api/v3/exchangeInfo"
	accountInfoEndpoint    = "/api/v3/account"
	getAllCoinsEndpoint    = "/api/v3/exchangeInfo"
	getAssetOrderBook      = "/api/v3/depth"
	getSymbolsOrderBook    = "/api/v3/ticker/bookTicker"
	getOrderInformation    = "/api/v3/order"
	cancelReplaceOrder     = "/api/v3/order/cancelReplace"
)
