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
	MetaData map[string]string
	*Exchange
}

func (i Instrument) SymbolAndExchangeID() string {
	return fmt.Sprintf("%s.%s", i.Symbol, i.Exchange.Name)
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

type TimeOfDay struct {
	Hour     int
	Minute   int
	Second   int
	Location *time.Location
}
