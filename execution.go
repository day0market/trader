package trader

type Execution struct {
	EventTime int64
	Price     float64
	Qty       float64
	BaseOrder *Order
}

func NewExecution(t int64, price, qty float64, order *Order) *Execution {
	return &Execution{
		EventTime: t,
		Price:     price,
		Qty:       qty,
		BaseOrder: order,
	}
}

func (e *Execution) Time() int64 {
	return e.EventTime
}

func (e *Execution) InstrumentID() int {
	return e.BaseOrder.instrumentID
}

func (e *Execution) Order() *Order {
	return e.BaseOrder
}
