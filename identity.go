package core

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed nameDB.csv
var nameDB string

// Names provides a collection of cultural names for seeding identifiers with.
//
// All credit goes to Kevin MacLeod of Incompetech for such a wonderful source database!
//
// Please check his stuff out, he's quite clever!
var Names = make([]GivenName, 0, 8888)

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

			if len(record) < 5 {
				continue // Some entries might be blank
			}

			if len(record) > 5 {
				record = record[:5]
				continue // Some rows might have extra columns
			}

			Verbosef(ModuleName, "error reading name database: %v\n", err)
			panic(err) // Otherwise, tell me the issue
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

		variationStrs := strings.Split(record[4], ", ")

		for _, variation := range variationStrs {
			if len(variation) > 0 {
				entry = GivenName{
					Name:        variation,
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
			}
		}

		i++
	}

	Verbosef(ModuleName, "name database loaded\n")
}

// LookupName finds the provided name in the Names slice, otherwise it returns nil.
func LookupName(name string) *GivenName {
	for _, n := range Names {
		if strings.EqualFold(n.Name, name) {
			return &n
		}
	}
	return nil
}

// RandomName returns a random name from the Names slice.
func RandomName() GivenName {
	return Names[rand.Intn(len(Names))]
}
