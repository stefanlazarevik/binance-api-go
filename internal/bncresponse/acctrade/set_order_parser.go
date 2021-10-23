package acctrade

import (
	"errors"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"strconv"
)

type OrderAnswer struct {
	Price    float64
	Quantity float64
}

func ParseSetOrder(response *http.Response) (OrderAnswer, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return OrderAnswer{}, err
	}

	responseI, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return OrderAnswer{}, errors.New("[bncresponse] -> set order response is not key/value pair array")
	}

	return prepareOrderAnswer(responseI)
}
func prepareOrderAnswer(responseI map[string]interface{}) (OrderAnswer, error) {
	var orderAnswer OrderAnswer
	var err error

	fills, isOkay := responseI[pnames.Fills].([]interface{})
	if !isOkay {
		return OrderAnswer{}, errors.New("[bncresponse] -> no such key as fills")
	}
	for i := 0; i < len(fills); i++ {
		priceI, isOkay := fills[i].(map[string]interface{})
		if !isOkay {
			return OrderAnswer{}, errors.New("[bncresponse] -> error when casting fills value to array")
		}
		priceStr, isOkay := priceI[pnames.Price].(string)
		if !isOkay {
			return OrderAnswer{}, errors.New("[bncresponse] -> no such key as price")
		}
		priceF, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return OrderAnswer{}, errors.New("[bncresponse] -> error when parsing priceStr to float64")
		}
		orderAnswer.Price += priceF
	}
	orderAnswer.Price = orderAnswer.Price / float64(len(fills))

	origQtyI := responseI[pnames.OrigQty]
	origQtyStr := fmt.Sprintf("%v", origQtyI)

	orderAnswer.Quantity, err = strconv.ParseFloat(origQtyStr, 64)
	if err != nil {
		return OrderAnswer{}, errors.New("[bncresponse] -> error when parsing origQtyStr to float64")
	}

	return orderAnswer, nil
}
