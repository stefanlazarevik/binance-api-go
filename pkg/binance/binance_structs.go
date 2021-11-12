package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"sync"
	"time"
)

const baseUrl = "https://api.binance.com"

//const baseUrl = "https://testnet.binance.vision"

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
				_, _ = mgr.client.Get(baseUrl)
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
	newOrderEndpoint       = "/api/v3/order"
	openOrdersEndpoint     = "/api/v3/openOrders"
	getPriceEndpoint       = "/api/v3/ticker/price"
	getCandlestickEndpoint = "/api/v3/klines"
	getServerTimeEndpoint  = "/api/v3/time"
	exchangeInfoEndpoint   = "/api/v3/exchangeInfo"
	accountInfoEndpoint    = "/api/v3/account"
	getAllCoinsEndpoint    = "/sapi/v1/capital/config/getall"
)
