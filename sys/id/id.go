package id

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/std/gender"
	"github.com/ignite-laboratories/core/sys/log"
	"math/rand"
	"regexp"
	"strings"
)

var ModuleName = "id"

//go:embed names.csv
var nameDB string

// Names provides a collection of cultural names for seeding identifiers with.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database!
// https://incompetech.com
//
// Please check his stuff out, he's quite clever!
var Names = make(NameDB, 0, 8888)

type NameDB []std.GivenName

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

			log.Verbosef(ModuleName, "error reading name database: %v\n", err)
			panic(err)
		}

		genderFunc := func(s string) gender.Gender {
			if s == "Male" {
				return gender.Male
			} else if s == "Female" {
				return gender.Female
			} else {
				return gender.NonBinary
			}
		}

		entry := std.GivenName{
			Name:        strings.TrimSpace(record[0]),
			Description: strings.TrimSpace(record[3]),
			Details: struct {
				Origin string
				Gender gender.Gender
			}{
				Origin: strings.TrimSpace(record[1]),
				Gender: genderFunc(strings.TrimSpace(record[2])),
			},
		}
		Names = append(Names, entry)

		i++
	}

	log.Verbosef(ModuleName, "name database loaded\n")
}

// LookupName finds the provided name in the Names slice, otherwise it returns nil.
func LookupName(name string, caseInsensitive ...bool) (std.GivenName, error) {
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
	return std.GivenName{}, fmt.Errorf("name not found")
}

// RandomName returns a random name from the Names slice.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
func RandomName(database ...NameDB) std.GivenName {
	if len(database) > 0 {
		names := database[0]
		return names[rand.Intn(len(names))]
	}
	return Names[rand.Intn(len(Names))]
}

// RandomNameFiltered returns a random name from the Names slice which passes the provided predicate check.
//
// If you'd prefer a random name from your own name database, provide it as a parameter.
func RandomNameFiltered(predicate func(std.GivenName) bool, database ...NameDB) std.GivenName {
	for {
		name := RandomName(database...)
		if predicate(name) {
			return name
		}
	}
}

/**
tiny
*/

var _usedTinyNames = make(map[string]*std.GivenName)

// RandomTinyName is a standard function for returning a name which satisfies tiny's requirements for implicit naming.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word names)
//   - At least three characters in length
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
//
// NOTE: This guarantees up to 2¹⁴ unique names before it begins recycling names.
func RandomTinyName() std.GivenName {
	return RandomNameFiltered(tinyNameFilter)
}

func tinyNameFilter(name std.GivenName) bool {
	var nonAlphaRegex = regexp.MustCompile(`^[a-zA-Z]+$`)

	if len(_usedTinyNames) >= 1<<14 {
		_usedTinyNames = make(map[string]*std.GivenName)
	}

	if nonAlphaRegex.MatchString(name.Name) && _usedTinyNames[name.Name] == nil && len(name.Name) > 2 {
		_usedTinyNames[name.Name] = &name
		return true
	}
	return false
}
