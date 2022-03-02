package protocols

type Config struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}


type Configer interface {
	ReadConfig() error
}