package trader

type IFeed interface {
	Subscribe(chan<- IFeedEvent, []Instrument)
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
	eventTime        int64
	sessionCloseTime int64
	volume           float64
	price            float64
	instrumentID     int
	auctionBased     bool
}

func NewSessionClosedEvent(et, sessCloseTime int64, price, vol float64, auctionBased bool, instID int) SessionClosedEvent {
	return SessionClosedEvent{
		eventTime:        et,
		sessionCloseTime: sessCloseTime,
		volume:           vol,
		price:            price,
		auctionBased:     auctionBased,
		instrumentID:     instID,
	}
}

func (s SessionClosedEvent) RelevantPrice() float64 {
	return s.price
}

func (s SessionClosedEvent) IsAuctionBased() bool {
	return s.auctionBased
}

func (s SessionClosedEvent) InstrumentID() int {
	return s.instrumentID
}

func (s SessionClosedEvent) Volume() float64 {
	return s.volume
}

func (s SessionClosedEvent) Datetime() int64 {
	return s.eventTime
}

func (s *SessionClosedEvent) SessionCloseTime() int64 {
	return s.sessionCloseTime
}

// *** After Session Close **
type AfterSessionCloseEvent struct {
	eventTime        int64
	sessionCloseTime int64
	instrumentID     int
}

func NewAfterSessionClosedEvent(et, sct int64, instID int) AfterSessionCloseEvent {
	return AfterSessionCloseEvent{
		eventTime:        et,
		sessionCloseTime: sct,
		instrumentID:     instID,
	}
}

func (s AfterSessionCloseEvent) Datetime() int64 {
	return s.eventTime
}

func (s AfterSessionCloseEvent) SessionCloseTime() int64 {
	return s.sessionCloseTime
}

func (s AfterSessionCloseEvent) InstrumentID() int {
	return s.instrumentID
}

// *** Session Will Close ***
type SessionWillCloseEvent struct {
	eventTime                int64
	expectedSessionCloseTime int64
	instrumentID             int
}

func NewSessionWillCloseEvent(et, esct int64, instID int) SessionWillCloseEvent {
	return SessionWillCloseEvent{
		eventTime:                et,
		expectedSessionCloseTime: esct,
		instrumentID:             instID,
	}
}

func (s *SessionWillCloseEvent) ExpectedSessionCloseTime() int64 {
	return s.expectedSessionCloseTime
}

func (s *SessionWillCloseEvent) Datetime() int64 {
	return s.eventTime
}

func (s *SessionWillCloseEvent) InstrumentID() int {
	return s.instrumentID
}

// *** Session Open ***
type SessionOpenEvent struct {
	eventTime       int64
	sessionOpenTime int64
	price           float64
	volume          float64
	instrumentID    int
	auctionBased    bool
}

func NewSessionOpenEvent(et, sot int64, price, vol float64, instID int, auction bool) SessionOpenEvent {
	return SessionOpenEvent{
		eventTime:       et,
		sessionOpenTime: sot,
		price:           price,
		volume:          vol,
		instrumentID:    instID,
		auctionBased:    auction,
	}
}

func (s *SessionOpenEvent) SessionOpenTime() int64 {
	return s.sessionOpenTime
}

func (s *SessionOpenEvent) Datetime() int64 {
	return s.eventTime
}

func (s *SessionOpenEvent) RelevantPrice() float64 {
	return s.price
}

func (s *SessionOpenEvent) IsAuctionBased() bool {
	return s.auctionBased
}

func (s *SessionOpenEvent) InstrumentID() int {
	return s.instrumentID
}

func (s *SessionOpenEvent) Volume() float64 {
	return s.volume
}

// *** Session Will Open ***
type SessionWillOpenEvent struct {
	eventTime       int64
	sessionOpenTime int64
	instrumentID    int
}

func NewSessionWillOpenEvent(et, sot int64, instID int) SessionWillOpenEvent {
	return SessionWillOpenEvent{
		eventTime:       et,
		sessionOpenTime: sot,
		instrumentID:    instID,
	}
}

func (s *SessionWillOpenEvent) InstrumentID() int {
	return s.instrumentID
}

func (s *SessionWillOpenEvent) Datetime() int64 {
	return s.eventTime
}

func (s *SessionWillOpenEvent) SessionExpectedOpenTime() int64 {
	return s.sessionOpenTime
}

//*** All Candles Closed ****
type AllCandlesClosedEvent struct {
	eventTime         int64
	expectedCloseTime int64
}

func NewAllCandlesClosedEvent(et, ect int64) AllCandlesClosedEvent {
	return AllCandlesClosedEvent{
		eventTime:         et,
		expectedCloseTime: ect,
	}
}

func (s *AllCandlesClosedEvent) Datetime() int64 {
	return s.eventTime
}

func (s *AllCandlesClosedEvent) ExpectedCloseTime() int64 {
	return s.expectedCloseTime
}

// *** No Market Data ***
type NoMarketDataEvent struct {
	eventTime   int64
	eventReason string
}

func NewNoMarketDataEvent(et int64, reason string) NoMarketDataEvent {
	return NoMarketDataEvent{
		eventTime:   et,
		eventReason: reason,
	}
}

func (s *NoMarketDataEvent) Datetime() int64 {
	return s.eventTime
}

func (s *NoMarketDataEvent) Reason() string {
	return s.eventReason
}
