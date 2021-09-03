package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
)

func GetSymbolsList(response *http.Response) ([]symbol.Assets, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	body, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return nil, errors.New("[mktdata] -> Response with symbols list body parsing failed")
	}

	symbolsListI, isOkay := body[pnames.Symbols].([]interface{})
	if !isOkay {
		return nil, errors.New("[mktdata] -> Symbols list not found in response body")
	}

	symbolsList := parseSymbolsAssets(symbolsListI)
	if len(symbolsList) == 0 {
		return nil, errors.New("[mktdata] -> No symbols got after parsing")
	}

	return symbolsList, nil
}

func parseSymbolsAssets(symbolsListI []interface{}) []symbol.Assets {
	symbolsList := make([]symbol.Assets, 0)
	for _, symbolInfoI := range symbolsListI {
		info, isOkay := symbolInfoI.(map[string]interface{})
		if !isOkay {
			continue
		}

		base, isOkay := info[pnames.BaseAsset].(string)
		if !isOkay {
			continue
		}

		quote, isOkay := info[pnames.QuoteAsset].(string)
		if !isOkay {
			continue
		}

		symbolsList = append(symbolsList, symbol.Assets{
			Base:  base,
			Quote: quote,
		})
	}

	return symbolsList
}
