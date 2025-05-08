package core

import "sync"

// FilterableMap is a type of map that provides primitive filtration methods.
type FilterableMap[K comparable, V any] map[K]V

// FilterableSlice is a type of slice that provides primitive filtration methods.
type FilterableSlice[T any] []T

// Where is a filtration method for maps. It calls the predicate for every entry
// in the map and returns a new map containing only the entries for which the
// predicate returned true.  This is effectively a "Where" clause.
//
// If a nil predicate is provided, an empty map is returned.
func (m FilterableMap[K, V]) Where(predicate func(K, V) bool) map[K]V {
	results := make(map[K]V)
	if predicate == nil {
		return results
	}

	for k, v := range m {
		if predicate(k, v) {
			results[k] = v
		}
	}
	return results
}

// Where is a filtration method for slices.  It calls the predicate for every entry
// in the slice and returns a new slice containing only the entries for which the
// predicate returned true.  This is effectively a "Where" clause.
//
// If a nil predicate is provided, an empty slice is returned.
func (s FilterableSlice[T]) Where(predicate func(int, T) bool) []T {
	if predicate == nil || len(s) == 0 {
		return make([]T, 0)
	}

	if len(s) < 1024 {
		results := make([]T, 0, len(s))
		for i, v := range s {
			if predicate(i, v) {
				results = append(results, v)
			}
		}
		return results
	}

	resultChan := make(chan T, len(s))
	chunkSize := 1024
	numGoroutines := (len(s) + chunkSize - 1) / chunkSize

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(s) {
			end = len(s)
		}

		go func(chunk []T, baseIndex int) {
			defer wg.Done()
			for i, item := range chunk {
				if predicate(baseIndex+i, item) {
					resultChan <- item
				}
			}
		}(s[start:end], start)
	}

	wg.Wait()
	
	results := make([]T, 0, len(s))
	for item := range resultChan {
		results = append(results, item)
	}

	return results
}
