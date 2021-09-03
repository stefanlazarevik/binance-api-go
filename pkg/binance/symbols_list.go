package binance

import (
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse"
	"github.com/posipaka-trade/binance-api-go/internal/bncresponse/mktdata"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"net/http"
)

func (manager *ExchangeManager) UpdateSymbolsList() error {
	req, err := http.NewRequest(http.MethodGet, baseUrl+exchangeInfoEndpoint, nil)
	if err != nil {
		return err
	}

	response, err := manager.client.Do(req)
	if err != nil {
		return err
	}

	defer bncresponse.CloseBody(response)
	manager.symbolsList, err = mktdata.GetSymbolsList(response)
	if err != nil {
		return err
	}
	return nil
}

func (manager *ExchangeManager) GetSymbolsList() []symbol.Assets {
	return manager.symbolsList
}
