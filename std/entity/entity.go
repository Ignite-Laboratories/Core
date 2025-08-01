package std

import (
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/name"
	"github.com/ignite-laboratories/core/sys/id"
)

// NewNamed creates a new entity, assigns it a unique identifier, and gives it a random name.
//
// If you'd prefer to directly name your entity, provide it as a parameter here.  Otherwise,
// a random entry from core.Names is chosen.  If you'd prefer to use a different random
// name database, please see NewNamedFromDB.
func NewNamed(str ...std.GivenName) std.NamedEntity {
	var given std.GivenName
	if len(str) > 0 {
		given = str[0]
	} else {
		given = name.Random()
	}

	ne := std.NamedEntity{
		GivenName: given,
	}
	ne.ID = id.Next()

	return ne
}

// NewNamedFromDB creates a new entity, assigns it a unique identifier, and gives it a random
// name from the provided name database.  If no database is provided, the default database is used.
//
// If you'd prefer to name your entity directly, please see NewNamed.
func NewNamedFromDB(db ...[]std.GivenName) std.NamedEntity {
	given := name.Random(db...)

	ne := std.NamedEntity{
		GivenName: given,
	}
	ne.ID = id.Next()

	return ne
}
