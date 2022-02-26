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

var quit chan bool

func threadKeepAliveBackgroundTask(quit chan bool) {
		select {
			
			case <- quit:
				fmt.Println("**Shutting down**")
				return
			default:
				keepAlive()
				ticker := time.NewTicker(1 * time.Minute)
				for _ = range ticker.C {
					keepAlive()
				}
			
		}

	
}

func keepAlive() {
						// reading threads from file
	err := entity.ReadThreads()

	if err != nil {
		fmt.Println("Error reading threads ")
		return
	}

	bot.SendMessage()
}

func getUserResfreshBackgroundTask(quit chan bool) {
	for {
		var userInput string
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput == "r" {
			fmt.Println("Restating thread keep alive task")
			quit <- true
			// go threadKeepAliveBackgroundTask()
		}
	}
}

func main() {

	for true {

		err := config.ReadConfig()

		if err != nil {
			fmt.Println(err.Error())

		}

		bot.Start()

		quit = make(chan bool)

		// keeping the task in the background up and running
		go threadKeepAliveBackgroundTask(quit)

		go getUserResfreshBackgroundTask(quit)

		// this keeps our threads running
		select {}

	}

}
