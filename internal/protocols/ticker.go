package protocols

type Ticker struct {
	ID int64 `json:"id"`
	ticker *Ticker
}

type TickerServices interface {
	KeepAlive()
	StartTicker()
	StopTicker()
}