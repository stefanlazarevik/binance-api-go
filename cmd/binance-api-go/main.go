package main

import (
	"github.com/posipaka-trade/posipaka-trade-cmn/exchangeapi/symbol"
)

func main() {

	//for {
	//	marginPriceMap, err := mgr.GetPricesMap(binance.BaseMarginUrl, binance.MarginGetPriceEndpoint)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	spotPriceMap, err := mgr.GetPricesMap(binance.BaseUrl, binance.GetPriceEndpoint)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	for marginSymbol, marginPrice := range marginPriceMap {
	//		spotPrice := spotPriceMap[marginSymbol]
	//		if spotPrice == 0 {
	//			continue
	//		}
	//
	//		growPercent := (marginPrice - spotPrice) / spotPrice * 100
	//
	//		if growPercent > 0.3 {
	//			log.Info.Printf("\nMargin price is bigger\nSymbol = %s\nSpot price = %f\nMargin price = %f\nGrow percent = %f", marginSymbol, spotPrice, marginPrice, growPercent)
	//		}
	//		if growPercent < -0.3 {
	//			log.Info.Printf("\nSpot price is bigger\nSymbol = %s\nSpot price = %f\nMargin price = %f\nGrow percent = %f", marginSymbol, spotPrice, marginPrice, growPercent)
	//		}
	//	}
	//	time.Sleep(1000 * time.Millisecond)
	//}

	//startTime := time.Now()
	//for {
	//	time.Sleep(600 * time.Millisecond)
	//	_, err := mgr.GetSymbolsLimits()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	if time.Now().Sub(startTime) >= time.Minute {
	//		fmt.Println("Minute passed.")
	//		startTime = time.Now()
	//	}
	//}
	//limits, _ := mgr.GetSymbolsLimits()
	//mgr.StoreSymbolsLimits(limits)
	//fmt.Println(mgr.GetAssetBalance("USDT"))
	//or, err := mgr.SetOrder(order.Parameters{
	//	Assets: symbol.Assets{
	//		Base:  "BUSD",
	//		Quote: "USDT",
	//	},
	//	Side:     order.Sell,
	//	Type:     order.Limit,
	//	Quantity: 15,
	//	Price:    1.5,
	//})
	//if err != nil {
	//	log.Error.Print(err)
	//}
	//log.Info.Print(or)
	////coins, _ := mgr.GetAllCoinsInfo()
	////fmt.Println(len(coins))
	////fmt.Println(coins)
}

type eurAssetStruct struct {
	symbol symbol.Assets
	price  float64
}

type usdtAssetStruct struct {
	symbol symbol.Assets
	price  float64
}
