package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse/mktdata"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"time"
)

func (manager *ExchangeManager) GetCurrentPrice(symbol symbol.Assets) (float64, error) {
	params := fmt.Sprintf("symbol=%s%s", symbol.Base, symbol.Quote)
	response, err := manager.client.Get(fmt.Sprint(baseUrl, getPriceEndpoint, "?", params))
	if err != nil {
		return 0, err
	}
	defer bncresponse.CloseBody(response)
	return mktdata.GetCurrentPrice(response)
}

func (manager *ExchangeManager) GetCandlestick(symbol symbol.Assets, interval string, limit int) ([]exchangeapi.Candlestick, error) {
	params := fmt.Sprintf("symbol=%s%s&interval=%s&limit=%d", symbol.Base, symbol.Quote, interval, limit)

	response, err := manager.client.Get(fmt.Sprint(baseUrl, getCandlestickEndpoint, "?", params))
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetCandlestick(response)
}

func (manager *ExchangeManager) GetServerTime() (time.Time, error) {
	response, err := manager.client.Get(fmt.Sprint(baseUrl, getServerTimeEndpoint))
	if err != nil {
		return time.Time{}, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetServerTime(response)
}

func (manager *ExchangeManager) GetSymbolsLimits() ([]symbol.Limits, error) {
	response, err := manager.client.Get(fmt.Sprintf("%s%s", baseUrl, exchangeInfoEndpoint))
	if err != nil {
		return []symbol.Limits{}, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetSymbolLimits(response)
}
