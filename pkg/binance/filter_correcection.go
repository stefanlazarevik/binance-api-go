package binance

import (
	"fmt"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"math"
	"strconv"
)

const accuracyFactor = 1000000000

func (manager *ExchangeManager) applyFilter(parameters order.Parameters) order.Parameters {
	limitsIdx := -1
	for idx, symbolLimits := range manager.symbolsLimits {
		if symbolLimits.Assets.IsEqual(parameters.Assets) {
			limitsIdx = idx
			break
		}
	}

	if limitsIdx == -1 {
		log.Warning.Printf("[binance] -> No limits found for %s/%s symbol",
			parameters.Assets.Base, parameters.Assets.Quote)
		return parameters
	}

	if parameters.Type == order.Limit {
		parameters.Price = filterCorrector(parameters.Price, manager.symbolsLimits[limitsIdx].Price)
		//log.Info.Printf("[binance] -> Price value correction. New price %f", parameters.Price)
	}

	if parameters.Side == order.Sell {
		parameters.Quantity = filterCorrector(parameters.Quantity, manager.symbolsLimits[limitsIdx].Base)
		parameters.Quantity, _ = strconv.ParseFloat(fmt.Sprintf("%."+
			fmt.Sprint(manager.symbolsLimits[limitsIdx].Base.Precision)+"f", parameters.Quantity), 64)
	} else {
		parameters.Quantity, _ = strconv.ParseFloat(fmt.Sprintf("%."+
			fmt.Sprint(manager.symbolsLimits[limitsIdx].Quote.Precision)+"f", parameters.Quantity), 64)
	}

	//log.Info.Printf("[binance] -> Quantity value correction on selling. New quantity %f", parameters.Quantity)

	return parameters
}

func filterCorrector(value float64, detail symbol.LimitDetail) float64 {
	valueInt := int(math.Round(value * accuracyFactor))
	minValueInt := int(math.Round(detail.MinSize * accuracyFactor))
	incrementInt := int(math.Round(detail.Increment * accuracyFactor))
	if ((valueInt - minValueInt) % incrementInt) != 0 {
		valueInt -= (valueInt - minValueInt) % incrementInt
		value = float64(valueInt) / accuracyFactor
	}

	if value > detail.MaxSize {
		value = detail.MaxSize
	}

	if value < detail.MinSize {
		value = detail.MinSize
	}

	return value
}
