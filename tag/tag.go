package tag

import (
	"errors"
	"fmt"
	"math/bits"
)

// Tag is a type that represents a property of a certain object's type.
type Tag uint64

const (
	// Nil represents nil value.
	Nil Tag = 1 << iota
	// Pointer represents pointer value.
	Pointer
	// Array represents array value.
	Array
	// ArrayOfArrays represents array of arrays value.
	ArrayOfArrays
	// ArrayOfSlices represents array of slices value.
	ArrayOfSlices
	// Slice represents slice value.
	Slice
	// SliceOfSlices represents slice of slices value.
	SliceOfSlices
	// SliceOfArrays represents slice of arrays value.
	SliceOfArrays
	// Struct represents struct value.
	Struct
	// Map represents map value.
	Map
	// Func represents function value.
	Func
	// Chan represents channel value.
	Chan
	// Bool represents boolean value.
	Bool
	// Int represents int value.
	Int
	// Int8 represents int8 value.
	Int8
	// Int16 represents int16 value.
	Int16
	// Int32 represents int32 value.
	Int32
	// Int64 represents int64 value.
	Int64
	// Uint represents uint value.
	Uint
	// Uint8 represents uint8 value.
	Uint8
	// Uint16 represents uint16 value.
	Uint16
	// Uint32 represents uint32 value.
	Uint32
	// Uint64 represents uint64 value.
	Uint64
	// Uintptr represents uintptr value.
	Uintptr
	// Float32 represents float32 value.
	Float32
	// Float64 represents float64 value.
	Float64
	// Complex64 represents complex64 value.
	Complex64
	// Complex128 represents complex128 value.
	Complex128
	// String represents string value.
	String
	// UnsafePointer represents unsafe pointer value.
	UnsafePointer

	// The overflowTagValue is a exceeding the limit of permissible
	// values for the Tag.
	overflowTagValue Tag = (1 << iota)
)

// Labels associates human-readable headings with type tags.
var Labels = map[string]Tag{
	"nil":             Nil,
	"pointer":         Pointer,
	"array":           Array,
	"array of arrays": ArrayOfArrays,
	"array of slices": ArrayOfSlices,
	"slice":           Slice,
	"slice of slices": SliceOfSlices,
	"slice of arrays": SliceOfArrays,
	"struct":          Struct,
	"map":             Map,
	"func":            Func,
	"chan":            Chan,
	"bool":            Bool,
	"int":             Int,
	"int8":            Int8,
	"int16":           Int16,
	"int32":           Int32,
	"int64":           Int64,
	"uint":            Uint,
	"uint8":           Uint8,
	"uint16":          Uint16,
	"uint32":          Uint32,
	"uint64":          Uint64,
	"uintptr":         Uintptr,
	"float32":         Float32,
	"float64":         Float64,
	"complex64":       Complex64,
	"complex128":      Complex128,
	"string":          String,
	"unsafe pointer":  UnsafePointer,
}

// Has returns true if value contains the specified flag.
func (t *Tag) Has(tag Tag) bool {
	v, _ := t.Contains(tag)
	return v
}

// Is returns true if value is the specified flag.
func (t *Tag) Is(tag Tag) bool {
	v, _ := t.Contains(tag)
	return v && t.IsSingle()
}

// IsEqual returns true if value is equal to the specified flag.
func (t *Tag) IsEqual(tag Tag) bool {
	return *t == tag
}

// IsSingle returns true if value contains single of the available flag.
func (t *Tag) IsSingle() bool {
	return bits.OnesCount(uint(*t)) == 1 &&
		*t <= Tag(overflowTagValue+1)>>1
}

// IsComplex returns true if value contains multiple flags.
func (t *Tag) IsComplex() bool {
	if *t == 0 {
		return false
	}

	return !t.IsSingle()
}

// Contains returns true if value contains the specified flag.
// Returns false and an error if the value is invalid or an
// invalid flag is specified.
func (t *Tag) Contains(flag Tag) (bool, error) {
	if flag == 0 {
		return false, errors.New("flag cannot be zero")
	}

	switch {
	case !flag.IsValid():
		return false, errors.New("incorrect flag value")
	case !t.IsValid():
		return false, errors.New("the object is damaged")
	}

	return *t&flag == flag, nil
}

// IsValid returns true if value contains zero, one or an
// unique sum of valid Tag flags. The zero value is a valid value.
func (t *Tag) IsValid() bool {
	// Check if object is zero, which is a valid value.
	if *t == 0 {
		return true
	}

	copy := *t
	// Iterate over all possible values of the constants and
	// check whether they are part of object.
	for tag := Tag(1); tag < overflowTagValue; tag <<= 1 {
		// If tag is part of the object, remove it from object.
		if copy&tag == tag {
			copy ^= tag
		}
	}

	// Check whether all bits of t were "turned off".
	// If t is zero, it means that all bits were matched values of constants,
	// and therefore t is valid.
	return copy == 0
}

// Set sets the specified flags ignores duplicates.
// The flags that were set previously will be discarded.
// Returns a new value if all is well or old value and an
// error if one or more invalid flags are specified.
func (t *Tag) Set(flags ...Tag) (Tag, error) {
	var r Tag

	for _, flag := range flags {
		if !flag.IsValid() {
			return *t, fmt.Errorf("the %d is invalid flag value", flag)
		}

		if ok, _ := r.Contains(flag); !ok {
			r += Tag(flag)
		}
	}

	*t = r
	return *t, nil
}

// Add adds the specified flags ignores duplicates or flags that value
// already contains. Returns a new value if all is well or old value and
// an error if one or more invalid flags are specified.
func (t *Tag) Add(flags ...Tag) (Tag, error) {
	r := *t

	for _, flag := range flags {
		if !flag.IsValid() {
			return *t, fmt.Errorf("the %d is invalid flag value", flag)
		}

		if ok, _ := r.Contains(flag); !ok {
			r += Tag(flag)
		}
	}

	*t = r
	return *t, nil
}

// Delete deletes the specified flags ignores duplicates or
// flags that were not set. Returns a new value if all is well or
// old value and an error if one or more invalid flags are specified.
func (t *Tag) Delete(flags ...Tag) (Tag, error) {
	r := *t

	for _, flag := range flags {
		if !flag.IsValid() {
			return *t, fmt.Errorf("the %d is invalid flag value", flag)
		}

		if ok, _ := r.Contains(flag); ok {
			r -= Tag(flag)
		}
	}

	*t = r
	return *t, nil
}

// All returns true if all of the specified flags are set.
func (t *Tag) All(flags ...Tag) bool {
	for _, flag := range flags {
		if ok, _ := t.Contains(flag); !ok {
			return false
		}
	}

	return true
}

// Any returns true if at least one of the specified flags is set.
func (t *Tag) Any(flags ...Tag) bool {
	for _, flag := range flags {
		if ok, _ := t.Contains(flag); ok {
			return ok
		}
	}

	return false
}

// IsNil returns true if value contains Nil flag.
func (t *Tag) IsNil() bool {
	v, _ := t.Contains(Nil)
	return v
}

// IsPointer returns true if value contains Pointer flag.
func (t *Tag) IsPointer() bool {
	v, _ := t.Contains(Pointer)
	return v
}

// IsArray returns true if value contains Array flag.
func (t *Tag) IsArray() bool {
	v, _ := t.Contains(Array)
	return v
}

// IsArrayOfArrays returns true if value contains ArrayOfArrays flag.
func (t *Tag) IsArrayOfArrays() bool {
	v, _ := t.Contains(ArrayOfArrays)
	return v
}

// IsArrayOfSlices returns true if value contains ArrayOfSlices flag.
func (t *Tag) IsArrayOfSlices() bool {
	v, _ := t.Contains(ArrayOfSlices)
	return v
}

// IsSlice returns true if value contains Slice flag.
func (t *Tag) IsSlice() bool {
	v, _ := t.Contains(Slice)
	return v
}

// IsSliceOfSlices returns true if value contains SliceOfSlices flag.
func (t *Tag) IsSliceOfSlices() bool {
	v, _ := t.Contains(SliceOfSlices)
	return v
}

// IsSliceOfArrays returns true if value contains SliceOfArrays flag.
func (t *Tag) IsSliceOfArrays() bool {
	v, _ := t.Contains(SliceOfArrays)
	return v
}

// IsStruct returns true if value contains Struct flag.
func (t *Tag) IsStruct() bool {
	v, _ := t.Contains(Struct)
	return v
}

// IsMap returns true if value contains Map flag.
func (t *Tag) IsMap() bool {
	v, _ := t.Contains(Map)
	return v
}

// IsFunction returns true if value contains Function flag.
func (t *Tag) IsFunction() bool {
	v, _ := t.Contains(Func)
	return v
}

// IsChannel returns true if value contains Channel flag.
func (t *Tag) IsChannel() bool {
	v, _ := t.Contains(Chan)
	return v
}

// IsBool returns true if value contains Bool flag.
func (t *Tag) IsBool() bool {
	v, _ := t.Contains(Bool)
	return v
}

// IsInt returns true if value contains Int flag.
func (t *Tag) IsInt() bool {
	v, _ := t.Contains(Int)
	return v
}

// IsInt8 returns true if value contains Int8 flag.
func (t *Tag) IsInt8() bool {
	v, _ := t.Contains(Int8)
	return v
}

// IsInt16 returns true if value contains Int16 flag.
func (t *Tag) IsInt16() bool {
	v, _ := t.Contains(Int16)
	return v
}

// IsInt32 returns true if value contains Int32 flag.
func (t *Tag) IsInt32() bool {
	v, _ := t.Contains(Int32)
	return v
}

// IsInt64 returns true if value contains Int64 flag.
func (t *Tag) IsInt64() bool {
	v, _ := t.Contains(Int64)
	return v
}

// IsUint returns true if value contains Uint flag.
func (t *Tag) IsUint() bool {
	v, _ := t.Contains(Uint)
	return v
}

// IsUint8 returns true if value contains Uint8 flag.
func (t *Tag) IsUint8() bool {
	v, _ := t.Contains(Uint8)
	return v
}

// IsUint16 returns true if value contains Uint16 flag.
func (t *Tag) IsUint16() bool {
	v, _ := t.Contains(Uint16)
	return v
}

// IsUint32 returns true if value contains Uint32 flag.
func (t *Tag) IsUint32() bool {
	v, _ := t.Contains(Uint32)
	return v
}

// IsUint64 returns true if value contains Uint64 flag.
func (t *Tag) IsUint64() bool {
	v, _ := t.Contains(Uint64)
	return v
}

// IsUintptr returns true if value contains Uintptr flag.
func (t *Tag) IsUintptr() bool {
	v, _ := t.Contains(Uintptr)
	return v
}

// IsFloat32 returns true if value contains Float32 flag.
func (t *Tag) IsFloat32() bool {
	v, _ := t.Contains(Float32)
	return v
}

// IsFloat64 returns true if value contains Float64 flag.
func (t *Tag) IsFloat64() bool {
	v, _ := t.Contains(Float64)
	return v
}

// IsComplex64 returns true if value contains Complex64 flag.
func (t *Tag) IsComplex64() bool {
	v, _ := t.Contains(Complex64)
	return v
}

// IsComplex128 returns true if value contains Complex128 flag.
func (t *Tag) IsComplex128() bool {
	v, _ := t.Contains(Complex128)
	return v
}

// IsString returns true if value contains String flag.
func (t *Tag) IsString() bool {
	v, _ := t.Contains(String)
	return v
}

// IsUnsafePointer returns true if value contains UnsafePointer flag.
func (t *Tag) IsUnsafePointer() bool {
	v, _ := t.Contains(UnsafePointer)
	return v
}
