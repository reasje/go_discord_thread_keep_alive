package bot

import (
	"fmt"

	"github.com/reasje/go_discord_thread_saver/internal/storage"
)



func SendMessage() {
	for _ , thread := range storage.Channels {
		// goBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { fmt.Println("Bot is ready") })
		// fmt.Println(thread.ID)
		// err := storage.GoBot.Open()

		// if err != nil {
		// 	Println("Error opening websocket : ")
		// }
		messsage , err := storage.GoBot.ChannelMessageSend(thread.ID, "Whatever1 , should be deleted ")

		if err != nil {
			fmt.Println("Error sending message")
			return
		}

		err = storage.GoBot.ChannelMessageDelete(thread.ID, messsage.ID)

		if err != nil {
			fmt.Println("Error deleting message")
			return
		}

	}
}