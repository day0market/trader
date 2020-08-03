package trader

type IFeed interface {
	Subscribe(chan<- IFeedEvent, []*Instrument)
	Unsubscribe()
}

type IFeedEvent interface {
	Datetime() int64
}

type ISessionEvent interface {
	IFeedEvent
	IInstrumentPriceContainer
	IsAuctionBased() bool
	Volume() float64
}

// *** Session Closed **
type SessionClosedEvent struct {
	EventTime     int64
	SessCloseTime int64
	SessVolume    float64
	Price         float64
	InstID        int
	Auction       bool
}

func NewSessionClosedEvent(et, sessCloseTime int64, price, vol float64, auctionBased bool, instID int) *SessionClosedEvent {
	return &SessionClosedEvent{
		EventTime:     et,
		SessCloseTime: sessCloseTime,
		SessVolume:    vol,
		Price:         price,
		Auction:       auctionBased,
		InstID:        instID,
	}
}

func (s SessionClosedEvent) RelevantPrice() float64 {
	return s.Price
}

func (s SessionClosedEvent) IsAuctionBased() bool {
	return s.Auction
}

func (s SessionClosedEvent) InstrumentID() int {
	return s.InstID
}

func (s SessionClosedEvent) Volume() float64 {
	return s.SessVolume
}

func (s SessionClosedEvent) Datetime() int64 {
	return s.EventTime
}

// *** After Session Close **
type AfterSessionCloseEvent struct {
	EventTime     int64
	SessCloseTime int64
	InstID        int
}

func NewAfterSessionClosedEvent(et, sct int64, instID int) *AfterSessionCloseEvent {
	return &AfterSessionCloseEvent{
		EventTime:     et,
		SessCloseTime: sct,
		InstID:        instID,
	}
}

func (s *AfterSessionCloseEvent) Datetime() int64 {
	return s.EventTime
}

// *** Session Will Close ***
type SessionWillCloseEvent struct {
	EventTime                int64
	ExpectedSessionCloseTime int64
	InstID                   int
}

func NewSessionWillCloseEvent(et, esct int64, instID int) *SessionWillCloseEvent {
	return &SessionWillCloseEvent{
		EventTime:                et,
		ExpectedSessionCloseTime: esct,
		InstID:                   instID,
	}
}

func (s *SessionWillCloseEvent) Datetime() int64 {
	return s.EventTime
}

// *** Session Open ***
type SessionOpenEvent struct {
	EventTime       int64
	SessionOpenTime int64
	Price           float64
	SessVolume      float64
	InstID          int
	Auction         bool
}

func NewSessionOpenEvent(et, sot int64, price, vol float64, instID int, auction bool) *SessionOpenEvent {
	return &SessionOpenEvent{
		EventTime:       et,
		SessionOpenTime: sot,
		Price:           price,
		SessVolume:      vol,
		InstID:          instID,
		Auction:         auction,
	}
}

func (s *SessionOpenEvent) Datetime() int64 {
	return s.EventTime
}

func (s *SessionOpenEvent) RelevantPrice() float64 {
	return s.Price
}

func (s *SessionOpenEvent) IsAuctionBased() bool {
	return s.Auction
}

func (s *SessionOpenEvent) InstrumentID() int {
	return s.InstID
}

func (s *SessionOpenEvent) Volume() float64 {
	return s.SessVolume
}

// *** Session Will Open ***
type SessionWillOpenEvent struct {
	EventTime       int64
	SessionOpenTime int64
	InstID          int
}

func NewSessionWillOpenEvent(et, sot int64, instID int) *SessionWillOpenEvent {
	return &SessionWillOpenEvent{
		EventTime:       et,
		SessionOpenTime: sot,
		InstID:          instID,
	}
}

func (s *SessionWillOpenEvent) Datetime() int64 {
	return s.EventTime
}

//*** All Candles Closed ****
type AllCandlesClosedEvent struct {
	EventTime         int64
	ExpectedCloseTime int64
}

func NewAllCandlesClosedEvent(et, ect int64) *AllCandlesClosedEvent {
	return &AllCandlesClosedEvent{
		EventTime:         et,
		ExpectedCloseTime: ect,
	}
}

func (s *AllCandlesClosedEvent) Datetime() int64 {
	return s.EventTime
}

// *** No Market Data ***
type NoMarketDataEvent struct {
	EventTime int64
	Reason    string
}

func NewNoMarketDataEvent(et int64, reason string) *NoMarketDataEvent {
	return &NoMarketDataEvent{
		EventTime: et,
		Reason:    reason,
	}
}

func (s *NoMarketDataEvent) Datetime() int64 {
	return s.EventTime
}
