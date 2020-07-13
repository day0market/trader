package trader

import (
	"math"
	"time"
)

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
	GetExecutionPrice() float64
	GetType() OrderType
	GetState() OrderState
	SetState(OrderState) bool
	GetQty() int64
	UpdateQty(int64) bool
	SetUpdateTime(t time.Time)
	GetUpdateTime() time.Time
	SetCreationTime(t time.Time)
	GetCreationTime() time.Time
	Instrument() IInstrument

	GetTIF() OrderTIF
	Replace(price float64) bool
	GetID() string
	GetContinuousTags() map[string]float64
	GetDiscreteTags() map[string]string
	SetPositionID(string)
	GetPositionID() string
}

type Order struct {
	orderType      OrderType
	orderState     OrderState
	orderSide      OrderSide
	executionPrice float64
	price          float64
	qty            int64
	time           time.Time
	creationTime   time.Time
	instrument     IInstrument
	tif            OrderTIF
	id             string
	posID          string
	continuousTags map[string]float64
	discreteTags   map[string]string
}

func (t *Order) GetExecutionPrice() float64 {
	panic("Not tested")
	return t.executionPrice
}

func (t *Order) SetUpdateTime(time time.Time) {
	t.time = time
}

func (t *Order) GetUpdateTime() time.Time {
	return t.time
}

func (t *Order) SetCreationTime(time time.Time) {
	t.creationTime = time
}

func (t *Order) GetCreationTime() time.Time {
	return t.creationTime
}

func (t *Order) SetPositionID(posID string) {
	t.posID = posID
}

func (t *Order) GetPositionID() string {
	return t.posID
}

func (t *Order) GetContinuousTags() map[string]float64 {
	return t.continuousTags
}

func (t *Order) GetDiscreteTags() map[string]string {
	return t.discreteTags
}

func (t *Order) GetSide() OrderSide {
	return t.orderSide
}
func (t *Order) GetPrice() float64 {
	return t.price
}
func (t *Order) GetType() OrderType {
	return t.orderType
}
func (t *Order) GetState() OrderState {
	return t.orderState
}
func (t *Order) SetState(s OrderState) bool {
	t.orderState = s
	return true
}
func (t *Order) GetQty() int64 {
	return t.qty
}
func (t *Order) UpdateQty(q int64) bool {
	t.qty = q
	return true
}

func (t *Order) Instrument() IInstrument {
	return t.instrument
}
func (t *Order) GetTIF() OrderTIF {
	return t.tif
}
func (t *Order) Replace(price float64) bool {
	t.price = price
	return true
}
func (t *Order) GetID() string {
	return t.id
}

type OrderParams struct {
	OrderSide      OrderSide
	OrderType      OrderType
	Price          float64
	Qty            int64
	Instrument     IInstrument
	TIF            OrderTIF
	ID             string
	ContinuousTags map[string]float64
	DiscreteTags   map[string]string
}

func NewOrderFromParams(p *OrderParams) IOrder {
	return &Order{
		orderType:      p.OrderType,
		orderState:     NewOrder,
		orderSide:      p.OrderSide,
		executionPrice: math.NaN(),
		price:          p.Price,
		qty:            p.Qty,
		time:           time.Time{},
		creationTime:   time.Time{},
		instrument:     p.Instrument,
		tif:            p.TIF,
		id:             p.ID,
		posID:          "",
		continuousTags: p.ContinuousTags,
		discreteTags:   p.DiscreteTags,
	}
}
