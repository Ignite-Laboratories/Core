package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize SDL2
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %s\n", err)
	}
	defer sdl.Quit()

	// Create a full-screen window
	window, err := sdl.CreateWindow(
		"SDL2 Full-Screen Example",    // Window title
		sdl.WINDOWPOS_CENTERED,        // Window position X
		sdl.WINDOWPOS_CENTERED,        // Window position Y
		1920,                          // Window width (ignored in full-screen mode)
		1080,                          // Window height (ignored in full-screen mode)
		sdl.WINDOW_FULLSCREEN_DESKTOP, // Full-screen mode flag
	)
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
	}
	defer window.Destroy()

	// Variable to manage the running state
	running := true

	// Create a channel to listen for OS signals
	signalChan := make(chan os.Signal, 1)

	// Notify the channel when a SIGINT (Ctrl+C) or SIGTERM is received
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Wait for a signal
		sig := <-signalChan
		fmt.Printf("\nReceived signal: %s. Exiting...\n", sig)
		running = false
	}()

	go func() {
		time.Sleep(time.Second * 20)
		fmt.Println("terminating")
		running = false
	}()

	// Run the event loop
	for running {
		// Handle SDL events (e.g., window close or keyboard/mouse input)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // Handle quit event (e.g., window close button)
				running = false
			}
		}

		// Query and print the mouse position
		x, y, _ := sdl.GetMouseState()
		fmt.Printf("Mouse position: X = %d, Y = %d\n", x, y)

		// Limit the loop frequency a bit
		sdl.Delay(16) // ~60 FPS
	}

	fmt.Println("Exiting program...")
}
