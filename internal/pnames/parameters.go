package pnames

// binance cmn parameters values
const (
	Filters    = "filters"
	Signature  = "signature"
	Timestamp  = "timestamp"
	ServerTime = "serverTime"
)

// types
const (
	Type       = "type"
	FilterType = "filterType"
)

// symbols info
const (
	Symbol  = "symbol"
	Symbols = "symbols"

	BaseAsset  = "baseAsset"
	QuoteAsset = "quoteAsset"
)

// price info
const (
	Price = "price"

	MinPrice = "minPrice"
	MaxPrice = "maxPrice"
	TickSize = "tickSize"
)

// quantity info
const (
	Quantity = "quantity"

	MinQuantity = "minQty"
	MaxQuantity = "maxQty"
	StepSize    = "stepSize"
)

// order info
const (
	OrderId = "orderId"
	Side    = "side"
	Status  = "status"

	QuoteOrderQty = "quoteOrderQty"
	TimeInForce   = "timeInForce"
)
