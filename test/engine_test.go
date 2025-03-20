package test

import (
	"github.com/ignite-laboratories/core"
	"testing"
	"time"
)

func Test_Engine_CannotStartWhileAlreadyRunning(t *testing.T) {
	// Fire engine A
	go func() {
		err := core.Impulse.Spark()
		if err != nil {
			t.Error("Expected no error, got one")
		}
	}()

	// Mute for a moment
	time.Sleep(100)

	// Fire engine B
	err := core.Impulse.Spark()
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Shut them both down
	core.Shutdown(100)
}
