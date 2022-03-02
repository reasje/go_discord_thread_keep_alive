package protocols


type Bot struct {
	Name string
	Token string
}

type BotServices interface {
	SendMessage()
	StartBot()
}