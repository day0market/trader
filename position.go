package trader

type PositionSide string

const (
	LongPosition    PositionSide = "Long"
	ShortPosition   PositionSide = "Short"
	NotOpenPosition PositionSide = "NotOpen"
)

type IPosition interface {
	InstrumentID() int
	OpenTime() int64
	CloseTime() int64
	Side() PositionSide
	OpenSize() float64
	ClosedPnL() float64
	OpenPnL() float64
	OpenPrice() float64
	FirstPrice() float64
	LastPrice() float64
	Executions() []*Execution
	ID() string
	IsClosed() bool
}
