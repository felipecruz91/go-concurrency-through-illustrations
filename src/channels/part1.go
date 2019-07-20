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
				fmt.Println("[Finder] - Ore found!")
				oreChan <- "ore"
			} else {
				fmt.Println("[Finder] - :(")
			}
		}
	}(theMine)

	// Miner
	go func() {

		//  Ranging over a channel will block until another item is sent on the channel
		// The only way to stop the go routine from blocking after all sends have occurred is by
		// closing the channel with ‘close(channel)’
		for foundOre := range oreChan {
			fmt.Println("[Miner] - Received: " + foundOre)

			// Simulate time to mine the ore
			time.Sleep(2 * time.Second)

			fmt.Println("[Miner] - Mine op. complete.")

			minedOreChan <- "minedOre"
		}

	}()

	// Smelter
	go func() {

		for smeltedOre := range minedOreChan {
			fmt.Println("[Smelter] - Received " + smeltedOre)
		}
	}()

	<-time.After(10 * time.Second)
}
