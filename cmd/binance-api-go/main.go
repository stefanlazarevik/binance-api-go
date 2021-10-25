package main

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/log"
	"os"
	"time"
)

func main() {
	log.Init("binance-api-go", true)
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	startTime := time.Now()
	for {
		time.Sleep(600 * time.Millisecond)
		_, err := mgr.GetSymbolsLimits()
		if err != nil {
			fmt.Println(err)
		}

		if time.Now().Sub(startTime) >= time.Minute {
			fmt.Println("Minute passed.")
			startTime = time.Now()
		}
	}
	//limits, _ := mgr.GetSymbolsLimits()
	//mgr.StoreSymbolsLimits(limits)
	//fmt.Println(mgr.GetAssetBalance("ETH"))
	//fmt.Println(mgr.SetOrder(order.Parameters{
	//	Assets: symbol.Assets{
	//		Base:  "ETH",
	//		Quote: "USDT",
	//	},
	//	Side:     order.Buy,
	//	Type:     order.Market,
	//	Quantity: 1000,
	//	Price:    0,
	//}))
	//coins, _ := mgr.GetAllCoinsInfo()
	//fmt.Println(len(coins))
	//fmt.Println(coins)
}
