package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"strconv"
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

func GetAllPricesList(response *http.Response, assetsList []symbol.Assets) ([]symbol.AssetPrice, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	priceI, isOk := bodyI.([]map[string]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> error when casting bodyI to priceI")
	}

	assetsPricesArr := make([]symbol.AssetPrice, 0)
	assetsPricesMap := make(map[string]float64, 0)

	for _, assets := range priceI {
		symbolAsset, isOk := assets[pnames.Symbol].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting asset symbol to string")
		}

		priceStr, isOk := assets[pnames.Price].(string)
		if !isOk {
			return nil, errors.New("[mktdata] -> error when casting price to string")
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, errors.New("[mkdata] -> error error when parsing priceStr to float64")
		}

		assetsPricesMap[symbolAsset] = price
	}
	for i := 0; i < len(assetsList); i++ {
		assetListSymbol := assetsList[i].Base + assetsList[i].Quote
		for j, value := range assetsPricesMap {
			if j == assetListSymbol {
				assetsPricesArr = append(assetsPricesArr, symbol.AssetPrice{
					Symbol: assetsList[i],
					Price:  value,
				})
			}
		}
	}

	return assetsPricesArr, nil
}
