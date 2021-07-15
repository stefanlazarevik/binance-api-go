package binance_api_go

import "net/http"

const burl = "https://api.binance.com"

type Account struct {
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	Balances         []Balance `json:"balances"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type OrdersParams struct {
	Symbol      string
	Side        string
	OrderType   string
	AssetCount  float64
	TimeInForce string
	Price       float64
}
type OrderResult struct {
	Symbol              string `json:"symbol"`
	OrderId             int    `json:"orderId"`
	OrderListId         int    `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int    `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

type PriceFilter struct {
	MinPrice float64
	MaxPrice float64
	TickSize float64
}

type LotSize struct {
	MinQuantity float64
	MaxQuantity float64
	StepSize    float64
}

type ExchangeInfo struct {
	PriceFilter PriceFilter
	LotSize     LotSize
}

type BodyAnswer struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	priceFilterType = "PRICE_FILTER"
	lotSizeType     = "LOT_SIZE"
)

const filled = "FILLED"

// trading side
const (
	Buy  = "BUY"
	Sell = "SELL"
)
const (
	Gtc = "GTC"
)

// trading order type
const (
	Limit  = "LIMIT"
	Market = "MARKET"
)

// Request methods
const (
	Get  = "GET"
	Post = "POST"
)
