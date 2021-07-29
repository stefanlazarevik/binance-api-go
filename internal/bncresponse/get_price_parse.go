package bncresponse

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"strconv"
)

func GetCurrentPrice(response *http.Response) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}

	priceI, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return 0, errors.New("[bncresponse] -> error when casting bodyI to priceI")
	}

	priceStr := priceI[pnames.Price]
	price, err := strconv.ParseFloat(priceStr.(string), 64)
	if err != nil {
		return 0, errors.New("[bncresponse] -> error when parsing priceStr to float64")
	}
	return price, nil
}
