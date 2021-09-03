package acctrade

import (
	"errors"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
	"strconv"
)

func ParseGetOrderList(response *http.Response) ([]order.Info, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return nil, err
	}

	body, isOkay := bodyI.([]map[string]interface{})
	if !isOkay {
		return nil, errors.New("[bncresponse] -> open order response is not an array")
	}

	ordersInfo := make([]order.Info, 0)
	for idx, _ := range body {
		info, err := getOrderInfo(body[idx])
		if err != nil {
			return nil, err
		}
		ordersInfo = append(ordersInfo, info)
	}

	return ordersInfo, nil
}

func getOrderInfo(orderInfoJson map[string]interface{}) (order.Info, error) {
	id, isOkay := orderInfoJson[pnames.OrderId].(float64)
	if !isOkay {
		return order.Info{},
			errors.New(fmt.Sprint("[bncresponse] -> error while parsing a value of ",
				pnames.OrderId, " key"))
	}

	status, isOkay := orderInfoJson[pnames.Status].(string)
	if !isOkay {
		return order.Info{}, errors.New(fmt.Sprint("[bncresponse] -> error while parsing a value of ",
			pnames.Status, " key"))
	}

	orderType, isOkay := orderInfoJson[pnames.Type].(string)
	if !isOkay {
		return order.Info{}, errors.New(fmt.Sprint("[bncresponse] -> error while parsing a value of ",
			pnames.Type, " key"))
	}

	priceStr, isOkay := orderInfoJson[pnames.Price].(string)
	if !isOkay {
		return order.Info{}, errors.New(fmt.Sprint("[bncresponse] -> error while parsing a value of ",
			pnames.Price, " key"))
	}

	// TODO Adopt quantity to limit orders
	//quantityStr, isOkay := orderInfoJson[paramnames.QuantityParam].(string)
	//if !isOkay {
	//	return order.Info{}, errors.New(fmt.Sprint("[bncresponse] -> error while parsing a value of ",
	//		paramnames.StatusParam, " key"))
	//}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return order.Info{}, err
	}

	return order.Info{
		Id:     fmt.Sprint(int(id)),
		Status: orderStatusFromSting(status),
		Type:   orderTypeFromString(orderType),
		Price:  price,
	}, nil
}

func orderStatusFromSting(status string) order.Status {
	switch status {
	case "NEW":
		return order.New
	case "FILLED":
		return order.Filled
	case "CANCELED":
		return order.Canceled
	case "REJECTED":
		return order.Rejected
	case "EXPIRED":
		return order.Expired
	default:
		return order.OtherStatus
	}
}

func orderTypeFromString(orderType string) order.Type {
	switch orderType {
	case "LIMIT":
		return order.Limit
	case "MARKET":
		return order.Market
	default:
		return order.OtherType
	}
}
