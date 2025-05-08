package core

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed names.csv
var nameDB string

// Names provides a collection of cultural names for seeding identifiers with.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database!
//
// Please check his stuff out, he's quite clever!
var Names = make(NameDB, 0, 8888)

type NameDB []GivenName

// GivenName represents a name, as well as its original cultural meaning.
//
// Your interpretation and meaning may absolutely vary. The true beauty of language
// is in such prismatic interpretations based entirely upon contextual experiences <3
//
//	tl;dr - you own your identifier, not the other way around!
type GivenName struct {
	Name        string
	Description string
	Details     struct {
		Origin string
		Gender Gender
	}
}

// NewName creates a new GivenName.  You may optionally provide a description during creation.
func NewName(name string, description ...string) GivenName {
	if len(description) > 0 {
		return GivenName{
			Name: name,
		}
	}
	return GivenName{
		Name:        name,
		Description: description[0],
	}
}

// Gender provides global identifiers for Male, Female, or NonBinary interpretations - as gender
// is ultimately self-defined.
type Gender int

const (
	Female Gender = iota
	Male
	NonBinary
)

func (n GivenName) String() string {
	if n.Description == "" {
		return fmt.Sprintf("%v", n.Name)
	}
	return fmt.Sprintf("%v - %v", n.Name, n.Description)
}

func initializeNameDB() {
	reader := csv.NewReader(strings.NewReader(nameDB))
	reader.Comma = '\t'

	i := 0
	for {
		record, err := reader.Read() // Read a single line
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file
			}

			Verbosef(ModuleName, "error reading name database: %v\n", err)
			panic(err)
		}

		genderFunc := func(s string) Gender {
			if s == "Male" {
				return Male
			} else if s == "Female" {
				return Female
			} else {
				return NonBinary
			}
		}

		entry := GivenName{
			Name:        strings.TrimSpace(record[0]),
			Description: strings.TrimSpace(record[3]),
			Details: struct {
				Origin string
				Gender Gender
			}{
				Origin: strings.TrimSpace(record[1]),
				Gender: genderFunc(strings.TrimSpace(record[2])),
			},
		}
		Names = append(Names, entry)

		i++
	}

	Verbosef(ModuleName, "name database loaded\n")
}

// LookupName finds the provided name in the Names slice, otherwise it returns nil.
func LookupName(name string, caseInsensitive ...bool) (GivenName, error) {
	for _, n := range Names {
		if len(caseInsensitive) > 0 && caseInsensitive[0] {
			if strings.EqualFold(n.Name, name) {
				return n, nil
			}
		} else {
			if n.Name == name {
				return n, nil
			}
		}
	}
	return GivenName{}, fmt.Errorf("name not found")
}

// RandomName returns a random name from the Names slice.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
func RandomName(database ...NameDB) GivenName {
	if len(database) > 0 {
		names := database[0]
		return names[rand.Intn(len(names))]
	}
	return Names[rand.Intn(len(Names))]
}
