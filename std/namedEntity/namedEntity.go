package std

import "github.com/ignite-laboratories/core/std/name"

// NamedEntity provides a given name to the entity.
type NamedEntity struct {
	name.GivenName
	Entity
}

// New creates a new entity, assigns it a unique identifier, and gives it a name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from core.Names is chosen.  If you'd prefer to use a different random
// name database, please see NewFromDB.
func New(name ...name.GivenName) NamedEntity {
	var given name.GivenName
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

// NewFromDB creates a new entity, assigns it a unique identifier, and gives it a random
// name from the provided name database.  If no database is provided, the default database is used.
//
// If you'd prefer to name your entity directly, please see New.
func NewFromDB(db ...NameDB) NamedEntity {
	given := RandomName(db...)

	ne := NamedEntity{
		GivenName: given,
	}
	ne.ID = NextID()

	return ne
}
