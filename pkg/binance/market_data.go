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
	response, err := manager.client.Get(fmt.Sprint(BaseUrl, GetPriceEndpoint, "?", params))
	if err != nil {
		return 0, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetCurrentPrice(response)
}

func (manager *ExchangeManager) GetPricesMap(baseUrl, endpoint string) ([]symbol.AssetInfo, error) {
	response, err := manager.client.Get(fmt.Sprint(baseUrl, endpoint))
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetPricesMap(response)
}

func (manager *ExchangeManager) GetCandlestick(symbol symbol.Assets, interval string, limit int) ([]exchangeapi.Candlestick, error) {
	params := fmt.Sprintf("symbol=%s%s&interval=%s&limit=%d", symbol.Base, symbol.Quote, interval, limit)
	response, err := manager.client.Get(fmt.Sprint(BaseUrl, getCandlestickEndpoint, "?", params))
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetCandlestick(response)
}

func (manager *ExchangeManager) GetServerTime() (time.Time, error) {
	response, err := manager.client.Get(fmt.Sprint(BaseUrl, getServerTimeEndpoint))
	if err != nil {
		return time.Time{}, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetServerTime(response)
}

func (manager *ExchangeManager) GetSymbolsLimits() ([]symbol.Limits, error) {
	response, err := manager.client.Get(fmt.Sprintf("%s%s", BaseUrl, exchangeInfoEndpoint))
	if err != nil {
		return []symbol.Limits{}, err
	}

	defer bncresponse.CloseBody(response)
	limits, err := mktdata.GetSymbolLimits(response)
	if err != nil {
		manager.checkReqError(err)
		return nil, err
	}

	return limits, nil
}
