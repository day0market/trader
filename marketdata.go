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
	GetName() string
	HasOpeningAuction() bool
	HasClosingAuction() bool
	OpenSessionDuration() int
	CloseSessionDuration() int
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
	OpenTime       time.Time
	CloseTime      time.Time
	Name           string
	Location       *time.Location
	OpeningAuction bool
	ClosingAuction bool
	OpenDuration   int
	CloseDuration  int
}

func (e *Exchange) OpenSessionDuration() int {
	return e.OpenDuration
}

func (e *Exchange) CloseSessionDuration() int {
	return e.CloseDuration
}

func (e *Exchange) GetMarketOpenTime() ITimeOfDay {
	return e.OpenTime
}

func (e *Exchange) HasOpeningAuction() bool {
	return e.OpeningAuction
}

func (e *Exchange) HasClosingAuction() bool {
	return e.ClosingAuction
}

func (e *Exchange) GetMarketCloseTime() ITimeOfDay {
	return e.CloseTime
}

func (e *Exchange) GetLocation() *time.Location {
	return e.Location
}

func (e *Exchange) GetName() string {
	return e.Name
}
