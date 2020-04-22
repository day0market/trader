package trader

import "time"

type OrderType string

func (t OrderType) isAuction() bool {
	if t == LimitOnClose || t == LimitOnOpen || t == MarketOnClose || t == MarketOnOpen {
		return true
	}
	return false
}

type OrderSide string
type OrderState string

func (s OrderState) IsActive() bool {
	if s == PartialFilledOrder || s == ConfirmedOrder {
		return true
	}

	return false
}

type OrderTIF string

const (
	OrderBuy  OrderSide = "B"
	OrderSell OrderSide = "S"

	NewOrder           OrderState = "NewOrder"
	ConfirmedOrder     OrderState = "ConfirmedOrder"
	FilledOrder        OrderState = "FilledOrder"
	PartialFilledOrder OrderState = "PartialFilledOrder"
	CanceledOrder      OrderState = "CanceledOrder"
	RejectedOrder      OrderState = "RejectedOrder"

	LimitOrder    OrderType = "LMT"
	MarketOrder   OrderType = "MKT"
	StopOrder     OrderType = "STP"
	LimitOnClose  OrderType = "LOC"
	LimitOnOpen   OrderType = "LOO"
	MarketOnClose OrderType = "MOC"
	MarketOnOpen  OrderType = "MOO"

	DayTIF OrderTIF = "DayTIF"
	GTCTIF OrderTIF = "GTCTIF"
)

type IOrder interface {
	GetSide() OrderSide
	GetPrice() float64
	GetType() OrderType
	GetState() OrderState
	SetState(OrderState) bool
	GetQty() int64
	UpdateQty(int64) bool
	SetUpdateTime(t time.Time)
	GetUpdateTime() time.Time
	GetInstrument() IInstrument

	GetTIF() OrderTIF
	Replace(price float64) bool
	GetID() string
}
