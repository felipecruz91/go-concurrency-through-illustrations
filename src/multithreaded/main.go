package main

import (
	"fmt"
	"sync"
)

func main() {

	theMine := []string{"rock", "ore", "ore", "rock", "ore"}

	goRoutines := 3
	var wg sync.WaitGroup

	// Tell the 'wg' WaitGroup how many threads/goroutines that are about to run concurrently.
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		// Spawn a thread for each iteration in the loop.
		go func(source []string, i int) {
			defer wg.Done()
			for _, item := range source {
				if item == "ore" {
					fmt.Printf("From Finder %d: Finded ore!\n", i)
				}
			}
		}(theMine, i)
	}

	// Wait for `wg.Done()` to be exectued the number of times
	//   specified in the `wg.Add()` call.
	// `wg.Done()` should be called the exact number of times
	//   that was specified in `wg.Add()`.
	// If the numbers do not match, `wg.Wait()` will either
	//   hang infinitely or throw a panic error.
	wg.Wait()

	fmt.Println("Finished for loop")
}
