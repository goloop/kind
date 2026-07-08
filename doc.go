// Package kind provides a small, cached reflection layer for classifying Go
// values and types.
//
// The main entry point is Of, which returns a Kind describing the dynamic type
// and value-specific state of a value:
//
//	k := kind.Of([]int{1, 2, 3})
//	k.IsSlice() // true
//	k.IsInt()   // true, because the slice element is int
//	k.Name()    // "[]int"
//
// Type descriptors are cached, so repeated calls for the same reflect.Type do
// not repeat the recursive classification work. Value-specific checks such as
// IsNil and IsZero are still evaluated per value.
package kind
