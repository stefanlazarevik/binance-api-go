package binance

import (
	"errors"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse/mktdata"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"time"
)

func (manager *ExchangeManager) GetCurrentPrice(symbol symbol.Assets) (float64, error) {
	params := fmt.Sprintf("symbol=%s%s", symbol.Base, symbol.Quote)

	if time.Now().Before(manager.nextRequestTime) {
		return 0, errors.New("[binance] -> Getting price is impossible due to block. Waiting until " +
			fmt.Sprint(manager.nextRequestTime.Unix()))
	}
	response, err := manager.client.Get(fmt.Sprint(baseUrl, getPriceEndpoint, "?", params))
	if err != nil {
		manager.checkReqError(err)
		return 0, err
	}
	defer bncresponse.CloseBody(response)
	return mktdata.GetCurrentPrice(response)
}

func (manager *ExchangeManager) GetCandlestick(symbol symbol.Assets, interval string, limit int) ([]exchangeapi.Candlestick, error) {
	params := fmt.Sprintf("symbol=%s%s&interval=%s&limit=%d", symbol.Base, symbol.Quote, interval, limit)

	if time.Now().Before(manager.nextRequestTime) {
		return nil, errors.New("[binance] -> Getting candlestick is impossible due to block. Waiting until " +
			fmt.Sprint(manager.nextRequestTime.Unix()))
	}
	response, err := manager.client.Get(fmt.Sprint(baseUrl, getCandlestickEndpoint, "?", params))
	if err != nil {
		manager.checkReqError(err)
		return nil, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetCandlestick(response)
}

func (manager *ExchangeManager) GetServerTime() (time.Time, error) {
	if time.Now().Before(manager.nextRequestTime) {
		return time.Time{}, errors.New("[binance] -> Getting server time is impossible due to block. Waiting until " +
			fmt.Sprint(manager.nextRequestTime.Unix()))
	}
	response, err := manager.client.Get(fmt.Sprint(baseUrl, getServerTimeEndpoint))
	if err != nil {
		manager.checkReqError(err)
		return time.Time{}, err
	}

	defer bncresponse.CloseBody(response)
	return mktdata.GetServerTime(response)
}

func (manager *ExchangeManager) GetSymbolsLimits() ([]symbol.Limits, error) {
	if time.Now().Before(manager.nextRequestTime) {
		return nil, errors.New("[binance] -> Getting symbol limits is impossible due to block. Waiting until " +
			fmt.Sprint(manager.nextRequestTime.Unix()))
	}
	response, err := manager.client.Get(fmt.Sprintf("%s%s", baseUrl, exchangeInfoEndpoint))
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
