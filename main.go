package main

import (
	"fmt"
	"strings"

	"github.com/reasje/go_discord_thread_saver/internal/metadata"
	"github.com/reasje/go_discord_thread_saver/internal/storage"
	"github.com/reasje/go_discord_thread_saver/services/bot"
	"github.com/reasje/go_discord_thread_saver/services/config"
	"github.com/reasje/go_discord_thread_saver/services/ticker"
)

// IDEA:  resstructure the project to make it more readable (handle protocols)
// IDEA: handeling failures (retry till success )



func threadKeepAliveBackgroundTask(c chan string) {
	
	for {
	 m := <-c
		
		if m == metadata.REFRESH {
			// if there is no an active ticker running
			// then It means
			if storage.Ticker != nil {
				ticker.StopTicker()
				go ticker.StartTicker()
			}else{
				go ticker.StartTicker()
			}
		}else {
			fmt.Println("Invalid Signal")
		}
	}
	
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
			quit <- metadata.REFRESH
			// go threadKeepAliveBackgroundTask()
		}
	}
}

func main() {
		err := config.ReadConfig()

		if err != nil {
			fmt.Println(err.Error())

		}

		bot.StartBot()

		// NOTE only one channel is used fo controlling app
		c := make(chan string)

		// keeping the task in the background up and running
		go threadKeepAliveBackgroundTask(c)

		// keeping the user input task running in the background
		go getUserResfreshBackgroundTask(c)

		// starts the process
		c <- metadata.REFRESH

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
