package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
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
