package main

import (
	"fmt"
	"time"
)

func chanel() {
	var msg chan string
	fmt.Println(msg)

	msg <- "Chanel Ninja"

	msg = make(chan string)
	fmt.Println(msg)

	go func() {
		time.Sleep(2 * time.Second)
		msg <- "Chanel Ninja"
	}()

	fmt.Println(<-msg)
}