package std

// MouseState provides the current state of the mouse.
type MouseState struct {
	GlobalPosition XY[int]
	WindowPosition XY[int]
	Buttons        struct {
		Left, Middle, Right bool
	}
}
