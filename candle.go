package trader

import "time"

type ICandle interface {
	IInstrumentPriceContainer

	GetOpen() float64
	GetHigh() float64
	GetLow() float64
	GetClose() float64
	GetVolume() float64
	GetTimeframe() string
	GetOpenDatetime() time.Time
}

type ICandleOpen interface {
	IInstrumentPriceContainer
	GetOpen() float64
	GetTimeframe() string
}

type Candle struct {
	Open         float64
	Close        float64
	High         float64
	Low          float64
	Volume       int64
	Timeframe    string
	Datetime     time.Time
	OpenDatetime time.Time
	Instrument   IInstrument
}

func (t *Candle) GetDatetime() time.Time {
	return t.Datetime
}
func (t *Candle) GetInstrument() IInstrument {
	return t.Instrument
}
func (t *Candle) GetRelevantPrice() float64 {
	return t.Close

}
func (t *Candle) GetOpen() float64 {
	return t.Open
}
func (t *Candle) GetHigh() float64 {
	return t.High
}
func (t *Candle) GetLow() float64 {
	return t.Low
}
func (t *Candle) GetClose() float64 {
	return t.Close
}

func (t *Candle) GetVolume() float64 {
	return float64(t.Volume)
}
func (t *Candle) GetTimeframe() string {
	return t.Timeframe
}

func (t *Candle) GetOpenDatetime() time.Time {
	return t.OpenDatetime

}

type CandleOpen struct {
	Open       float64
	Timeframe  string
	Datetime   time.Time
	Instrument IInstrument
}

func (t *CandleOpen) GetDatetime() time.Time {
	return t.Datetime
}
func (t *CandleOpen) GetInstrument() IInstrument {
	return t.Instrument
}
func (t *CandleOpen) GetRelevantPrice() float64 {
	return t.Open

}
func (t *CandleOpen) GetOpen() float64 {
	return t.Open
}
func (t *CandleOpen) GetTimeframe() string {
	return t.Timeframe
}
