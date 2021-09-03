package mktdata

import (
	"errors"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
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
	bodyI, err := bncresponse.GetResponseBody(response)
	if err != nil {
		return symbol.Limits{}, err
	}

	symbolInfo, err := parseSymbolInfo(bodyI)
	if err != nil {
		return symbol.Limits{}, err
	}

	filters, isOk := symbolInfo[pnames.Filters].([]interface{})
	if !isOk {
		return symbol.Limits{}, errors.New("[mktdata] -> filters tree casting failed")
	}

	limits, err := parseSymbolFilters(filters)
	if err != nil {
		return symbol.Limits{}, err
	}

	limits.Assets, err = parseAsset(symbolInfo)
	if err != nil {
		return symbol.Limits{}, err
	}

	limits, err = parseSymbolPrecision(symbolInfo, limits)
	if err != nil {
		return symbol.Limits{}, err
	}

	return limits, nil
}

func parseSymbolPrecision(symbolInfo map[string]interface{}, limits symbol.Limits) (symbol.Limits, error) {
	basePrecision, isOkay := symbolInfo[pnames.BaseAssetPrecision].(float64)
	if !isOkay {
		return limits, errors.New("[mktdata] -> failed base asset parsing")
	}

	quotePrecision, isOkay := symbolInfo[pnames.QuoteAssetPrecision].(float64)
	if !isOkay {
		return limits, errors.New("[mktdata] -> failed base asset parsing")
	}

	limits.Base.Precision = int(basePrecision)
	limits.Quote.Precision = int(quotePrecision)
	return limits, nil
}

func parseAsset(symbolInfo map[string]interface{}) (symbol.Assets, error) {
	var assets symbol.Assets
	var isOkay bool
	assets.Base, isOkay = symbolInfo[pnames.BaseAsset].(string)
	if !isOkay {
		return symbol.Assets{}, errors.New("[mktdata] -> failed base asset parsing")
	}

	assets.Quote, isOkay = symbolInfo[pnames.QuoteAsset].(string)
	if !isOkay {
		return symbol.Assets{}, errors.New("[mktdata] -> failed quote asset parsing")
	}

	return assets, nil
}

func parseSymbolInfo(bodyI interface{}) (map[string]interface{}, error) {
	body, isOk := bodyI.(map[string]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> error when casting bodyI to timeI")
	}

	symbols, isOk := body[pnames.Symbols].([]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> symbols tree not found in exchange information")
	}

	if len(symbols) == 0 {
		return nil, errors.New("[mktdata] -> no symbols were returned by exchange information response")
	}

	symbolInfo, isOk := symbols[0].(map[string]interface{})
	if !isOk {
		return nil, errors.New("[mktdata] -> failed to cast symbol to key/value pairs")
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
		return symbol.Limits{}, errors.New("[mktdata] -> min price for price filter not parsed")
	}

	maxPrice, isOkay := filter[pnames.MaxPrice].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[mktdata] -> max price for price filter not parsed")
	}

	tickSize, isOkay := filter[pnames.TickSize].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[mktdata] -> tick size for price filter not parsed")
	}

	var err error
	limits.Price.MinSize, err = strconv.ParseFloat(minPrice, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> MinPrice. " + err.Error())
	}

	limits.Price.MaxSize, err = strconv.ParseFloat(maxPrice, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> MaxPrice. " + err.Error())
	}

	limits.Price.Increment, err = strconv.ParseFloat(tickSize, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> tickSize. " + err.Error())
	}

	return limits, nil
}

func parseLotSize(filter map[string]interface{}, limits symbol.Limits) (symbol.Limits, error) {
	minQty, isOkay := filter[pnames.MinQuantity].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[mktdata] -> min quantity filter not parsed")
	}

	maxQty, isOkay := filter[pnames.MaxQuantity].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[mktdata] -> max quantity filter not parsed")
	}

	stepSize, isOkay := filter[pnames.StepSize].(string)
	if !isOkay {
		return symbol.Limits{}, errors.New("[mktdata] -> quantity step size filter not parsed")
	}

	var err error
	limits.Base.MinSize, err = strconv.ParseFloat(minQty, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> MinQty. " + err.Error())
	}

	limits.Base.MaxSize, err = strconv.ParseFloat(maxQty, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> MaxQty. " + err.Error())
	}

	limits.Base.Increment, err = strconv.ParseFloat(stepSize, 64)
	if err != nil {
		return symbol.Limits{}, errors.New("[mktdata] -> StepSize. " + err.Error())
	}

	return limits, nil
}
