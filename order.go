package trader

import (
	"math"
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

type Order struct {
	orderType      OrderType
	orderState     OrderState
	orderSide      OrderSide
	executionPrice float64
	price          float64
	qty            float64
	time           int64
	creationTime   int64
	instrumentID   int
	tif            OrderTIF
	id             string
	posID          string
	continuousTags map[string]float64
	discreteTags   map[string]string
}

func (t *Order) ExecutionPrice() float64 {
	panic("Not tested")
	return t.executionPrice
}

func (t *Order) SetUpdateTime(time int64) {
	t.time = time
}

func (t *Order) UpdateTime() int64 {
	return t.time
}

func (t *Order) SetCreationTime(time int64) {
	t.creationTime = time
}

func (t *Order) CreationTime() int64 {
	return t.creationTime
}

func (t *Order) SetPositionID(posID string) {
	t.posID = posID
}

func (t *Order) PositionID() string {
	return t.posID
}

func (t *Order) ContinuousTags() map[string]float64 {
	return t.continuousTags
}

func (t *Order) DiscreteTags() map[string]string {
	return t.discreteTags
}

func (t *Order) Side() OrderSide {
	return t.orderSide
}
func (t *Order) Price() float64 {
	return t.price
}
func (t *Order) Type() OrderType {
	return t.orderType
}
func (t *Order) State() OrderState {
	return t.orderState
}
func (t *Order) SetState(s OrderState) bool {
	t.orderState = s
	return true
}
func (t *Order) Qty() float64 {
	return t.qty
}
func (t *Order) UpdateQty(q float64) bool {
	t.qty = q
	return true
}

func (t *Order) InstrumentID() int {
	return t.instrumentID
}
func (t *Order) TIF() OrderTIF {
	return t.tif
}
func (t *Order) Replace(price float64) bool {
	t.price = price
	return true
}
func (t *Order) ID() string {
	return t.id
}

func (t *Order) SetID(id string) {
	t.id = id
}

type OrderParams struct {
	OrderSide      OrderSide
	OrderType      OrderType
	Price          float64
	Qty            float64
	InstrumentID   int
	TIF            OrderTIF
	ID             string
	ContinuousTags map[string]float64
	DiscreteTags   map[string]string
}

func NewOrderFromParams(p *OrderParams) Order {
	return Order{
		orderType:      p.OrderType,
		orderState:     NewOrder,
		orderSide:      p.OrderSide,
		executionPrice: math.NaN(),
		price:          p.Price,
		qty:            p.Qty,
		time:           0,
		creationTime:   0,
		instrumentID:   p.InstrumentID,
		tif:            p.TIF,
		id:             p.ID,
		posID:          "",
		continuousTags: p.ContinuousTags,
		discreteTags:   p.DiscreteTags,
	}
}
