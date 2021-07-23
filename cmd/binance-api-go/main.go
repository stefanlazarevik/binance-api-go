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

	_, err := mgr.GetOrdersList(exchangeapi.AssetsSymbol{
		Base:  "ETH",
		Quote: "BUSD",
	})

	if err != nil {
		panic(err.Error())
	}
}
