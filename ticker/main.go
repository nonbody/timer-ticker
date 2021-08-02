package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ticker := time.NewTicker(3 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("sleep 1s")
				time.Sleep(time.Second)
			}
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	log.Println("started")
	<-s

	log.Println("gracefully shutting down...")

}
