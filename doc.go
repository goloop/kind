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
// The scalar predicates (IsInt, IsString, ...) are leaf-aware: a container
// reports the leaf type of its element chain, so a []int reports IsInt, and
// for maps the leaf follows the value type, not the key. When you need a
// strict scalar test that excludes containers, compare Kind with the reflect
// kind instead:
//
//	k.Kind() == reflect.Int // strictly an int, not []int or map[K]int
//
// Type descriptors are cached, so repeated calls for the same reflect.Type do
// not repeat the recursive classification work. Value-specific checks such as
// IsNil and IsZero are still evaluated per value.
package kind
