package main

import (
	"fmt"
	"time"
)

// Given a moment, determine the moment that would be after a gigasecond has passed.
// A gigasecond is 10^9 (1,000,000,000) seconds.
func main() {
	epoch := time.Now().UTC()
	oneGigasecond := time.Second * 1000000000
	epochElapsed := epoch.Add(oneGigasecond)

	epochFormated := epoch.UTC().Format("02/01/2006 15:04:05")
	epochElapsedFormated := epochElapsed.UTC().Format("02/01/2006 15:04:05")

	fmt.Printf("Elapsed time from %v \nThis is future: %v \n", epochFormated, epochElapsedFormated)
}
