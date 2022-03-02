package ticker

import (
	"fmt"

	"github.com/reasje/go_discord_thread_saver/services/bot"
	"github.com/reasje/go_discord_thread_saver/services/channels"
)

func KeepAlive() {
	// reading threads from file
	err := channels.ReadChannels()

	if err != nil {
		fmt.Println("Error reading threads ")
		return
	}

	bot.SendMessage()
}