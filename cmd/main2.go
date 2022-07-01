package main

import (
	"log"

	"github.com/chrispassas/peak"
)

func main() {
	defer func() {
		log.Printf("peak mem:%d\n", peak.PeakMemory())
		log.Printf("peak goroutines:%d\n", peak.PeakGoRoutines())
		log.Printf("peak fd:%d\n", peak.PeakFileDescriptors())
	}()

	// Program goes here
}
