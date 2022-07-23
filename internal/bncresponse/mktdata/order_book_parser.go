package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"strconv"
)

func GetAssetOrderBook(response *http.Response) ([]symbol.OrderBook, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	orderBookI, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return nil, errors.New("[mktdata] -> error when casting order book body to map[string]interface")
	}

	asksIArr, isOk := orderBookI[pnames.Asks].([]interface{})
	if isOk != true {
		return nil, errors.New("[mktdata] -> error when casting asks to []interface")
	}

	orderBook := make([]symbol.OrderBook, len(asksIArr))

	for i, value := range asksIArr {
		askI, isOk := value.([]interface{})
		if isOk != true {
			return nil, errors.New("[mktdata] -> error when casting order book body to map[string]interface")
		}

		priceStr, isOk := askI[0].(string)
		if isOk != true {
			return nil, errors.New("[mktdata] -> error when casting order book price to string")
		}

		quantityStr, isOk := askI[1].(string)
		if isOk != true {
			return nil, errors.New("[mktdata] -> error when casting order book quantity to string")
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, err
		}

		quantity, err := strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			return nil, err
		}

		orderBook[i] = symbol.OrderBook{
			Price:    price,
			Quantity: quantity,
		}
	}

	return orderBook, nil
}
