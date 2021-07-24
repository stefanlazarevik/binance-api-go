package bncresponse

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCurrentPrice(response *http.Response) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}

	body, err := json.Marshal(bodyI)
	if err != nil {
		return 0, err
	}

	pricesMap := map[string]string{}
	err = json.Unmarshal(body, &pricesMap)
	if err != nil {
		return 0, err
	}

	priceStr := pricesMap["price"]
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
}
