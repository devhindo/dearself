package main

import (
	"github.com/devhindo/dearself/server/config"
	"github.com/devhindo/dearself/server/api"
)

// todo make it a gin framework
// todo also use the todo extension

func main() {
	config.LoadEnv()
	api.RUN()
}





