package main

import (
	"fmt"
	"github.com/posipaka-trade/binance-api-go/pkg/binance"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi"
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
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
	defer mgr.Finish()

	time.Sleep(5 * time.Second)

	startTime := time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(time.Since(startTime).String())

	time.Sleep(time.Second)
	startTime = time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(time.Since(startTime).String())

	time.Sleep(time.Second)
	startTime = time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(time.Since(startTime).String())

	time.Sleep(time.Second)
	startTime = time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})
	fmt.Println(time.Since(startTime).String())

	time.Sleep(time.Second)
	startTime = time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})

	time.Sleep(110 * time.Second)
	startTime = time.Now()
	_, _ = mgr.GetCurrentPrice(symbol.Assets{
		Base:  "ETH",
		Quote: "USDT",
	})

	fmt.Println(time.Since(startTime).String())
}
