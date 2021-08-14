package bncresponse

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func ParseSetOrder(response *http.Response) (float64, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return 0, err
	}

	responseI, isOkay := bodyI.(map[string]interface{})
	if !isOkay {
		return 0, errors.New("[bncresponse] -> set order response is not key/value pair array")
	}

	origQtyI := responseI["origQty"]
	origQtyStr := fmt.Sprintf("%v", origQtyI)
	origQtyF, err := strconv.ParseFloat(origQtyStr, 64)
	if err != nil {
		return 0, errors.New("[bncresponse] -> Error when parsing origQtyStr to float64")

	}
	return origQtyF, nil
}
