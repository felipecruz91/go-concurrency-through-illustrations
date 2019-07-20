package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	oreChan := make(chan string)
	minedOreChan := make(chan string)

	// Finder
	go func(source []string) {

		for _, item := range source {
			if item == "ore" {
				oreChan <- "ore"
			}
		}
	}(theMine)

	// Miner
	go func() {

		//  Ranging over a channel will block until another item is sent on the channel
		// The only way to stop the go routine from blocking after all sends have occurred is by
		// closing the channel with ‘close(channel)’
		for foundOre := range oreChan {
			fmt.Println("Miner received: " + foundOre)
			minedOreChan <- "minedOre"
		}

	}()

	// Smelter
	go func() {

		for smeltedOre := range minedOreChan {
			fmt.Println("Smelter: " + smeltedOre)
		}
	}()

	<-time.After(5 * time.Second)
}
