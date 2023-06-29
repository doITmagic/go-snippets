package slices

// SameUnderlying check if the first slice and the second slice have the same underlying array
func SameUnderlying[T any](first, second []T) bool {
	return cap(first) > 0 && cap(second) > 0 && &first[0:cap(first)][cap(first)-1] == &second[0:cap(second)][cap(second)-1]
}
