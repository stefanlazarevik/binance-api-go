package main

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"os"
)

func main() {
	log.Init("binance-api-go", true)
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	//err := mgr.UpdateSymbolsList()
	//if err != nil {
	//	panic(err)
	//}
	//
	//symbols := mgr.GetSymbolsList()
	//fmt.Print(symbols)
	limits, _ := mgr.GetSymbolsLimits()
	mgr.StoreSymbolsLimits(limits)
	fmt.Println(mgr.GetSymbolsList())
}
