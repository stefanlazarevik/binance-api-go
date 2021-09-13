package main

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	cmn "github.com/posipaka-trade/posipaka-trade-cmn"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"os"
)

func main() {
	cmn.InitLoggers("binance-api-go")
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
	limits, _ := mgr.GetSymbolLimits()
	mgr.StoreSymbolLimits(limits)
	fmt.Println(mgr.GetSymbolsList())
}
