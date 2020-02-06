package trader

import "time"

type ICandle interface {
	IInstrumentPriceContainer

	GetOpen() float64
	GetHigh() float64
	GetLow() float64
	GetClose() float64
	GetTimeframe() string
	GetCloseDatetime() time.Time
}

type ICandleOpen interface {
	IInstrumentPriceContainer
	GetOpen() float64
	GetTimeframe() string
}

type Candle struct {
	Open       float64
	Close      float64
	High       float64
	Low        float64
	Volume     int64
	Timeframe  string
	Datetime   time.Time
	Instrument IInstrument
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
func (t *Candle) GetTimeframe() string {
	return t.Timeframe
}

func (t *Candle) GetCloseDatetime() time.Time {
	delta := 1
	if t.Timeframe == "D" {
		delta = 8 * 60
	}
	if t.Timeframe == "W" {
		delta = 24 * 6
	}

	return t.Datetime.Add(time.Minute*time.Duration(delta) - time.Second)

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
