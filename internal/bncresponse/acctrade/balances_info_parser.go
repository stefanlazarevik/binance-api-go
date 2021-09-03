package acctrade

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"net/http"
	"strconv"
)

func ParseBalancesInfo(response *http.Response, quote string) (float64, error) {
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return 0, err
	}
	balancesI, isOk := bodyI.(map[string]interface{})
	if isOk != true {
		return 0, errors.New("[bncresponse] -> Error when casting bodyI to balancesI")
	}
	balancesArrI, isOk := balancesI[pnames.Balances].([]interface{})
	if isOk != true {
		return 0, errors.New("[bncresponse] -> Error when casting balancesI to balancesArr")
	}
	var freeBalance float64

	//for _, value := range balancesArrI {
	//	inter, isOk := value.(map[string]interface{})
	//	if isOk != true {
	//		return 0, errors.New("[bncresponse] -> Error when casting value to inter ")
	//	}
	//	if inter[pnames.Asset].(string) == quote {
	//		free, isOk := inter[pnames.Free].(string)
	//		if isOk != true {
	//			return 0, errors.New("[bncresponse] -> Parsing free balance for"+quote+ "failed")
	//		}
	//		balance, err = strconv.ParseFloat(free, 64)
	//		if err != nil {
	//			return 0, errors.New("[bncresponse] ->  "+quote+" free balance conversion failed. Error: " + err.Error())
	//		}
	//		return balance, nil
	//	}
	//	continue
	//}
	for _, value := range balancesArrI {
		assetInfo, isOk := value.(map[string]interface{})
		if !isOk {
			continue
		}

		assetName, isOk := assetInfo[pnames.Asset].(string)
		if !isOk {
			continue
		}

		if assetName != quote {
			continue
		}

		freeBalanceStr, isOk := assetInfo[pnames.Free].(string)
		if !isOk {
			continue
		}

		freeBalance, err = strconv.ParseFloat(freeBalanceStr, 64)
		if err != nil {
			continue
		}
		return freeBalance, nil
	}
	return 0, errors.New("[bncresponse] -> Information about " + quote + " not found or its parsing failed.")
}
