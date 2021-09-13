package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

func (manager *ExchangeManager) GetSymbolsList() ([]symbol.Assets, error) {
	limitsArr, err := manager.GetSymbolLimits()
	if err != nil {
		return []symbol.Assets{}, err
	}

	var assetsArr []symbol.Assets

	for i := 0; i < len(limitsArr); i++ {
		asset := limitsArr[i].Assets
		assetsArr = append(assetsArr, asset)
	}

	return assetsArr, nil
}

func (manager *ExchangeManager) StoreSymbolLimits(limits []symbol.Limits) {

	manager.symbolsLimits = limits

}
