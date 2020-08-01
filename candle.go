package trader

type Candle struct {
	Open          float64
	Close         float64
	High          float64
	Low           float64
	Volume        float64
	Timeframe     string
	CloseDatetime int64
	OpenDatetime  int64
	InstID        int
}

func NewCandle(o, h, l, c, v float64, tf string, closeTS, openTS int64, instID int) Candle {
	return Candle{
		Open:          o,
		Close:         c,
		High:          h,
		Low:           l,
		Volume:        v,
		Timeframe:     tf,
		CloseDatetime: closeTS,
		OpenDatetime:  openTS,
		InstID:        instID,
	}
}

func (t Candle) Datetime() int64 {
	return t.CloseDatetime
}
func (t Candle) InstrumentID() int {
	return t.InstID
}
func (t Candle) RelevantPrice() float64 {
	return t.Close
}

type CandleOpen struct {
	Open         float64
	OpenDatetime int64
	Timeframe    string
	InstID       int
}

func NewCandleOpen(o float64, datetime int64, tf string, instID int) *CandleOpen {
	return &CandleOpen{
		Open:         o,
		OpenDatetime: datetime,
		Timeframe:    tf,
		InstID:       instID,
	}
}

func (t CandleOpen) Datetime() int64 {
	return t.OpenDatetime
}
func (t CandleOpen) InstrumentID() int {
	return t.InstID
}
func (t CandleOpen) RelevantPrice() float64 {
	return t.Open
}
