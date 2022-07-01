package peak

import (
	"os"
	"runtime"
	"sync"
	"time"
)

var m2 sync.RWMutex
var interval = time.Second * 2

var m sync.RWMutex
var peakMemory uint64
var peakGoRoutines uint64
var peakFileDescriptors uint64

func init() {
	go update()
}

func update() {
	for {
		updatePeakMemory()
		updatePeakGoroutines()
		updateFileDescriptors()
		m2.RLock()
		time.Sleep(interval)
		m2.RUnlock()
	}
}

func updatePeakMemory() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	m.Lock()
	defer m.Unlock()
	if mem.Alloc > peakMemory {
		peakMemory = mem.Alloc
	}
}

func updatePeakGoroutines() {
	count := uint64(runtime.NumGoroutine())
	m.Lock()
	defer m.Unlock()
	if count > peakGoRoutines {
		peakGoRoutines = count
	}
}

func updateFileDescriptors() {
	var fdPath string
	switch runtime.GOOS {
	case "darwin":
		fdPath = "/dev/fd"
	case "linux", "freebsd":
		fdPath = "/proc/self/fd"
	default:
		return
	}

	entries, _ := os.ReadDir(fdPath)
	m.Lock()
	defer m.Unlock()
	if peakFileDescriptors < uint64(len(entries)) {
		peakFileDescriptors = uint64(len(entries))
	}
}

// PeakMemory return highest number of bytes in use
func PeakMemory() uint64 {
	m.RLock()
	defer m.RUnlock()
	return peakMemory
}

// PeakGoRoutines return highest number of observed goroutines
func PeakGoRoutines() uint64 {
	m.RLock()
	defer m.RUnlock()
	return peakGoRoutines
}

// PeakFileDescriptors return highest number of observed file descriptors
func PeakFileDescriptors() uint64 {
	m.RLock()
	defer m.RUnlock()
	return peakFileDescriptors
}

// Reset set metrics back to 0
func Reset() {
	m.Lock()
	defer m.Unlock()
	peakMemory = 0
	peakGoRoutines = 0
	peakFileDescriptors = 0
}

// SetInterval override default sampling interval
func SetInterval(dur time.Duration) {
	if dur > 0 {
		m2.Lock()
		interval = dur
		m2.Unlock()
	}
}
