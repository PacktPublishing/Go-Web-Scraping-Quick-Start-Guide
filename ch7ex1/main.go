package main

import (
	"fmt"
	"time"
)

func startTicker() {
	ticks := 0
	for true {
		fmt.Println(ticks)
		ticks++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	println("Starting ticker")
	go startTicker()
	time.Sleep(10 * time.Second)
}
