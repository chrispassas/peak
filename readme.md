# Peak tracks peak usage of memory goroutines and file descriptors that have ocurred since program startup

## Why use this?
If you need to know how many CPU's , File Descriptors and Memory your program uses at its peak. Maybe for sizing a VM or container resources.

## Quick example
```go
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
```

## Example
```go
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

```

## Output
```
start
peak mem:120840
peak mem string:0.12 MB
peak goroutines:2
peak fd:0
x:1
x:0
x:8
x:4
x:2
x:3
x:6
x:5
x:7
x:9
peak mem:960272
peak mem string:0.92 MB
peak goroutines:12
peak fd:10
```
