package binance

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

//func (manager *ExchangeManager) UpdateSymbolsList() error {
//	req, err := http.NewRequest(http.MethodGet, baseUrl+exchangeInfoEndpoint, nil)
//	if err != nil {
//		return err
//	}
//
//	response, err := manager.client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	defer bncresponse.CloseBody(response)
//	manager.symbolsList, err = mktdata.GetSymbolsList(response)
//	if err != nil {
//		return err
//	}
//	return nil
//}

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
