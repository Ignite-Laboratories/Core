package core

// Entity provides 'ID' and 'Name' fields to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID uint64

	// Name is the given name for this entity.  It may or may not be assigned.
	Name string
}
