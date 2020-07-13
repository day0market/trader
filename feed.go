package trader

import "time"

type IFeed interface {
	Subscribe(chan<- IFeedEvent, []IInstrument)
	Unsubscribe()
}

type IFeedEvent interface {
	GetDatetime() time.Time
}

type ISessionEvent interface {
	IFeedEvent
	IInstrumentPriceContainer
	IsAuctionBased() bool
	GetVolume() int64
}

// *** Session Closed **
type ISessionClosedEvent interface {
	IFeedEvent
	ISessionEvent
	GetSessionCloseTime() time.Time
}

type SessionClosedEvent struct {
	EventTime        time.Time
	SessionCloseTime time.Time
	Price            float64
	Volume           int64
	AuctionBased     bool
	Instrument       IInstrument
}

func (s *SessionClosedEvent) GetRelevantPrice() float64 {
	return s.Price
}

func (s *SessionClosedEvent) IsAuctionBased() bool {
	return s.AuctionBased
}

func (s *SessionClosedEvent) GetInstrument() IInstrument {
	return s.Instrument
}

func (s *SessionClosedEvent) GetVolume() int64 {
	return s.Volume
}

func (s *SessionClosedEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *SessionClosedEvent) GetSessionCloseTime() time.Time {
	return s.SessionCloseTime
}

// *** After Session Close **
type IAfterSessionCloseEvent interface {
	IFeedEvent
	GetSessionCloseTime() time.Time
}

type AfterSessionCloseEvent struct {
	EventTime        time.Time
	SessionCloseTime time.Time
}

func (s *AfterSessionCloseEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *AfterSessionCloseEvent) GetSessionCloseTime() time.Time {
	return s.SessionCloseTime
}

// *** Session Will Close ***
type ISessionWillCloseEvent interface {
	IFeedEvent
	GetExpectedSessionCloseTime() time.Time
}
type SessionWillCloseEvent struct {
	EventTime                time.Time
	ExpectedSessionCloseTime time.Time
}

func (s *SessionWillCloseEvent) GetExpectedSessionCloseTime() time.Time {
	return s.ExpectedSessionCloseTime
}

func (s *SessionWillCloseEvent) GetDatetime() time.Time {
	return s.EventTime
}

// *** Session Open ***
type ISessionOpenEvent interface {
	IFeedEvent
	ISessionEvent
	GetSessionOpenTime() time.Time
}

type SessionOpenEvent struct {
	EventTime       time.Time
	SessionOpenTime time.Time
	Instrument      IInstrument
	Price           float64
	AuctionBased    bool
	Volume          int64
}

func (s *SessionOpenEvent) GetSessionOpenTime() time.Time {
	return s.SessionOpenTime
}

func (s *SessionOpenEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *SessionOpenEvent) GetRelevantPrice() float64 {
	return s.Price
}

func (s *SessionOpenEvent) IsAuctionBased() bool {
	return s.AuctionBased
}

func (s *SessionOpenEvent) GetInstrument() IInstrument {
	return s.Instrument
}

func (s *SessionOpenEvent) GetVolume() int64 {
	return s.Volume
}

// *** Session Will Open ***
type ISessionWillOpenEvent interface {
	IFeedEvent
	GetSessionExpectedOpenTime() time.Time
}

type SessionWillOpenEvent struct {
	EventTime       time.Time
	SessionOpenTime time.Time
}

func (s *SessionWillOpenEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *SessionWillOpenEvent) GetSessionExpectedOpenTime() time.Time {
	return s.SessionOpenTime
}

//*** All Candles Closed ****
type IAllCandlesClosedEvent interface {
	IFeedEvent
	GetExpectedCloseTime() time.Time
}

type AllCandlesClosedEvent struct {
	EventTime         time.Time
	ExpectedCloseTime time.Time
}

func (s *AllCandlesClosedEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *AllCandlesClosedEvent) GetExpectedCloseTime() time.Time {
	return s.ExpectedCloseTime
}

// *** No Market Data ***
type INoMarketDataEvent interface {
	IFeedEvent
	Reason() string
}

type NoMarketDataEvent struct {
	EventTime   time.Time
	EventReason string
}

func (s *NoMarketDataEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *NoMarketDataEvent) Reason() string {
	return s.EventReason
}
