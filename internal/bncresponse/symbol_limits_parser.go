package bncresponse

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/pnames"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
	"strconv"
)

const (
	priceFilterType = "PRICE_FILTER"
	lotSizeType     = "LOT_SIZE"
)

func GetSymbolLimits(response *http.Response) (symbol.Limits, error) {
	bodyI, err := getResponseBody(response)
	if err != nil {
		return symbol.Limits{}, err
	}

	symbolInfo, err := parseSymbolInfo(bodyI)
	if err != nil {
		return symbol.Limits{}, err
	}

	filters, isOk := symbolInfo[pnames.Filters].([]interface{})
	if !isOk {
		return symbol.Limits{}, errors.New("[bncresponse] -> filters tree casting failed")
	}

	limits, err := parseSymbolFilters(filters)
	if err != nil {
		return symbol.Limits{}, err
	}

	limits.Assets.Base, isOk = symbolInfo[pnames.BaseAsset].(string)
	if !isOk {
		return symbol.Limits{}, errors.New("[bncresponse] -> failed base asset parsing")
	}

	limits.Assets.Quote, isOk = symbolInfo[pnames.QuoteAsset].(string)
	if !isOk {
		return symbol.Limits{}, errors.New("[bncresponse] -> failed quote asset parsing")
	}

	return limits, nil
}

func parseSymbolInfo(bodyI interface{}) (map[string]interface{}, error) {
	body, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return nil, errors.New("[bncresponse] -> error when casting bodyI to timeI")
	}

	symbols, isOk := body[pnames.Symbols].([]interface{})
	if !isOk {
		return nil, errors.New("[bncresponse] -> symbols tree not found in exchange information")
	}

	if len(symbols) == 0 {
		return nil, errors.New("[bncresponse] -> no symbols were returned by exchange information response")
	}

	symbolInfo, isOk := symbols[0].(map[string]interface{})
	if !isOk {
		return nil, errors.New("[bncresponse] -> failed to cast symbol to key/value pairs")
	}

	return symbolInfo, nil
}

func parseSymbolFilters(filters []interface{}) (symbol.Limits, error) {
	var limits symbol.Limits
	for _, filterI := range filters {
		filter, isOk := filterI.(map[string]interface{})
		if !isOk {
			continue
		}

		var err error
		filterType, isOk := filter[pnames.FilterType].(string)
		if filterType == priceFilterType {
			limits, err = parsePriceFilterJson(filter, limits)
			if err != nil {
				return symbol.Limits{}, err
			}
		} else if filterType == lotSizeType {
			limits, err = parseLotSize(filter, limits)
			if err != nil {
				return symbol.Limits{}, err
			}
		}
	}
	return limits, nil
}

func parsePriceFilterJson(filter map[string]interface{}, limits symbol.Limits) (symbol.Limits, error) {
	minPrice, isOkay := filter[pnames.MinPrice].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> min price for price filter not parsed")
	}

	maxPrice, isOkay := filter[pnames.MaxPrice].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> max price for price filter not parsed")
	}

	tickSize, isOkay := filter[pnames.TickSize].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> tick size for price filter not parsed")
	}

	var err error
	limits.Price.MinSize, err = strconv.ParseFloat(minPrice, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> MinPrice. " + err.Error())
	}

	limits.Price.MaxSize, err = strconv.ParseFloat(maxPrice, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> MaxPrice. " + err.Error())
	}

	limits.Price.Increment, err = strconv.ParseFloat(tickSize, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> tickSize. " + err.Error())
	}

	return limits, nil
}

func parseLotSize(filter map[string]interface{}, limits symbol.Limits) (symbol.Limits, error) {
	minQty, isOkay := filter[pnames.MinQuantity].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> min quantity filter not parsed")
	}

	maxQty, isOkay := filter[pnames.MaxQuantity].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> max quantity filter not parsed")
	}

	stepSize, isOkay := filter[pnames.StepSize].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[bncresponse] -> quantity step size filter not parsed")
	}

	var err error
	limits.Base.MinSize, err = strconv.ParseFloat(minQty, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> MinQty. " + err.Error())
	}

	limits.Base.MaxSize, err = strconv.ParseFloat(maxQty, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> MaxQty. " + err.Error())
	}

	limits.Base.Increment, err = strconv.ParseFloat(stepSize, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[bncresponse] -> StepSize. " + err.Error())
	}

	return limits, nil
}
