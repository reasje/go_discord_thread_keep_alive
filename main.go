package main

import (
	"fmt"
	// "strconv"
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

// var quit chan bool

// This app is designed to work with a single ticker
// This ticker handles the message sending action 
var ticker *time.Ticker 

func threadKeepAliveBackgroundTask(c chan string) {
	
	for {
	 m := <-c
		
		if m == config.REFRESH {
			fmt.Println("**Refreshing On down**")
			// if there is no an active ticker running
			// then It means
			if ticker != nil {
				ticker.Stop()
				go startTicker()
			}else{
				go startTicker()
			}
		}else {
			fmt.Println("Invalid Signal")
		}
	}
	
}

// This will start the operation and assign the ticker 
// to the active ticker  
func startTicker()   {
	ReadAndSend()
	ticker = time.NewTicker(10 * time.Second)
	for _ = range ticker.C {
		ReadAndSend()
	}
}

func ReadAndSend() {
						// reading threads from file
	err := entity.ReadThreads()

	if err != nil {
		fmt.Println("Error reading threads ")
		return
	}

	bot.SendMessage()
}

func getUserResfreshBackgroundTask(quit chan string) {
	for {
		var userInput string
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput == "r" {
			fmt.Println("Restating thread keep alive task")
			// sc := make(chan os.Signal, 1)
			// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
			// <-sc
			quit <- config.REFRESH
			// go threadKeepAliveBackgroundTask()
		}
	}
}

func main() {

	// for  {

		err := config.ReadConfig()

		if err != nil {
			fmt.Println(err.Error())

		}

		bot.Start()

		c := make(chan string)
		// quitFinal := make(chan bool)

		// keeping the task in the background up and running
		go threadKeepAliveBackgroundTask(c)

		go getUserResfreshBackgroundTask(c)

		c <- config.REFRESH

		// for range quitStart{
		// 	// quitValue   := <- quitStart
		// 	// quitString := strconv.FormatBool(quitValue)
		// 	quitFinal <- true
		// }
		// quitValue , ok  := <- quitStart
		// if ok {
		// 	quitString := strconv.FormatBool(quitValue)
		// 	fmt.Println("**Shutting down* %v *" +  quitString)
		// 	quitFinal <- true
		// }


		// this keeps our threads running
		select {}

	// }

}
