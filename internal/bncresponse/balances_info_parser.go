package bncresponse

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"strconv"
)

func ParseBalancesInfo(response *http.Response, quote string) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}
	balancesI, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return 0, errors.New("[bncresponse] -> Error when casting bodyI to balancesI in ParseBalancesInfo")
	}
	balancesArrI, isOk := balancesI[pnames.Balances].([]interface{})
	if isOk != true {
		return 0, errors.New("[bncresponse] -> Error when casting balancesI to balancesArr in ParseBalancesInfo")
	}
	var balance float64

	for _, value := range balancesArrI {
		inter, isOk := value.(map[string]interface{})
		if isOk != true {
			return 0, errors.New("[bncresponse] -> Error when casting value to inter in ParseBalancesInfo")
		}
		if inter[pnames.Asset].(string) == quote {
			free, isOk := inter[pnames.Free].(string)
			if isOk != true {
				return 0, errors.New("[bncresponse] -> Error when casting Free to free in ParseBalancesInfo")
			}
			balance, err = strconv.ParseFloat(free, 64)
			if err != nil {
				return 0, errors.New("[bncresponse] -> Error when parsing Free to balance in ParseBalancesInfo")
			}
			return balance, nil
		}
		continue
	}

	return balance, nil
}
