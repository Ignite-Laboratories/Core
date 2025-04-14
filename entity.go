package core

// Entity provides 'ID' and 'Name' fields to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID uint64
}

type NamedEntity struct {
	GivenName
	Entity
}

// NewNamedEntity creates a new entity, assigns it a unique identifier, and gives it a name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from core.Names is chosen.
func NewNamedEntity(name ...GivenName) NamedEntity {
	var given GivenName
	if len(name) > 0 {
		given = name[0]
	} else {
		given = RandomName()
	}

	ne := NamedEntity{
		GivenName: given,
	}
	ne.ID = NextID()

	return ne
}
