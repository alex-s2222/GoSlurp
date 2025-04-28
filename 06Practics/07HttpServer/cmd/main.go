package main

import (
	"context"
	"httpServer/internals/app"
	"httpServer/internals/cfg"
	"log"
	"os"
	"os/signal"
)

func main(){
	config := cfg.LoadAndStoreConfig()
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c
		log.Printf("system call: %v", oscall)
		server.ShutDown()
		cancel()
	}()

	server.Serve()
}