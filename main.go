package main

import (
	"fmt"


	"github.com/reasje/go_discord_thread_saver/entity"
	"time"
	// "log"
	// "os"
	// "os/signal"
	"github.com/bwmarrin/discordgo"
	// "github.com/reasje/go_discord_thread_saver/bot"
	// "github.com/reasje/go_discord_thread_saver/config"
)

func backgroundTask() {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tick at")
		err := entity.ReadThreads()

		if err != nil {
			fmt.Println("Error reading threads ")
			return
		}
		s, err := discordgo.New("Bot " + "OTQ1MjMyODc2NjQxODY5ODM0.YhNKtw.4zt-cAvzBa9h65tQDCjmrX9dPd8")
		
		if err != nil {
			fmt.Println("Error sending message")
		}

		s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) { fmt.Println("Bot is ready") })
		message, err := s.ChannelMessageSend("945225329377697873", "Hello, world!")
		if err != nil {
			fmt.Println("Error sending message")
		}
		print(message)
		// s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
		// err = s.Open()
		// if err != nil {
		// 	log.Fatalf("Cannot open the session: %v", err)
		// }
		// defer s.Close()
		// stop := make(chan os.Signal, 1)
		// signal.Notify(stop, os.Interrupt)
		// <-stop
		// log.Println("Graceful shutdown")
		// sendMessageChannel := make(chan int)
		// go bot.SendMessage(sendMessageChannel)
		// <-sendMessageChannel
	}
}

func main() {


	err := entity.ReadThreads()

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	// err = config.ReadConfig()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// bot.Start()

	go backgroundTask()




	// go func() {

	// 	time.Sleep(5 * time.Second)
	// 	<-ticker.C
	// 	fmt.Println("Tick at")

	// }()

	// for {
	// 	var userInput string
	// 	fmt.Scanln(&userInput)
	// 	userInput = strings.ToLower(userInput)
	// 	if userInput == "r" {
	// 		// TODO get the list from file  agian
	// 		fmt.Println("List of games:")
	// 	}
	// }

	select {}

}
