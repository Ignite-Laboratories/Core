package core

// FilterableMap is a type of map that provides primitive filtration methods.
type FilterableMap[K comparable, V any] map[K]V

// FilterableSlice is a type of slice that provides primitive filtration methods.
type FilterableSlice[T any] []T

// Where is a filtration method for maps. It calls the predicate for every entry
// in the map and returns a new map containing only the entries for which the
// predicate returned true.  This is effectively a "Where" clause.
//
// If a nil predicate is provided, a nil map is returned.
func (m FilterableMap[K, V]) Where(predicate func(K, V) bool) map[K]V {
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

// Select is a filtration method for slices.  It calls the predicate for every entry
// in the slice and returns a new slice containing only the entries for which the
// predicate returned true.  This is effectively a "Select" statement.
//
// If a nil predicate is provided, a nil slice is returned.
func (s FilterableSlice[T]) Select(predicate func(int, T) bool) []T {
	if predicate == nil {
		return nil
	}

	results := make([]T, 0, len(s))
	for i, v := range s {
		if predicate(i, v) {
			results = append(results, v)
		}
	}
	return results
}
