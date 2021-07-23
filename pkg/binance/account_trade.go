package binance

import (
	"encoding/json"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/reqresp/paramnames"
	"github.com/posipaka-trade/binance-api-go/internal/reqresp/parser"
	"github.com/posipaka-trade/binance-api-go/internal/sha256encryptor"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func (manager *ExchangeManager) GetOrdersList(symbol exchangeapi.AssetsSymbol) ([]exchangeapi.OrderInfo, error) {
	params := fmt.Sprintf("%s=%s%s&%s=%d", paramnames.SymbolParam, symbol.Base, symbol.Quote,
		paramnames.TimestampParam, time.Now().UnixNano()/int64(time.Millisecond))
	params = fmt.Sprintf("%s&%s=%s", params, paramnames.SignatureParam,
		sha256encryptor.EncryptMessage(params, manager.apiKey.Secret))

	request, err := http.NewRequest(http.MethodGet, fmt.Sprint(baseUrl, openOrdersEndpoint, "?", params), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-MBX-APIKEY", manager.apiKey.Key)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := manager.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	return parser.ParseGetOrderListResponse(response)
}

func (manager *ExchangeManager) SetOrder(parameters exchangeapi.OrderParameters) (float64, error) {
	requestBody := manager.createOrderRequestBody(&parameters)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprint(baseUrl, newOrderEndpoint), strings.NewReader(requestBody))
	if err != nil {
		return 0, err
	}

	request.Header.Set("X-MBX-APIKEY", manager.apiKey.Key)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := manager.client.Do(request)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	return parser.ParseSetOrderResponse(response)
}

func (manager *ExchangeManager) createOrderRequestBody(params *exchangeapi.OrderParameters) string {
	body := url.Values{}
	body.Add(paramnames.SymbolParam, fmt.Sprint(params.Symbol.Base, params.Symbol.Quote))
	body.Add(paramnames.SideParam, orderSideAlias[params.Side])
	body.Add(paramnames.TypeParam, orderTypeAlias[params.Type])

	if params.Type == exchangeapi.Limit {
		body.Add(paramnames.TimeInForceParam, "GTC")
		body.Add(paramnames.PriceParam, fmt.Sprint(params.Price))
		body.Add(paramnames.QuantityParam, fmt.Sprint(params.Quantity))
	} else if params.Type == exchangeapi.Market {
		body.Add(paramnames.QuantityParam, fmt.Sprint(params.Quantity))
	}

	body.Add(paramnames.SignatureParam, sha256encryptor.EncryptMessage(body.Encode(), manager.apiKey.Secret))
	return body.Encode()
}
func (manager *ExchangeManager) GetCurrentPrice(symbol exchangeapi.AssetsSymbol) (float64, error) {
	params := fmt.Sprintf("symbol=%s", symbol)

	response, err := manager.client.Get(fmt.Sprint(baseUrl, getPriceEndpoint, "?", params))
	if err != nil {
		return 0, err
	}
	body, err := ioutil.ReadAll(response.Body)
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
func (manager *ExchangeManager) GetCandlestick(symbol, interval string, limit int) ([]Candlesticks, error) {
	params := fmt.Sprintf("symbol=%s&interval=%s&limit=%d", symbol, interval, limit)

	response, err := manager.client.Get(fmt.Sprint(baseUrl, getCandlestickEndpoint, "?", params))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var candleData []Candlesticks

	err = json.Unmarshal(body, &candleData)
	if err != nil {
		return nil, err
	}

	return candleData, nil
}
