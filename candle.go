package trader

type Candle struct {
	open         float64
	close        float64
	high         float64
	low          float64
	volume       float64
	timeframe    string
	datetime     int64
	openDatetime int64
	instrumentID int
}

func NewCandle(o, h, l, c, v float64, tf string, closeTS, openTS int64, instID int) Candle {
	return Candle{
		open:         o,
		close:        c,
		high:         h,
		low:          l,
		volume:       v,
		timeframe:    tf,
		datetime:     closeTS,
		openDatetime: openTS,
		instrumentID: instID,
	}
}

func (t Candle) Datetime() int64 {
	return t.datetime
}
func (t Candle) InstrumentID() int {
	return t.instrumentID
}
func (t Candle) RelevantPrice() float64 {
	return t.close
}

func (t Candle) Open() float64 {
	return t.open
}
func (t Candle) High() float64 {
	return t.high
}
func (t Candle) Low() float64 {
	return t.low
}
func (t Candle) Close() float64 {
	return t.close
}

func (t Candle) Volume() float64 {
	return t.volume
}
func (t Candle) Timeframe() string {
	return t.timeframe
}

func (t *Candle) OpenDatetime() int64 {
	return t.openDatetime

}

type CandleOpen struct {
	open         float64
	datetime     int64
	timeframe    string
	instrumentID int
}

func NewCandleOpen(o float64, datetime int64, tf string, instID int) CandleOpen {
	return CandleOpen{
		open:         o,
		datetime:     datetime,
		timeframe:    tf,
		instrumentID: instID,
	}
}

func (t CandleOpen) Datetime() int64 {
	return t.datetime
}
func (t CandleOpen) InstrumentID() int {
	return t.instrumentID
}
func (t CandleOpen) RelevantPrice() float64 {
	return t.open
}

func (t CandleOpen) Open() float64 {
	return t.open
}
func (t CandleOpen) Timeframe() string {
	return t.timeframe
}
