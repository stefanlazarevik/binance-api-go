package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (binanceApi *BinanceApi) MakeSignature(baseUrl string) string {
	h := hmac.New(sha256.New, []byte(binanceApi.ApiSecret))
	h.Write([]byte(baseUrl))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}

func (binanceApi *BinanceApi) DoRequest(method string, finalUrl string) ([]byte, error) {
	req, err := http.NewRequest(method, finalUrl, nil)
	if err != nil {
		return nil, err
	}

	binanceApi.HeaderAdd(req)

	res, resErr := binanceApi.Client.Do(req)

	body, bodyErr := ioutil.ReadAll(res.Body)
	tradeBotError := TradeBotErrorCheck(body, res, resErr, bodyErr)
	if tradeBotError != nil {
		res.Body.Close()
		return nil, tradeBotError
	}

	return body, nil
}

func (binanceApi *BinanceApi) HeaderAdd(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-MBX-APIKEY", binanceApi.ApiKey)
}

func parseExchangeInfoRespond(respond map[string]interface{}) (ExchangeInfo, bool) {
	symbolsI, exist := respond["symbols"]
	if !exist {
		return ExchangeInfo{}, false
	}

	bytes, err := json.Marshal(symbolsI)
	if err != nil {
		return ExchangeInfo{}, false
	}

	var symbols []map[string]interface{}
	err = json.Unmarshal(bytes, &symbols)
	if err != nil {
		return ExchangeInfo{}, false
	}

	filtersI, exist := symbols[0]["filters"]
	if !exist {
		return ExchangeInfo{}, false
	}

	bytes, err = json.Marshal(filtersI)
	if err != nil {
		return ExchangeInfo{}, false
	}

	var filters []map[string]interface{}
	err = json.Unmarshal(bytes, &filters)
	if err != nil {
		return ExchangeInfo{}, false
	}

	var priceFilter PriceFilter
	var lotSize LotSize
	for _, filter := range filters {
		filterType, isOkay := filter["filterType"].(string)
		if !isOkay {
			continue
		}

		if filterType == priceFilterType {
			priceFilter, isOkay = parsePriceFilterJson(filter)
			if !isOkay {
				return ExchangeInfo{}, false
			}
		} else if filterType == lotSizeType {
			lotSize, isOkay = parseLotSizeJson(filter)
			if !isOkay {
				return ExchangeInfo{}, false
			}
		}
	}

	return ExchangeInfo{
		PriceFilter: priceFilter,
		LotSize:     lotSize,
	}, true
}

func parsePriceFilterJson(filter map[string]interface{}) (PriceFilter, bool) {
	minPrice, isOkay := filter["minPrice"].(string)
	if !isOkay {
		return PriceFilter{}, false
	}
	maxPrice, isOkay := filter["maxPrice"].(string)
	if !isOkay {
		return PriceFilter{}, false
	}
	tickSize, isOkay := filter["tickSize"].(string)
	if !isOkay {
		return PriceFilter{}, false
	}

	var priceFilter PriceFilter
	var err error
	priceFilter.MinPrice, err = strconv.ParseFloat(minPrice, 64)
	priceFilter.MaxPrice, err = strconv.ParseFloat(maxPrice, 64)
	priceFilter.TickSize, err = strconv.ParseFloat(tickSize, 64)
	if err != nil {
		return PriceFilter{}, false
	}
	return priceFilter, true
}

func parseLotSizeJson(filter map[string]interface{}) (LotSize, bool) {
	minQty, isOkay := filter["minQty"].(string)
	if !isOkay {
		return LotSize{}, false
	}
	maxQty, isOkay := filter["maxQty"].(string)
	if !isOkay {
		return LotSize{}, false
	}
	stepSize, isOkay := filter["stepSize"].(string)
	if !isOkay {
		return LotSize{}, false
	}

	var lotSize LotSize
	var err error
	lotSize.MinQuantity, err = strconv.ParseFloat(minQty, 64)
	lotSize.MaxQuantity, err = strconv.ParseFloat(maxQty, 64)
	lotSize.StepSize, err = strconv.ParseFloat(stepSize, 64)
	if err != nil {
		return LotSize{}, false
	}
	return lotSize, true
}
func (c *Candlesticks) UnmarshalJSON(data []byte) error {

	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	c.OpenTime, _ = v[0].(int64)
	c.Open, _ = v[1].(string)
	c.High, _ = v[2].(string)
	c.Low, _ = v[3].(string)
	c.Close, _ = v[4].(string)
	c.Volume, _ = v[5].(string)
	c.CloseTime, _ = v[6].(int64)
	c.QuoteAssetVolume, _ = v[7].(string)
	c.NumberOfTrade, _ = v[8].(int64)
	c.TakerBuyBaseAssetVolume, _ = v[9].(string)
	c.TakerBuyQuoteAssetVolume, _ = v[10].(string)
	c.Ignore, _ = v[11].(string)
	return nil
}
