package main

import (
	"fmt"
)

func main() {
	var chans = []chan string{}
	agg := make(chan string)
	for _, ch := range chans {
		go func(c chan string) {
			for msg := range c {
				agg <- msg
			}
		}(ch)
	}

	select {
	case msg := <-agg:
		fmt.Println("received ", msg)
	}
}
