package trader

type Execution struct {
	time  int64
	price float64
	qty   float64
	order *Order
}

func NewExecution(t int64, price, qty float64, order *Order) *Execution {
	return &Execution{
		time:  t,
		price: price,
		qty:   qty,
		order: order,
	}
}

func (e *Execution) Time() int64 {
	return e.time
}

func (e *Execution) InstrumentID() int {
	return e.order.instrumentID
}

func (e *Execution) Order() *Order {
	return e.order
}

func (e *Execution) Qty() float64 {
	return e.qty
}

func (e *Execution) Price() float64 {
	return e.price
}
