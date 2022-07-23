package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
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

func GetPricesMap(response *http.Response, arbitrageAssets []string) ([]symbol.AssetInfo, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	priceI, isOk := bodyI.([]map[string]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> error when casting bodyI to priceI")
	}

	assetPricesArr := make([]symbol.AssetInfo, 0)

	for _, assets := range priceI {
		symbolAsset, isOk := assets[pnames.Symbol].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting asset symbol to string")
		}

		priceStr, isOk := assets[pnames.Price].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting price to string")
		}

		if strings.Contains(symbolAsset, "LUNA") || strings.Contains(symbolAsset, "WRX") || strings.Contains(symbolAsset, "BTT") {
			continue
		}
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, errors.New("[mkdata] -> error error when parsing priceStr to float64")
		}

		if strings.Contains(symbolAsset, arbitrageAssets[0]) || strings.Contains(symbolAsset, arbitrageAssets[1]) {
			assetInfo := symbol.AssetInfo{
				Symbol: symbolAsset,
				Price:  price,
			}

			assetPricesArr = append(assetPricesArr, assetInfo)
		}
	}

	return assetPricesArr, nil
}
