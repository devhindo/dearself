package main

import (
	//"fmt"
	//"net/http"
	//"time"

	"github.com/devhindo/dearself/server/config"
	//"github.com/devhindo/dearself/server/api"
    "github.com/devhindo/dearself/server/db"
)

// todo make it a gin framework
// todo also use the todo extension

func main() {

    config.LoadEnv()
    db.RunScheduledMailsJob()

    /*

	go func() {
        for {
            _, err := http.Get("https://dearself.onrender.com/ping")
            if err != nil {
                // Handle error
                fmt.Println(err)
            }
            // Wait for 14 minutes before sending the next request
            time.Sleep(14 * time.Minute)
			fmt.Println("pinged")
        }
    }()

    go func() {
        for {
            db.RunScheduledMailsJob()
            // send scheduled mails every 24 hours
            time.Sleep(24 * time.Hour)
        }
    }()

	api.RUN()
    */
}





