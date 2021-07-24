package main

import (
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"os"
)

func main() {
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	_, err := mgr.SetOrder(exchangeapi.OrderParameters{
		Symbol: exchangeapi.AssetsSymbol{
			Base:  "ETH",
			Quote: "USDT",
		},
		Side:     exchangeapi.Buy,
		Type:     exchangeapi.Market,
		Quantity: 12,
	})

	if err != nil {
		panic(err.Error())
	}
}
