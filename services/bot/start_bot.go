package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/reasje/go_discord_thread_saver/internal/storage"
)

func StartBot() {

	//creating new bot session
	var err error
	storage.GoBot, err = discordgo.New("Bot " + storage.Config.Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// // opens a connection to discord
	// err = storage.GoBot.Open()
	// //Error handling
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}
