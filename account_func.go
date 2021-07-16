package binance_api_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

var mgr exchangeapi.ApiConnector

func NewBinanceApi(exchangeapi.ApiKey) *BinanceApi {
	bObj := B{}
	bObj.aStr = "dsfg"


	return &BinanceApi{
		ApiKey:    cred.ApiKey,
		ApiSecret: cred.ApiSecret,
	}
}

func (binanceApi *BinanceApi) NewFullMarketOrder(orderParams OrdersParams) (float64, error) {
	var orderResult OrderResult
	endpoint := "/api/v3/order?"
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	var params string

	if orderParams.Side == Buy {
		params = fmt.Sprintf("symbol=%s&side=%s&type=%s&quoteOrderQty=%f&newOrderRespType=RESULT&timestamp=%s",
			orderParams.Symbol, orderParams.Side, orderParams.OrderType, orderParams.AssetCount, timestamp)
	} else {
		availableBalF := math.Floor(orderParams.AssetCount*100000) / 100000
		availableBalanceStr := fmt.Sprintf("%f", availableBalF)

		params = fmt.Sprintf("symbol=%s&side=%s&type=%s&quantity=%s&newOrderRespType=RESULT&timestamp=%s",
			orderParams.Symbol, orderParams.Side, orderParams.OrderType, availableBalanceStr, timestamp)
	}

	signature := binanceApi.MakeSignature(params)
	finalUrl := fmt.Sprintf("%s%s%s&signature=%s", burl, endpoint, params, signature)

	body, tradeBotError := binanceApi.DoRequest(Post, finalUrl)
	if tradeBotError != nil {
		return 0, tradeBotError
	}

	err := json.Unmarshal(body, &orderResult)
	if err != nil {
		return 0, err
	}
	//if string(body) == "{}" {        //Uncomment it if you want to test order
	//	return nil, freeBal
	//} else {
	//	return &cmn.TradeBotError{Msg: "Unreal"}, freeBal
	//}

	if orderResult.Status == filled { //	Uncomment it if you make a real order
		boughtCoin, err := strconv.ParseFloat(orderResult.OrigQty, 64)
		if err != nil {
			return 0, err
		}
		return boughtCoin, nil
	} else {
		return 0, &cmn.FakeBuyError{Message: "Unparsed orderResult struct"}
	}
}

func (binanceApi BinanceApi) GetPrice(symbol string) (float64, error) {
	var price float64
	pricesMap := map[string]string{}
	endPoint := "/api/v3/ticker/price?"
	params := fmt.Sprintf("symbol=%s", symbol)
	finalUrl := burl + endPoint + params

	body, tradeBotError := binanceApi.DoRequest(Get, finalUrl)
	if tradeBotError != nil {
		return 0, tradeBotError
	}

	err := json.Unmarshal(body, &pricesMap)
	if err != nil {
		return 0, err
	}
	priceStr := pricesMap["price"]
	price, err = strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func (binanceApi *BinanceApi) AccountInfo(symbol string) (float64, error) {
	timestampStr := strconv.Itoa(int(time.Now().Unix() * 1000))
	account := new(Account)
	params := "timestamp=" + timestampStr
	endpoint := "/api/v3/account?"

	signature := binanceApi.MakeSignature(params)

	finalUrl := fmt.Sprintf("%s%s%s&signature=%s", burl, endpoint, params, signature)

	body, tradeBotError := binanceApi.DoRequest(Get, finalUrl)
	if tradeBotError != nil {
		return 0, tradeBotError
	}

	err := json.Unmarshal(body, account)
	if err != nil {
		return 0, err
	}

	var freeMoney float64

	for _, a := range account.Balances {
		if a.Asset == symbol {
			freeMoney, err = strconv.ParseFloat(a.Free, 64)
			if err != nil {
				return 0, err
			}

		}
	}

	return freeMoney, nil
}

func (binanceApi *BinanceApi) Ping() error {
	endPoint := "/api/v3/ping"
	finalUrl := burl + endPoint

	_, tradeBotError := binanceApi.DoRequest(Get, finalUrl)
	if tradeBotError != nil {
		return tradeBotError
	}

	return nil
}
func (binanceApi *BinanceApi) GetOpenOrders(symbol string) (bool, error) {
	endPoint := "/api/v3/openOrders?"
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	params := fmt.Sprintf("symbol=%s&timestamp=%s", symbol, timestamp)
	signature := binanceApi.MakeSignature(params)
	finalUrl := fmt.Sprintf("%s%s%s&signature=%s", burl, endPoint, params, signature)

	body, tradeBotError := binanceApi.DoRequest(Get, finalUrl)
	if tradeBotError != nil {
		return false, tradeBotError
	}

	var isClosed bool
	if string(body) == "[]" {
		isClosed = true
	}

	return isClosed, nil
}
func (binanceApi *BinanceApi) NewLimitOrder(orderParams OrdersParams) (bool, error) {
	var orderResult OrderResult
	endpoint := "/api/v3/order?"
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	var params string

	if orderParams.Side == Buy {
		params = fmt.Sprintf("symbol=%s&side=%s&type=%s&quantity=%f&newOrderRespType=RESULT&timeInForce=%s&price=%f&timestamp=%s",
			orderParams.Symbol, orderParams.Side, orderParams.OrderType, orderParams.AssetCount, Gtc, orderParams.Price, timestamp)

	} else {
		params = fmt.Sprintf("symbol=%s&side=%s&type=%s&quantity=%f&newOrderRespType=RESULT&timeInForce=%s&price=%f&timestamp=%s",
			orderParams.Symbol, orderParams.Side, orderParams.OrderType, orderParams.AssetCount, Gtc, orderParams.Price, timestamp)
	}
	signature := binanceApi.MakeSignature(params)
	finalUrl := fmt.Sprintf("%s%s%s&signature=%s", burl, endpoint, params, signature)

	body, tradeBotError := binanceApi.DoRequest(Post, finalUrl)
	if tradeBotError != nil {
		return false, tradeBotError
	}
	err := json.Unmarshal(body, &orderResult)
	if err != nil {
		return false, err
	}

	var isOpened bool
	if orderResult.Status == "NEW" {
		isOpened = true
	} else {
		return false, &cmn.FakeBuyError{Message: "Unparsed orderResult struct"}
	}
	return isOpened, nil
}

func (binanceApi *BinanceApi) ExchangeInfo(symbol string) (ExchangeInfo, error) {
	requestParams := fmt.Sprintf("%s/api/v3/exchangeInfo?symbol=%s", burl, symbol)
	body, err := binanceApi.DoRequest(Get, requestParams)
	if err != nil {
		return ExchangeInfo{}, err
	}

	var respond map[string]interface{}
	err = json.Unmarshal(body, &respond)
	if err != nil {
		return ExchangeInfo{}, err
	}

	exchangeInfo, isOkay := parseExchangeInfoRespond(respond)
	if isOkay {
		return exchangeInfo, nil
	}

	return ExchangeInfo{}, errors.New("exchange info parse failed")
}
