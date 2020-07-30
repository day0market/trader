package trader

type OrderCancel struct {
	time         int64
	instrumentID int
	order        *Order
}

func NewOrderCancel(time int64, instID int, order *Order) *OrderCancel {
	return &OrderCancel{
		time:         time,
		instrumentID: instID,
		order:        order,
	}
}

func (c *OrderCancel) Time() int64 {
	return c.time
}

func (c *OrderCancel) InstrumentID() int {
	return c.instrumentID
}

func (c *OrderCancel) Order() *Order {
	return c.order
}

type IOrderEvent interface {
	Time() int64
	InstrumentID() int
	Order() *Order
}
