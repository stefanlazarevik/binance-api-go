package bncresponse

import (
	"errors"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"net/http"
	"strconv"
	"time"
)

func GetCandlestick(response *http.Response) ([]exchangeapi.Candlestick, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return nil, err
	}

	g := bodyI.([]interface{})
	var candleStickArr []exchangeapi.Candlestick

	for i := 0; i < len(g); i++ {

		v := g[i].([]interface{})
		var c exchangeapi.Candlestick

		openPrice, err := strconv.ParseFloat(v[1].(string), 64)
		if err != nil {
			return nil, errors.New("[bncresponse] -> error when parsing openPrice to float64")
		}
		maxPrice, err := strconv.ParseFloat(v[2].(string), 64)
		if err != nil {
			return nil, errors.New("[bncresponse] -> error when parsing maxPrice to float64")
		}
		minPrice, err := strconv.ParseFloat(v[3].(string), 64)
		if err != nil {
			return nil, errors.New("[bncresponse] -> error when parsing minPrice to float64")
		}
		closePrice, err := strconv.ParseFloat(v[4].(string), 64)
		if err != nil {
			return nil, errors.New("[bncresponse] -> error when parsing closePrice to float64")
		}
		volume, err := strconv.ParseFloat(v[5].(string), 64)
		if err != nil {
			return nil, errors.New("[bncresponse] -> error when parsing volume to float64")
		}
		openTimeF, isOk := v[0].(float64)
		if !isOk {
			return nil, errors.New("[bncresponse] -> error when parsing openTime to float64")
		}
		closeTimeF, isOk := v[6].(float64)
		if !isOk {
			return nil, errors.New("[bncresponse] -> error when parsing closeTime to float64")
		}
		tradesNumber, isOk := v[8].(float64)
		if !isOk {
			return nil, errors.New("[bncresponse] -> error when parsing tradesNumber to float64")
		}

		c.OpenTime = time.Unix(int64(openTimeF), 999999999)
		c.OpenPrice = openPrice
		c.MaxPrice = maxPrice
		c.MinPrice = minPrice
		c.ClosePrice = closePrice
		c.Volume = volume
		c.CloseTime = time.Unix(int64(closeTimeF), 999999999)
		c.TradesNumber = tradesNumber

		candleStickArr = append(candleStickArr, c)
	}

	return candleStickArr, nil
}
