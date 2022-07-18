package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"strconv"
	"strings"
)

func GetCurrentPrice(response *http.Response) (float64, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return 0, err
	}

	priceI, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return 0, errors.New("[mktdata] -> error when casting bodyI to priceI")
	}

	priceStr := priceI[pnames.Price]
	price, err := strconv.ParseFloat(priceStr.(string), 64)
	if err != nil {
		return 0, errors.New("[mktdata] -> error when parsing priceStr to float64")
	}
	return price, nil
}

func GetPricesMap(response *http.Response) (map[string]float64, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	priceI, isOk := bodyI.([]map[string]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> error when casting bodyI to priceI")
	}
	pricesMap := make(map[string]float64)

	for _, t := range priceI {
		symbol, isOk := t[pnames.Symbol].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting bodyI to priceI")
		}

		priceStr, isOk := t[pnames.Price].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting bodyI to priceI")
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, errors.New("[mkdata] -> error error when parsing priceStr to float64")
		}

		if strings.Contains(symbol, "1000") {
			symbol = strings.ReplaceAll(symbol, "1000", "")
			price = price / 1000
		}
		if strings.Contains(symbol, "2") {
			symbol = strings.ReplaceAll(symbol, "2", "")
		}
		if symbol == "SCUSDT" || symbol == "ICPUSDT" || symbol == "TLMUSDT" {
			continue
		}
		pricesMap[symbol] = price
	}

	return pricesMap, nil
}
