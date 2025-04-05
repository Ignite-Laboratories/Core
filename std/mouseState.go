package std

// MouseState provides the current state of the mouse.
type MouseState struct {
	Position XY[int]
	Buttons  struct {
		Left, Middle, Right bool
	}
}
