package trader

import "time"

type IBroker interface {
	NewOrder(o Order) error
	CancelOrder(o Order) error
	OnMarketEvent(event IInstrumentPriceContainer)
	OnTimeEvent(t time.Time)
	GetOrderById(id string) Order
	OpenPosition(i Instrument) IPosition
	AllOpenPositions() []IPosition
	AllClosedPositions() []IPosition
	ClosedPositions(i Instrument) []IPosition
	AllOpenOrders() []Order
	AllOrders() []Order
	OpenOrders(i Instrument) []Order
}
