package main

import (
	"fmt"
	"strings"
)

func genetateMsg(ch1 chan<- string) {

	ch1 <- "Msg has been generated from GenFunction"

}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	msg := <-ch1
	msg = strings.ToTitle(msg)
	ch2 <- msg + " and modified in Relay function"

}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go genetateMsg(ch1)

	go relayMsg(ch1, ch2)

	msg := <-ch2
	fmt.Println(msg)

}
