package binance

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse/mktdata"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
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

func (manager *ExchangeManager) GetAllPricesList() ([]symbol.AllPricesList, error) {
	response, err := manager.client.Get(fmt.Sprint(BaseUrl, GetPriceEndpoint))
	if err != nil {
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetAllPricesList(response)
}

func (manager *ExchangeManager) GetAssetOrderBook(asset symbol.Assets, orderBookDepth int) (symbol.OrderBook, error) {
	params := fmt.Sprintf("symbol=%s%s&limit=%d", asset.Base, asset.Quote, orderBookDepth)

	response, err := manager.client.Get(fmt.Sprint(BaseUrl, getAssetOrderBook, "?", params))
	if err != nil {
		return symbol.OrderBook{}, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetAssetOrderBook(response)
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

func (manager *ExchangeManager) GetAllTradingCoins() ([]symbol.Assets, error) {
	response, err := manager.client.Get(fmt.Sprint(BaseUrl, getAllCoinsEndpoint))
	if err != nil {
		return []symbol.Assets{}, err
	}

	defer bncresponse.CloseBody(response)

	return mktdata.GetAllTradingCoins(response)
}

func (manager *ExchangeManager) GetSymbolsBookTicker(assets []symbol.Assets) (map[string]symbol.OrderBook, error) {
	var assetsStr string
	for i := 0; i < len(assets); i++ {
		if i == 0 {
			assetsStr += "["
		}
		if i == len(assets)-1 {
			assetsStr += `"` + assets[i].Base + assets[i].Quote + `"]`
		} else {
			assetsStr += `"` + assets[i].Base + assets[i].Quote + `",`
		}
	}
	var request string
	if assets == nil {
		request = fmt.Sprintf("%s%s", BaseUrl, getSymbolsOrderBook)
	} else {
		request = fmt.Sprintf("%s%s?%s=%s", BaseUrl, getSymbolsOrderBook, pnames.Symbols, assetsStr)
	}

	response, err := manager.client.Get(request)
	if err != nil {
		return map[string]symbol.OrderBook{}, err
	}

	defer bncresponse.CloseBody(response)

	return mktdata.GetSymbolsBookTicker(response)
}
