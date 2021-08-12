package main

import (
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"os"
)

func main() {
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	//_, err := mgr.SetOrder(exchangeapi.OrderParameters{
	//	Symbol: exchangeapi.AssetsSymbol{
	//		Base:  "ETH",
	//		Quote: "USDT",
	//	},
	//	Side:     exchangeapi.Buy,
	//	Type:     exchangeapi.Market,
	//	Quantity: 12,
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

	_, err := mgr.GetSymbolLimits(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	if err != nil {
		panic(err.Error())
	}
}
