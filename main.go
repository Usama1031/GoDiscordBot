package main

import (
	"fmt"

	"github.com/usama1031/golangbot/bot"
	"github.com/usama1031/golangbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
