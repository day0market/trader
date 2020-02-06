package trader

type ITick interface {
	IInstrumentPriceContainer

	IsValid() bool

	GetLastPrice() float64
	GetLastSize() int64
	GetBidPrice() float64
	GetAskPrice() float64
	HasTrade() bool
	HasQuote() bool

	GetAskSize() int64
	GetBidSize() int64

	IsClosing() bool
	IsOpening() bool
}
