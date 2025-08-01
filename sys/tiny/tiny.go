package tiny

import (
	"github.com/ignite-laboratories/core/std"
	"regexp"
)

var _usedTinyNames = make(map[string]*std.GivenName)

// NameFilter is a standard function for returning a name which satisfies tiny's requirements for implicit naming.
// Currently, these are our explicit filters -
//
//   - Only the standard 26 letters of the English alphabet (case-insensitive)
//   - No whitespace or special characters (meaning only single word alpha-explicit names)
//   - At least three characters in length
//   - At least 2¹⁴ unique names before beginning to recycling names
//   - Names are case-sensitive in uniqueness.
//
// These filters will never be reduced - if any changes are made, they will only be augmented.
func NameFilter(name std.GivenName) bool {
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
