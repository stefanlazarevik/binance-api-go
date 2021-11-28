package acctrade

import (
	"errors"
	"fmt"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"net/http"
	"strconv"
	"time"
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
		info, err := retrieveOrderInfo(body[idx])
		if err != nil {
			return nil, err
		}
		ordersInfo = append(ordersInfo, info)
	}

	return ordersInfo, nil
}

func ParseOrderInfoResponse(response *http.Response) (order.Info, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return order.Info{}, err
	}

	body, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return order.Info{}, errors.New("[bncresponse] -> Set order response is not key/value pair array")
	}

	return retrieveOrderInfo(body)
}

func retrieveOrderInfo(body map[string]interface{}) (order.Info, error) {
	orderIdI, isOkay := body[pnames.OrderId]
	if !isOkay {
		return order.Info{}, errors.New("[bncresponse] -> Field `orderId` does not exist")
	}

	orderId, isOkay := orderIdI.(float64)
	if !isOkay {
		return order.Info{}, errors.New("[bncresponse] -> Field in casting interface to float64")
	}

	orderInfo := order.Info{
		Id: fmt.Sprint(int(orderId)),
	}

	var err error
	orderInfo.Type, err = getOrderType(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Status, err = getOrderStatus(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.Side, err = getOrderSide(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.TransactionTime, err = getTransactionTime(body)
	if err != nil {
		return order.Info{}, err
	}

	orderInfo.BaseQuantity, err = getOriginQuantity(body)
	if err != nil {
		return order.Info{}, err
	}

	if orderInfo.Status == order.Filled {
		orderInfo.Price = calculateFilledOrderPrice(body)
	}

	return orderInfo, nil
}

func calculateFilledOrderPrice(body map[string]interface{}) float64 {
	fills, isOkay := body[pnames.Fills].([]interface{})
	if !isOkay {
		return 0
	}

	orderPrice := 0.0
	for i := 0; i < len(fills); i++ {
		priceI, isOkay := fills[i].(map[string]interface{})
		if !isOkay {
			return 0
		}
		priceStr, isOkay := priceI[pnames.Price].(string)
		if !isOkay {
			return 0
		}
		priceF, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return 0
		}

		orderPrice += priceF
	}

	return orderPrice / float64(len(fills))
}

func getOriginQuantity(body map[string]interface{}) (float64, error) {
	origQtyStr, isOkay := body[pnames.OrigQty].(string)
	if !isOkay {
		return 0, errors.New("[bncresponse] -> Field `origQty` does not exist")
	}

	quantity, err := strconv.ParseFloat(origQtyStr, 64)
	if err != nil {
		return 0, errors.New("[bncresponse] -> error when parsing quantity to float64")
	}

	return quantity, nil
}

func getTransactionTime(body map[string]interface{}) (time.Time, error) {
	transactionTime, isOkay := body[pnames.TransactTime].(float64)
	if !isOkay {
		transactionTime, isOkay = body[pnames.UpdateTime].(float64)
		if !isOkay {
			return time.Time{}, errors.New("[bncresponse] -> Transaction time field does not exist")
		}

		return time.UnixMilli(int64(transactionTime)), nil
	}

	return time.UnixMilli(int64(transactionTime)), nil
}

func getOrderSide(body map[string]interface{}) (order.Side, error) {
	side, isOkay := body[pnames.Side].(string)
	if !isOkay {
		return order.UnknownSide, errors.New("[bncresponse] -> Field `side` does not exist")
	}

	switch side {
	case "BUY":
		return order.Buy, nil
	case "SELL":
		return order.Sell, nil
	default:
		return order.UnknownSide, nil
	}
}

func getOrderStatus(body map[string]interface{}) (order.Status, error) {
	status, isOkay := body[pnames.Status].(string)
	if !isOkay {
		return order.UnknownStatus, errors.New("[bncresponse] -> Field `status` does not exist")
	}

	switch status {
	case "NEW":
		return order.Open, nil
	case "FILLED":
		return order.Filled, nil
	case "PARTIALLY_FILLED":
		return order.PartiallyFilled, nil
	case "PENDING_CANCEL":
		return order.Canceled, nil
	case "CANCELED":
		return order.Canceled, nil
	case "REJECTED":
		return order.Rejected, nil
	case "EXPIRED":
		return order.Expired, nil
	default:
		return order.UnknownStatus, nil
	}
}

func getOrderType(body map[string]interface{}) (order.Type, error) {
	orderType, isOkay := body[pnames.Type].(string)
	if !isOkay {
		return order.UnknownType, errors.New("[bncresponse] -> Field `type` does not exist")
	}

	switch orderType {
	case "LIMIT":
		return order.Limit, nil
	case "MARKET":
		return order.Market, nil
	default:
		return order.UnknownType, nil
	}
}
