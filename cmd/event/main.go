package main

import (
	"fmt"
	"kumparan/config/env"
	"kumparan/handler"
	"runtime"
)

func main() {
	fmt.Println("[EVENT RUN]")
	env.LoadEnv()

	// Init dependencies
	service := handler.InitHandler()

	go service.ArticleModule.EventCreateArticle()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	runtime.Goexit()
}
