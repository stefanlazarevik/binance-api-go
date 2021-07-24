package main

import (
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/order"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
	"os"
)

func main() {
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    os.Args[1],
		Secret: os.Args[2],
	})

	_, err := mgr.SetOrder(order.Parameters{
		Assets: symbol.Assets{
			Base:  "ETH",
			Quote: "USDT",
		},
		Side:     order.Buy,
		Type:     order.Market,
		Quantity: 12,
	})

	if err != nil {
		panic(err.Error())
	}
}
