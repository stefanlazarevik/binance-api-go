package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"strconv"
)

func GetAssetOrderBook(response *http.Response) (symbol.OrderBook, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return symbol.OrderBook{}, err
	}

	orderBookI, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return symbol.OrderBook{}, errors.New("[mktdata] -> error when casting order book body to map[string]interface")
	}

	asksIArr, isOk := orderBookI[pnames.Asks].([]interface{})
	if isOk != true {
		return symbol.OrderBook{}, errors.New("[mktdata] -> error when casting asks to []interface")
	}

	bidsIArr, isOk := orderBookI[pnames.Bids].([]interface{})
	if isOk != true {
		return symbol.OrderBook{}, errors.New("[mktdata] -> error when casting bids to []interface")
	}

	var orderBook symbol.OrderBook

	asksArr, err := getAssetAsk(asksIArr)
	if err != nil {
		return symbol.OrderBook{}, err
	}

	bidsArr, err := getAssetBid(bidsIArr)
	if err != nil {
		return symbol.OrderBook{}, err
	}

	orderBook.Ask = asksArr
	orderBook.Bid = bidsArr

	return orderBook, nil
}

func getAssetAsk(asksIArr []interface{}) ([]symbol.BidAsk, error) {
	askArr := make([]symbol.BidAsk, len(asksIArr))

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

		askArr[i] = symbol.BidAsk{
			Price:    price,
			Quantity: quantity,
		}
	}
	return askArr, nil
}

func getAssetBid(bidsIArr []interface{}) ([]symbol.BidAsk, error) {
	bidsArr := make([]symbol.BidAsk, len(bidsIArr))

	for i, value := range bidsIArr {
		bidsI, isOk := value.([]interface{})
		if isOk != true {
			return nil, errors.New("[mktdata] -> error when casting order book bids to []interface")
		}

		priceStr, isOk := bidsI[0].(string)
		if isOk != true {
			return nil, errors.New("[mktdata] -> error when casting order book price to string")
		}

		quantityStr, isOk := bidsI[1].(string)
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

		bidsArr[i] = symbol.BidAsk{
			Price:    price,
			Quantity: quantity,
		}
	}
	return bidsArr, nil
}
