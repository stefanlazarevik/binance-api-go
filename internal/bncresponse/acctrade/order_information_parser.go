package acctrade

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
)

func ParseOrderInfo(response *http.Response) (order.Info, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return order.Info{}, err
	}

	body, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return order.Info{}, errors.New("[bncresponse] -> Set order response is not key/value pair array")
	}

	return RetrieveSetOrderInfo(body)
}
