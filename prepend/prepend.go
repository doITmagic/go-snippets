// Package prepend
// Go have a built-in append function that allows for the appending of a value to a slice.
// Just like a joke Package prepend provides a method to prepend a value to a slice of interfaces
package prepend

// Prepender is an interface that defines the Prepend method
type Prepender []interface {
	Prepend(value interface{}, slice []interface{}) []interface{}
}

// Prepend prepends (add to the beginning at the slice) a value to a slice of interfaces
func (p Prepender) Prepend(v interface{}, slice []interface{}) []interface{} {
	return append([]interface{}{v}, slice...)
}
