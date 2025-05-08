package core

// Entity provides an 'ID' field to any composite types.
type Entity struct {
	// ID is the unique identifier for this entity, relative to its home world.
	ID uint64
}

// NamedEntity provides a given name to the entity.
type NamedEntity struct {
	GivenName
	Entity
}

// NewNamedEntity creates a new entity, assigns it a unique identifier, and gives it a name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from core.Names is chosen.  If you'd prefer to use a different random
// name database, please see NewNamedEntityFromDB.
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

// NewNamedEntityFromDB creates a new entity, assigns it a unique identifier, and gives it a random
// name from the provided name database.  If no database is provided, the default database is used.
//
// If you'd prefer to name your entity directly, please see NewNamedEntity.
func NewNamedEntityFromDB(db ...NameDB) NamedEntity {
	given := RandomName(db...)

	ne := NamedEntity{
		GivenName: given,
	}
	ne.ID = NextID()

	return ne
}
