package main

import (
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
)

func main() {
	mgr := binance.New(exchangeapi.ApiKey{
		Key:    "gYCpJ8cHY9aS09qnBcktaG2WB44BwiRF3nmNQQTkDGHTC39Zm5CSeVbv7MF5sIDL", //os.Args[1],
		Secret: "1x6qZCz7tJnvWFyJJ0d40nhPneA6SV8U9arMxXr5lPX68zETlnaKm5XJCuGl9Ljy", //os.Args[2],
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
