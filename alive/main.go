package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var timer time.Timer

	cnt := 1
	go func() {
		for alive := true; alive; {
			log.Println("next")
			timer = *time.NewTimer(2 * time.Second)
			<-timer.C
			log.Println("round:", cnt)
			cnt++
			log.Println("sleep")
			time.Sleep(5 * time.Second)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	log.Println("started")
	<-s

	log.Println("gracefully shutting down...")
	log.Println("timer Stop:", timer.Stop())
}
