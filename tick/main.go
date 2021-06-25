package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go func() {
		for now := range time.Tick(2 * time.Second) {
			log.Println(now)
			log.Println("sleep")
			time.Sleep(5 * time.Second)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	log.Println("started")
	<-s

	log.Println("gracefully shutting down...")

}
