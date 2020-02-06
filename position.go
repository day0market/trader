package trader

type IPosition interface {
	GetOpenSize() int64
	GetClosedPnL() float64
	GetOpenPnL() float64
	GetOpenPrice() float64
}
