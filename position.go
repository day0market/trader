package trader

import (
	"math"
)

type PositionSide string

const (
	LongPosition    PositionSide = "Long"
	ShortPosition   PositionSide = "Short"
	NotOpenPosition PositionSide = "NotOpen"
)

type Position struct {
	openPrice    float64
	closePrice   float64
	closedPnL    float64
	openPnL      float64
	openSize     float64
	tradedSize   float64
	closeTime    int64
	openTime     int64
	instrumentID int
	executions   []*Execution
	side         PositionSide
	id           string
}

func (p *Position) IsClosed() bool {
	if p.tradedSize != 0 && p.openSize == 0 {
		return true
	}
	return false
}

func (p *Position) Instrument() int {
	return p.instrumentID
}

func (p *Position) ID() string {
	return p.id
}

func (p *Position) OpenTime() int64 {
	return p.openTime
}

func (p *Position) CloseTime() int64 {
	return p.closeTime
}

func (p *Position) Side() PositionSide {
	return p.side
}

func (p *Position) OpenSize() float64 {
	return p.openSize
}

func (p *Position) ClosedPnL() float64 {
	return p.closedPnL
}

func (p *Position) OpenPnL() float64 {
	return p.openPnL
}

func (p *Position) OpenPrice() float64 {
	return p.openPrice
}

func (p *Position) FirstPrice() float64 {
	if len(p.executions) == 0 {
		return math.NaN()
	}
	return p.executions[0].Price()
}

func (p *Position) LastPrice() float64 {
	if len(p.executions) == 0 {
		return math.NaN()
	}
	return p.executions[len(p.executions)-1].Price()
}

func (p *Position) IsOpen() bool {
	return p.openSize != 0
}

func (p *Position) Executions() []*Execution {
	return p.executions
}

func (p *Position) AddExecution(e *Execution) {
	p.tradedSize += math.Abs(e.Qty())
	p.executions = append(p.executions, e)
	e.order.SetPositionID(p.id)
}
