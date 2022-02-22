package main

import (
	"fmt"
	"strings"

	"time"

	"github.com/reasje/go_discord_thread_saver/entity"

	// "log"
	// "os"
	// "os/signal"
	// "github.com/bwmarrin/discordgo"
	"github.com/reasje/go_discord_thread_saver/bot"
	"github.com/reasje/go_discord_thread_saver/config"
)

func threadKeepAliveBackgroundTask() {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {

		// reading threads from file
		err := entity.ReadThreads()

		if err != nil {
			fmt.Println("Error reading threads ")
			return
		}

		bot.SendMessage()
	}
}

func getUserResfreshBackgroundTask() {
	for {
		var userInput string
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput == "r" {
			fmt.Println("List of games:")
		}
	}
}

func main() {


	err := entity.ReadThreads()

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	err = config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	// keeping the task in the background up and running
	go threadKeepAliveBackgroundTask()

	go getUserResfreshBackgroundTask()


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
