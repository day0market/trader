package trader

import (
	"fmt"
	"time"
)

type IInstrument interface {
	GetSymbol() string
	GetExchange() IExchange
	GetIdentifier() string
	GetMetadata() map[string]string
}

type IExchange interface {
	GetMarketOpenTime() ITimeOfDay
	GetMarketCloseTime() ITimeOfDay
	GetLocation() *time.Location
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

func (i *Instrument) GetIdentifier() string {
	return fmt.Sprintf("%s.%s", i.Symbol, i.Exchange.Name)
}

func (i *Instrument) GetMetadata() map[string]string {
	return nil
}

type Exchange struct {
	OpenTime  time.Time
	CloseTime time.Time
	Name      string
	Location  *time.Location
}

func (e *Exchange) GetMarketOpenTime() ITimeOfDay {
	return e.OpenTime
}

func (e *Exchange) GetMarketCloseTime() ITimeOfDay {
	return e.CloseTime
}

func (e *Exchange) GetLocation() *time.Location {
	return e.Location
}
