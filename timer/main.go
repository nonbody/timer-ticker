package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				log.Println("sleep 1s")
				time.Sleep(time.Second)
				timer = time.NewTimer(3 * time.Second)
			}
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	log.Println("started")
	<-s

	log.Println("gracefully shutting down...")

}
