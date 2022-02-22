package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/reasje/go_discord_thread_saver/config"
	"github.com/reasje/go_discord_thread_saver/entity"
)

var BotId string
var goBot *discordgo.Session

func Start() {

    //creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	

	err = goBot.Open()
//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    //If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

func SendMessage(c chan int) {
	for _ , thread := range entity.Threads {
		fmt.Println(thread.ID)
		_ , err := goBot.ChannelMessageSend(thread.ID, "Whatever , should be deleted ")

		if err != nil {
			fmt.Println("Error sending message")
			return
		}

		goBot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
		err = goBot.Open()
		if err != nil {
			log.Fatalf("Cannot open the session: %v", err)
		}
		defer goBot.Close()
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop

		// err = goBot.ChannelMessageDelete(thread.ID, messsage.ID)

		// if err != nil {
		// 	fmt.Println("Error deleting message")
		// 	return
		// }

	}
	c <- 1
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	
    //Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
  //If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
}