package trader

import "time"

type IFeed interface {
	Subscribe(chan<- IFeedEvent, []IInstrument)
	Unsubscribe()
}

type IFeedEvent interface {
	GetDatetime() time.Time
}

type IAllCandlesClosedEvent interface {
	IFeedEvent
	GetExpectedCloseTime() time.Time
}

type ISessionClosedEvent interface {
	IFeedEvent
	GetSessionCloseTime() time.Time
}

type ISessionWillOpenEvent interface {
	IFeedEvent
	GetSessionExpectedOpenTime() time.Time
}

type INoMarketDataEvent interface {
	IFeedEvent
	Reason() string
}

type SessionClosedEvent struct {
	EventTime        time.Time
	SessionCloseTime time.Time
}

func (s *SessionClosedEvent) GetDatetime() time.Time {
	return s.EventTime
}

func (s *SessionClosedEvent) GetSessionCloseTime() time.Time {
	return s.SessionCloseTime
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
