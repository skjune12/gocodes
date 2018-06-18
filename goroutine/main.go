package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")

	// channel
	finished := make(chan bool)

	go func() {
		log.Println("Sleep1 Started.")
		time.Sleep(time.Second * 1)
		log.Println("Sleep1 Finished.")
		finished <- true
	}()

	go func() {
		log.Println("Sleep2 Started.")
		time.Sleep(time.Second * 2)
		log.Println("Sleep2 Finished.")
		finished <- true
	}()

	go func() {
		log.Println("Sleep3 Started.")
		time.Sleep(time.Second * 3)
		log.Println("Sleep3 Finished")
		finished <- true
	}()

	for i := 0; i < 3; i++ {
		<-finished
	}

	log.Println("all finished")
}
