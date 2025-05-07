package std

// FuzzyMap is a type of map that provides primitive searching methods.
type FuzzyMap[K comparable, V any] map[K]V

// Search finds map entries that test 'true' with the provided predicate, then
// returns the found entries.
//
// If a nil predicate is provided, a nil map is returned.
func (m FuzzyMap[K, V]) Search(predicate func(K, V) bool) map[K]V {
	if predicate == nil {
		return nil
	}

	results := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return results
}
