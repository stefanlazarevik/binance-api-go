package bncresponse

import (
	"encoding/json"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"io/ioutil"
	"net/http"
)

func GetCandlestick(response *http.Response) (exchangeapi.Candlestick, error) {
	_, err := getResponseBody(response)
	if err != nil {
		return exchangeapi.Candlestick{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return exchangeapi.Candlestick{}, err
	}
	//var candleData Candlesticks
	var c exchangeapi.Candlestick
	var v []interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		return exchangeapi.Candlestick{}, err
	}

	// TODO rename field according to new struct
	//c.OpenTime, _ = v[0].(int64)
	//c.Open, _ = v[1].(string)
	//c.High, _ = v[2].(string)
	//c.Low, _ = v[3].(string)
	//c.Close, _ = v[4].(string)
	//c.Volume, _ = v[5].(string)
	//c.CloseTime, _ = v[6].(int64)
	//c.QuoteAssetVolume, _ = v[7].(string)
	//c.NumberOfTrade, _ = v[8].(int64)
	//c.TakerBuyBaseAssetVolume, _ = v[9].(string)
	//c.TakerBuyQuoteAssetVolume, _ = v[10].(string)
	//c.Ignore, _ = v[11].(string)

	return c, nil
}
