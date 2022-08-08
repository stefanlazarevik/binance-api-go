package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
)

func GetAllTradingCoins(response *http.Response) ([]symbol.Assets, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}
	coinsI, isOkay := bodyI.(map[string]interface{})
	if isOkay != true {
		return nil, errors.New("[bncresponse] -> Error when casting body to map[string]interface{} in all coins receiving")
	}

	symbols, isOkay := coinsI["symbols"].([]interface{})
	if isOkay != true {
		return nil, errors.New("[bncresponse] -> Error when casting symbols to []interface{}")
	}

	assets := make([]symbol.Assets, 0)
	for _, symbolI := range symbols {

		symbolInfo, isOkay := symbolI.(map[string]interface{})
		if isOkay != true {
			return nil, errors.New("[bncresponse] -> Error when casting symbols info to map[string]interface{}")
		}

		if symbolInfo[pnames.Status].(string) == pnames.Trading {

			assets = append(assets, symbol.Assets{
				Base:  symbolInfo["baseAsset"].(string),
				Quote: symbolInfo["quoteAsset"].(string),
			})
		}
	}

	return assets, nil
}
