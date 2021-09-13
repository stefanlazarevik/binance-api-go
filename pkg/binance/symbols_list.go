package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

func (manager *ExchangeManager) GetSymbolsList() []symbol.Assets {
	limitsArr := manager.symbolsLimits

	var assetsArr []symbol.Assets

	for i := 0; i < len(limitsArr); i++ {
		asset := limitsArr[i].Assets
		assetsArr = append(assetsArr, asset)
	}

	return assetsArr
}

func (manager *ExchangeManager) StoreSymbolsLimits(limits []symbol.Limits) {

	manager.symbolsLimits = limits

}
