package main

import (
	"fmt"
	"time"
)

func main() {

	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "chanel 1. 200ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			message2 <- "chanel 2. 1 second"
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}

}
