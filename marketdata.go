package trader

import "time"

type IInstrument interface {
	GetSymbol() string
	GetExchange() IExchange
}

type IExchange interface {
	GetMarketOpenTime() ITimeOfDay
	GetMarketCloseTime() ITimeOfDay
}

type ITimeOfDay interface {
	Hour() int
	Minute() int
	Second() int
	Location() *time.Location
}

type IInstrumentPriceContainer interface {
	GetDatetime() time.Time
	GetInstrument() IInstrument
	GetRelevantPrice() float64
}

type Instrument struct {
	Symbol string
	*Exchange
}

func (i Instrument) GetSymbol() string {
	return i.Symbol
}

func (i Instrument) GetExchange() IExchange {
	return i.Exchange
}

type Exchange struct {
	OpenTime  time.Time
	CloseTime time.Time
	Name      string
}

func (e *Exchange) GetMarketOpenTime() ITimeOfDay {
	return e.OpenTime
}

func (e *Exchange) GetMarketCloseTime() ITimeOfDay {
	return e.CloseTime
}
