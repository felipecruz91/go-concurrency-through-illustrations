package main

import "fmt"

func main() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}

	foundOre := finder(theMine)
	minedOre := miner(foundOre)
	smelter(minedOre)
}

func finder(source []string) []string {

	var ores []string

	for _, item := range source {

		if item == "ore" {
			ores = append(ores, "ore")
		}
	}

	fmt.Printf("From Finder: %v\n", ores)

	return ores
}

func miner(source []string) []string {

	var minedOre []string

	for _, item := range source {

		if item == "ore" {
			minedOre = append(minedOre, "minedOre")
		}
	}

	fmt.Printf("From Miner: %v\n", minedOre)

	return minedOre
}

func smelter(source []string) {

	var smeltedOre []string

	for _, item := range source {

		if item == "minedOre" {
			smeltedOre = append(smeltedOre, "smeltedOre")
		}
	}

	fmt.Printf("From Smelter: %v\n", smeltedOre)
}
