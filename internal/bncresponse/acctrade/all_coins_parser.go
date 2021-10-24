package acctrade

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
)

func ParseAllCoinsResponse(response *http.Response) ([]string, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}
	coinsI, isOkay := bodyI.([]map[string]interface{})
	if isOkay != true {
		return nil, errors.New("[bncresponse] -> Error when casting bodyI to coinsI")
	}
	coinArray := make([]string, 0)
	for i := 0; i < len(coinsI); i++ {
		coin, isOkay := coinsI[i][pnames.Coin].(string)
		if isOkay != true {
			return nil, errors.New("[bncresponse] -> Error when parsing coin to string")
		}
		coinArray = append(coinArray, coin)
	}
	return coinArray, nil
}
