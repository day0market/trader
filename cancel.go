package trader

type OrderCancel struct {
	EventTime int64
	InstID    int
	BaseOrder *Order
}

func NewOrderCancel(time int64, instID int, order *Order) *OrderCancel {
	return &OrderCancel{
		EventTime: time,
		InstID:    instID,
		BaseOrder: order,
	}
}

func (c *OrderCancel) Time() int64 {
	return c.EventTime
}

func (c *OrderCancel) InstrumentID() int {
	return c.InstID
}

func (c *OrderCancel) Order() *Order {
	return c.BaseOrder
}

type IOrderEvent interface {
	Time() int64
	InstrumentID() int
	Order() *Order
}
