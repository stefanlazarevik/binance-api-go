package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/reqresp/paramnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"io/ioutil"
	"net/http"
	"strconv"
)

// responses json keys
const (
	codeKey = "code" // error code
	msgKey  = "msg"  // error message
)

func ParseGetOrderListResponse(response *http.Response) ([]exchangeapi.OrderInfo, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return nil, err
	}

	body, isOkay := bodyI.([]map[string]interface{})
	if !isOkay {
		return nil, errors.New("[binance-api-go.internal.parser] -> open order response is not an array")
	}

	ordersInfo := make([]exchangeapi.OrderInfo, 0)
	for idx, _ := range body {
		info, err := getOrderInfo(body[idx])
		if err != nil {
			return nil, err
		}
		ordersInfo = append(ordersInfo, info)
	}

	return ordersInfo, nil
}

func ParseSetOrderResponse(response *http.Response) (float64, error) {
	_, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}
	// TODO add set order parsing
	return 0, nil
}

func getResponseBody(response *http.Response) (interface{}, error) {
	if response.StatusCode/100 != 2 && response.Body == nil {
		return nil, &exchangeapi.ExchangeError{
			Type:    exchangeapi.HttpErr,
			Code:    response.StatusCode,
			Message: response.Status,
		}
	}

	respondBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if respondBody[0] == '[' && respondBody[len(respondBody)-1] == ']' {
		var body []map[string]interface{}
		err = json.Unmarshal(respondBody, &body)
		if err != nil {
			return nil, err
		}

		return body, err
	}

	var body map[string]interface{}
	err = json.Unmarshal(respondBody, &body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode/100 != 2 {
		return nil, parseBinanceError(body)
	}

	return body, nil
}

func parseBinanceError(body map[string]interface{}) error {
	code, isOkay := body[codeKey].(float64)
	if !isOkay {
		return errors.New("[binance-api-go.internal.parser] -> error code key not found")
	}

	message, isOkay := body[msgKey].(string)
	if !isOkay {
		return errors.New("[binance-api-go.internal.parser] -> failed to parse binance error message")
	}

	return &exchangeapi.ExchangeError{
		Type:    exchangeapi.BinanceErr,
		Code:    int(code),
		Message: message,
	}
}

func getOrderInfo(orderInfoJson map[string]interface{}) (exchangeapi.OrderInfo, error) {
	id, isOkay := orderInfoJson[paramnames.OrderIdParam].(float64)
	if !isOkay {
		return exchangeapi.OrderInfo{},
			errors.New(fmt.Sprint("[binance-api-go.internal.parser] -> error while parsing a value of ",
				paramnames.OrderIdParam, " key"))
	}

	status, isOkay := orderInfoJson[paramnames.StatusParam].(string)
	if !isOkay {
		return exchangeapi.OrderInfo{}, errors.New(fmt.Sprint("[binance-api-go.internal.parser] -> error while parsing a value of ",
			paramnames.StatusParam, " key"))
	}

	orderType, isOkay := orderInfoJson[paramnames.TypeParam].(string)
	if !isOkay {
		return exchangeapi.OrderInfo{}, errors.New(fmt.Sprint("[binance-api-go.internal.parser] -> error while parsing a value of ",
			paramnames.TypeParam, " key"))
	}

	priceStr, isOkay := orderInfoJson[paramnames.PriceParam].(string)
	if !isOkay {
		return exchangeapi.OrderInfo{}, errors.New(fmt.Sprint("[binance-api-go.internal.parser] -> error while parsing a value of ",
			paramnames.PriceParam, " key"))
	}

	// TODO Adopt quantity to limit orders
	//quantityStr, isOkay := orderInfoJson[paramnames.QuantityParam].(string)
	//if !isOkay {
	//	return exchangeapi.OrderInfo{}, errors.New(fmt.Sprint("[binance-api-go.internal.parser] -> error while parsing a value of ",
	//		paramnames.StatusParam, " key"))
	//}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return exchangeapi.OrderInfo{}, err
	}

	return exchangeapi.OrderInfo{
		Id:       fmt.Sprint(int(id)),
		Status:   orderStatusFromSting(status),
		Type:     orderTypeFromString(orderType),
		Price:    price,
		Quantity: 0,
	}, nil
}

func orderStatusFromSting(status string) exchangeapi.OrderStatus {
	switch status {
	case "NEW":
		return exchangeapi.NewOrder
	case "FILLED":
		return exchangeapi.Filled
	case "CANCELED":
		return exchangeapi.Canceled
	case "REJECTED":
		return exchangeapi.Rejected
	case "EXPIRED":
		return exchangeapi.Expired
	default:
		return exchangeapi.OtherStatus
	}
}

func orderTypeFromString(orderType string) exchangeapi.OrderType {
	switch orderType {
	case "LIMIT":
		return exchangeapi.Limit
	case "MARKET":
		return exchangeapi.Market
	default:
		return exchangeapi.OtherType
	}
}
