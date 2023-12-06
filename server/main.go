package main

import (
	"github.com/devhindo/dearself/server/config"
	"github.com/devhindo/dearself/server/api"

	"fmt"
	"net/http"
	"time"
)

// todo make it a gin framework
// todo also use the todo extension

func main() {

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

	config.LoadEnv()
	api.RUN()
}





