package core

// System is a structure that can manage impulse activation.
type System struct {
	Entity
	*Activation
	*Engine
}
