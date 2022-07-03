package peak

import (
	"testing"
	"time"
)

func TestPeakMemory(t *testing.T) {
	if PeakMemory() == 0 {
		t.Errorf("PeakMemory() should not be 0")
	}
}

func TestPeakFileDescriptors(t *testing.T) {
	if PeakFileDescriptors() == 0 {
		t.Errorf("PeakFileDescriptors() should not be 0")
	}
}

func TestPeakGoRoutines(t *testing.T) {
	if PeakGoRoutines() == 0 {
		t.Errorf("PeakGoRoutines() should not be 0")
	}
}

func TestReset(t *testing.T) {
	Reset()
}

func TestSetInterval(t *testing.T) {
	SetInterval(time.Second * 2)
}

func TestPeakMemoryString(t *testing.T) {
	if out := PeakMemoryString(); len(out) == 0 {
		t.Errorf("PeakMemoryString() should not return empty string")
	}
}
