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
	if core.Self.ID != 1 {
		t.Error("Expected core.Self.ID to be 1, but got ", core.Self.ID)
	}
}

func Test_Core_Shutdown(t *testing.T) {
	core.Shutdown(1)
	time.Sleep(100)
	if core.Alive {
		t.Error("Shutdown did not set Alive to false")
	}
}

func Test_Core_Shutdown_ShouldDelay(t *testing.T) {
	go core.Shutdown(250)
	if !core.Alive {
		t.Error("Shutdown did not delay")
	}
}

func Test_Core_Shutdown_DelayShouldWork(t *testing.T) {
	core.Shutdown(250)
	time.Sleep(1000)
	if core.Alive {
		t.Error("Shutdown did not set Alive to false")
	}
}
