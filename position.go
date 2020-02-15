package trader

import "time"

type PositionSide string

const (
	LongPosition    PositionSide = "Long"
	ShortPosition   PositionSide = "Short"
	NotOpenPosition PositionSide = "NotOpen"
)

type IPosition interface {
	Instrument() IInstrument
	OpenTime() time.Time
	CloseTime() time.Time
	Side() PositionSide
	GetOpenSize() int64
	GetClosedPnL() float64
	GetOpenPnL() float64
	GetOpenPrice() float64
}
