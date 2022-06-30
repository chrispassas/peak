package main

import (
	"log"
	"os"
	"time"

	"github.com/chrispassas/peak"
)

func main() {
	log.Printf("start")
	log.Printf("pid:%d", os.Getpid())

	// os.Stderr.Close()
	os.Stdout.Close()
	// os.Stdin.Close()

	for x := 0; x < 10; x++ {
		log.Printf("x:%d", x)
		time.Sleep(time.Second * 2)
		log.Printf("peak mem:%d", peak.PeakMemory())
		log.Printf("peak goroutines:%d", peak.PeakGoRoutines())
		log.Printf("peak fd:%d", peak.PeakFileHandles())
	}

	log.Printf("end")
}
