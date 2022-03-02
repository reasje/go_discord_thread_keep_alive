package protocols

type Channel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ChannelServices interface {
	ReadThreads() error
}
