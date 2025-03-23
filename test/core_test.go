package test

import (
	"github.com/ignite-laboratories/core"
	"testing"
	"time"
)

func Test_Core_Initialization(t *testing.T) {
	if !core.Alive {
		t.Error("Expected core.Alive to be true, but got false")
	}
	if core.ID != 1 {
		t.Error("Expected core.ID to be 1, but got ", core.ID)
	}
}

func Test_Core_Shutdown(t *testing.T) {
	core.Shutdown(0)
	time.Sleep(time.Millisecond * 100)
	if core.Alive {
		t.Error("Shutdown did not set Alive to false")
	}
	core.Alive = true // reset for the next test
}

func Test_Core_Shutdown_ShouldDelay(t *testing.T) {
	core.Shutdown(time.Millisecond * 100)
	if !core.Alive {
		t.Error("Shutdown did not delay")
	}
	time.Sleep(time.Millisecond * 100) // ensure shutdown triggers
	core.Alive = true                  // reset for the next test
}

func Test_Core_Shutdown_DelayShouldWork(t *testing.T) {
	core.Shutdown(0)
	time.Sleep(time.Millisecond * 100)
	if core.Alive {
		t.Error("Shutdown did not set Alive to false")
	}
	core.Alive = true // reset for the next test
}
