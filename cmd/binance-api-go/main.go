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

	//limits, err := mgr.GetSymbolLimits(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//mgr.AddLimits(limits)
	//_, err = mgr.SetOrder(order.Parameters{
	//	Assets: symbol.Assets{
	//		Base:  "ETH",
	//		Quote: "USDT",
	//	},
	//	Price:    3120.58789214685,
	//	Side:     order.Buy,
	//	Type:     order.Limit,
	//	Quantity: 0.00473815,
	//})
	//price, err := mgr.GetCurrentPrice(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT"})
	//fmt.Println(price)
	//candleStick, err := mgr.GetCandlestick(symbol.Assets{Base: "ETH", Quote: "USDT"}, "1h", 1)
	//fmt.Println(candleStick)
	//if err != nil {
	//	panic(err.Error())
	//}

	//_, err := mgr.GetSymbolLimits(symbol.Assets{
	//	Base:  "ETH",
	//	Quote: "USDT",
	//})
	balance, err := mgr.BalancesInfo("USDT")
	fmt.Println(balance)
	if err != nil {
		panic(err.Error())
	}
}
