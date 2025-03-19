package test

import (
	"github.com/ignite-laboratories/core"
	"testing"
	"time"
)

func Test_Engine_CannotStartWhileAlreadyRunning(t *testing.T) {
	engine := core.NewEngine()

	go func() {
		err := engine.Start()
		if err == nil {
			t.Error("Expected no error, got one")
		}
	}()

	time.Sleep(100)

	go func() {
		err := engine.Start()
		if err == nil {
			t.Error("Expected error, got nil")
		}
	}()

	DelayKill(100)
}
