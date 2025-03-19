package core

// Entity can be embedded into your struct to gain an ID field and GetID() function.
type Entity struct {
	// ID is this entities unique identifier.
	ID uint64
}

// GetID returns this entity's unique ID.
func (e Entity) GetID() uint64 {
	return e.ID
}
