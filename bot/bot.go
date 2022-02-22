package bot

import (
	"fmt"
	"log"
	// "os"
	// "os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/reasje/go_discord_thread_saver/config"
	"github.com/reasje/go_discord_thread_saver/entity"
)

var BotId string
var goBot *discordgo.Session

func Start() {

    //creating new bot session
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)

//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}


	// opens a connection to discord
	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    //If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

func SendMessage() {
	for _ , thread := range entity.Threads {
		// goBot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { fmt.Println("Bot is ready") })
		fmt.Println(thread.ID)
		messsage , err := goBot.ChannelMessageSend(thread.ID, "Whatever , should be deleted ")

		if err != nil {
			fmt.Println("Error sending message")
			return
		}

		// goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
		// err = goBot.Open()
		if err != nil {
			log.Fatalf("Cannot open the session: %v", err)
		}
		// defer goBot.Close()
		// stop := make(chan os.Signal, 1)
		// signal.Notify(stop, os.Interrupt)
		// <-stop

		err = goBot.ChannelMessageDelete(thread.ID, messsage.ID)

		if err != nil {
			fmt.Println("Error deleting message")
			return
		}

	}
}
