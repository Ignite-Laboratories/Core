package core

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity.
	ID uint64

	// Name is the given name for this entity.  It may or may not be assigned.
	Name string
}
