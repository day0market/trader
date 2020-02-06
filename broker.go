package trader

type IBroker interface {
	NewOrder(o IOrder) error
	CancelOrder(o IOrder) error
	OnMarketEvent(event IInstrumentPriceContainer)
	GetOrderById(id string) IOrder
	OpenPosition(i IInstrument) IPosition
	AllOpenPositions() []IPosition
	AllClosedPositions() []IPosition
	ClosedPositions(i IInstrument) []IPosition
	AllOpenOrders() []IOrder
	OpenOrders(i IInstrument) []IOrder
}
