package main

import (
	"dot-go/src"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	server := src.InitServer()

	go func() {
		server.Run()

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

}
