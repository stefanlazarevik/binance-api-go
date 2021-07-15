package binance_api_go

type BinanceFunc interface {
	NewFullMarketOrder(orderParams OrdersParams) (float64, error)
	GetPrice(symbol string) (float64, error)
	AccountInfo(symbol string) (float64, error)
	Ping() error
	GetOpenOrders(symbol string) (bool, error)
	NewLimitOrder(orderParams OrdersParams) (bool, error)
}
