package main

func main() {
	//log.Init("binance-api-go", true)
	//mgr := binance.New(exchangeapi.ApiKey{
	//	Key:    os.Args[1],
	//	Secret: os.Args[2],
	//})

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
