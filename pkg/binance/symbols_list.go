package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

func (manager *ExchangeManager) GetSymbolsList() []symbol.Assets {
	var assetsArr []symbol.Assets
	for i := 0; i < len(manager.symbolsLimits); i++ {
		asset := manager.symbolsLimits[i].Assets
		assetsArr = append(assetsArr, asset)
	}

	return assetsArr
}

func (manager *ExchangeManager) StoreSymbolsLimits(limits []symbol.Limits) {
	manager.symbolsLimits = limits
}
