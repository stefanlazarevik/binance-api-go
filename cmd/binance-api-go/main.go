package main

import (
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"os"
)

func main() {
	log.Init("binance-api-go", true)
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	limits, err := mgr.GetSymbolsLimits()
	mgr.StoreSymbolsLimits(limits)
	if err != nil {
		panic(err)
	}

	_, err = mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "TROY",
			Quote: "BUSD",
		},
		Side:     order.Sell,
		Type:     order.Limit,
		Quantity: 1075.494,
		Price:    0.018588 * 1.1,
	})

	if err != nil {
		panic(err)
	}
	//err := mgr.UpdateSymbolsList()
	//if err != nil {
	//	panic(err)
	//}
	//
	//symbols := mgr.GetSymbolsList()
	//fmt.Print(symbols)
}
