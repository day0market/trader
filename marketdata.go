package trader

import (
	"fmt"
	"time"
)

type IInstrumentPriceContainer interface {
	Datetime() int64
	InstrumentID() int
	RelevantPrice() float64
}

type Instrument struct {
	Symbol   string
	SystemID int
	*Exchange
}

func (i Instrument) GetSymbol() string {
	return i.Symbol
}

func (i Instrument) GetExchange() *Exchange {
	return i.Exchange
}

func (i Instrument) GetIdentifier() string {
	return fmt.Sprintf("%s.%s", i.Symbol, i.Exchange.Name)
}

func (i *Instrument) GetMetadata() map[string]string {
	return nil
}

type Exchange struct {
	OpenTime       TimeOfDay
	CloseTime      TimeOfDay
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

func (e *Exchange) GetMarketOpenTime() TimeOfDay {
	return e.OpenTime
}

func (e *Exchange) HasOpeningAuction() bool {
	return e.OpeningAuction
}

func (e *Exchange) HasClosingAuction() bool {
	return e.ClosingAuction
}

func (e *Exchange) GetMarketCloseTime() TimeOfDay {
	return e.CloseTime
}

func (e *Exchange) GetLocation() *time.Location {
	return e.Location
}

func (e *Exchange) GetName() string {
	return e.Name
}

type TimeOfDay struct {
	Hour     int
	Minute   int
	Second   int
	Location *time.Location
}
