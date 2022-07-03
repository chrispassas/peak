package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/chrispassas/peak"
)

func main() {
	fmt.Printf("start\n")
	// Print peak values
	fmt.Printf("peak mem:%d\n", peak.PeakMemory())
	fmt.Printf("peak mem string:%s\n", peak.PeakMemoryString())
	fmt.Printf("peak goroutines:%d\n", peak.PeakGoRoutines())
	fmt.Printf("peak fd:%d\n", peak.PeakFileDescriptors())

	time.Sleep(time.Second * 2)

	// Make some goroutines and use some memeory
	var wg sync.WaitGroup
	var data []string
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Second * 5)
			fmt.Printf("x:%d\n", x)
		}(x, &wg)
		for i := 0; i < 1000; i++ {
			data = append(data, fmt.Sprintf("%d", i))
		}
	}
	wg.Wait()

	fmt.Printf("peak mem:%d\n", peak.PeakMemory())
	fmt.Printf("peak mem string:%s\n", peak.PeakMemoryString())
	fmt.Printf("peak goroutines:%d\n", peak.PeakGoRoutines())
	fmt.Printf("peak fd:%d\n", peak.PeakFileDescriptors())

}
