// Package kind provides utilities to analyze and represent the type
// details of Go values.
//
// The Kind struct represents the type details of an instance, including
// simple types like int, uint, string, bool, etc., as well as complex
// types like slices, maps, structs, arrays, interfaces, and pointers.
//
// The package exposes the Of function that takes a value as input and
// returns a Kind instance representing the type of that value.
// The Kind struct contains various boolean flags and methods
// to query the type details.
//
// It also provides a function, deepEqualKind, to compare two Kind instances
// for equality, allowing users to check if two values have the same type and
// value representation.
//
// Example usage:
//
//	kind := kind.Of(42)
//	fmt.Println(kind.IsInt(), kind.Name()) // true "int"
//
//	kind := kind.Of([]int{1, 2, 3})
//	fmt.Println(kind.IsSlice(), kind.IsInt(), kind.Name()) // true true "[]int"
//
// The Kind struct has methods like IsComplex, IsSigned, IsUnsigned, IsNil,
// IsPointer, etc., to query various attributes of the type. It also provides
// As<Type>() methods to extract the underlying value as a specific Go type
// (e.g., AsBool, AsString, AsInt, AsFloat32, etc.).
//
// Additionally, Kind provides MapKeyKind and MapValueKind methods to get
// the Kind instances of the key and value types of a map, respectively.
//
// The package is useful for situations where type introspection or reflection
// is needed to understand the nature of Go values, especially in cases where
// nested or complex types are involved.
//
// Note: The package assumes that the value passed to the Of function is
// a valid Go value, and it does not handle all possible types. For instance,
// certain edge cases like unsafe pointers, function pointers, and unexported
// fields in structs might not be fully supported. Users should exercise
// caution and test thoroughly in their specific use cases.
//
// This package is written in pure Go and does not have
// any external dependencies.
package kind

import (
	"reflect"
	"strings"

	"github.com/goloop/kind/tag"
)

// Kind is a struct that represents detailed information about the type of an instance.
type Kind struct {
	name         string      // name of the type
	value        interface{} // original value
	mapKeyKind   *Kind       // representing the key type of a map
	mapValueKind *Kind       // representing the value type of a map
	tag          tag.Tag     // tag of the type
}

// IsComplex returns true if the Kind instance represents a complex type.
// Types like int, uint, string, etc are simple types. That is, they can be
// determined by one indicator, for example, IsInt(), IsUint(), IsString, etc..
// Types like struct, map, slice are complex types, for example:
// map[string]int == IsMap() and IsString() and IsInt() - have several attr.
func (k *Kind) IsComplex() bool {
	return k.tag.IsComplex()
}

// MapKeyKind returns the Kind instance of the map key.
func (k *Kind) MapKeyKind() *Kind {
	if k.tag.Has(tag.Map) && k.mapKeyKind != nil {
		return k.mapKeyKind
	}

	return &Kind{name: "nil", tag: tag.Nil}
}

// MapValueKind returns the Kind instance of the map key.
func (k *Kind) MapValueKind() *Kind {
	if k.tag.Has(tag.Map) && k.mapValueKind != nil {
		return k.mapValueKind
	}

	return &Kind{name: "nil", tag: tag.Nil}
}

// Name returns the name of the Kind instance.
func (k *Kind) Name() string {
	return k.name
}

// IsNil returns true if the Kind instance represents a nil type.
func (k *Kind) IsNil() bool {
	return k.tag.Has(tag.Nil)
}

// IsPointer returns true if the Kind instance represents a pointer type.
func (k *Kind) IsPointer() bool {
	return k.tag.Has(tag.Pointer)
}

// IsArray returns true if the Kind instance represents an array type.
func (k *Kind) IsArray() bool {
	return k.tag.Has(tag.Array)
}

// IsSlice returns true if the Kind instance represents a slice type.
func (k *Kind) IsSlice() bool {
	return k.tag.Has(tag.Slice)
}

// IsSliceOfSlices returns true if the Kind instance represents
// a slice of slices type.
func (k *Kind) IsSliceOfSlices() bool {
	return k.tag.Has(tag.SliceOfSlices)
}

// IsArrayOfSlices returns true if the Kind instance represents
// an array of slices type.
func (k *Kind) IsArrayOfSlices() bool {
	return k.tag.Has(tag.ArrayOfSlices)
}

// IsSliceOfArrays returns true if the Kind instance represents
// a slice of arrays type.
func (k *Kind) IsSliceOfArrays() bool {
	return k.tag.Has(tag.SliceOfArrays)
}

// IsArrayOfArrays returns true if the Kind instance represents
// an array of arrays type.
func (k *Kind) IsArrayOfArrays() bool {
	return k.tag.Has(tag.ArrayOfArrays)
}

// IsMap returns true if the Kind instance represents a map type.
func (k *Kind) IsMap() bool {
	return k.tag.Has(tag.Map)
}

// IsStruct returns true if the Kind instance represents a struct type.
func (k *Kind) IsStruct() bool {
	return k.tag.Has(tag.Struct)
}

// IsFunction returns true if the Kind instance represents a function type.
func (k *Kind) IsFunction() bool {
	return k.tag.Has(tag.Func)
}

// IsChannel returns true if the Kind instance represents a channel type.
func (k *Kind) IsChannel() bool {
	return k.tag.Has(tag.Chan)
}

// IsBool returns true if the Kind instance represents a bool type.
func (k *Kind) IsBool() bool {
	return k.tag.Has(tag.Bool)
}

// IsString returns true if the Kind instance represents a string type.
func (k *Kind) IsString() bool {
	return k.tag.Has(tag.String)
}

// IsInt8 returns true if the Kind instance represents an int8 type.
func (k *Kind) IsInt8() bool {
	return k.tag.Has(tag.Int8)
}

// IsInt16 returns true if the Kind instance represents an int16 type.
func (k *Kind) IsInt16() bool {
	return k.tag.Has(tag.Int16)
}

// IsInt32 returns true if the Kind instance represents an int32 type.
func (k *Kind) IsInt32() bool {
	return k.tag.Has(tag.Int32)
}

// IsInt64 returns true if the Kind instance represents an int64 type.
func (k *Kind) IsInt64() bool {
	return k.tag.Has(tag.Int64)
}

// IsUint8 returns true if the Kind instance represents a uint8 type.
func (k *Kind) IsUint8() bool {
	return k.tag.Has(tag.Uint8)
}

// IsUint16 returns true if the Kind instance represents a uint16 type.
func (k *Kind) IsUint16() bool {
	return k.tag.Has(tag.Uint16)
}

// IsUint32 returns true if the Kind instance represents a uint32 type.
func (k *Kind) IsUint32() bool {
	return k.tag.Has(tag.Uint32)
}

// IsUint64 returns true if the Kind instance represents a uint64 type.
func (k *Kind) IsUint64() bool {
	return k.tag.Has(tag.Uint64)
}

// IsInt returns true if the Kind instance represents an int type.
func (k *Kind) IsInt() bool {
	return k.tag.Has(tag.Int)
}

// IsUint returns true if the Kind instance represents a uint type.
func (k *Kind) IsUint() bool {
	return k.tag.Has(tag.Uint)
}

// IsUintptr returns true if the Kind instance represents a uintptr type.
func (k *Kind) IsUintptr() bool {
	return k.tag.Has(tag.Uintptr)
}

// IsFloat32 returns true if the Kind instance represents a float32 type.
func (k *Kind) IsFloat32() bool {
	return k.tag.Has(tag.Float32)
}

// IsFloat64 returns true if the Kind instance represents a float64 type.
func (k *Kind) IsFloat64() bool {
	return k.tag.Has(tag.Float64)
}

// IsComplex64 returns true if the Kind instance represents a complex64 type.
func (k *Kind) IsComplex64() bool {
	return k.tag.Has(tag.Complex64)
}

// IsComplex128 returns true if the Kind instance represents a complex128 type.
func (k *Kind) IsComplex128() bool {
	return k.tag.Has(tag.Complex128)
}

// IsNumber returns true if the Kind instance represents a number type.
func (k *Kind) IsNumber() bool {
	return k.IsAnyInt() || k.IsAnyFloat() || k.IsAnyComplex()
}

// IsAnyInt returns true if the Kind instance represents an integer type.
func (k *Kind) IsAnyInt() bool {
	return k.tag.Any(tag.Int8, tag.Int16, tag.Int32, tag.Int64, tag.Int,
		tag.Uint8, tag.Uint16, tag.Uint32, tag.Uint64, tag.Uint)
}

// IsAnyFloat returns true if the Kind instance represents a float type.
func (k *Kind) IsAnyFloat() bool {
	return k.tag.Any(tag.Float32, tag.Float64)
}

// IsAnyComplex returns true if the Kind instance represents a complex type.
func (k *Kind) IsAnyComplex() bool {
	return k.tag.Any(tag.Complex64, tag.Complex128)
}

// IsUnsigned returns true if the Kind instance represents an unsigned type.
func (k *Kind) IsUnsigned() bool {
	return k.tag.Any(tag.Uint8, tag.Uint16, tag.Uint32, tag.Uint64, tag.Uint)
}

// IsSigned returns true if the Kind instance represents a signed type.
func (k *Kind) IsSigned() bool {
	return k.tag.Any(tag.Int8, tag.Int16, tag.Int32, tag.Int64, tag.Int,
		tag.Float32, tag.Float64, tag.Complex64, tag.Complex128)
}

// Is returns true if the name of the Kind instance is equal to the given name.
//
// Example usage:
//
//	kind := kind.Of(42)
//	fmt.Println(kind.Is("int")) // true
//
//	kind := kind.Of([]int{1, 2, 3})
//	fmt.Println(kind.Is("[]int")) // true
func (k *Kind) Is(name string) bool {
	// Try to find the name in the tag labels.
	name = strings.Replace(strings.ToLower(name), " ", "", -1)
	if t, ok := tag.Labels[name]; ok {
		return k.tag.Has(t)
	}

	return k.name == name
}

// String returns the name of the Kind instance.
func (k *Kind) String() string {
	return k.name
}

// AsBool returns the value of the Kind as bool.
func (k *Kind) AsBool() (bool, bool) {
	if k.IsBool() {
		return k.value.(bool), true
	}

	return false, false
}

// AsString returns the value of the Kind as string.
func (k *Kind) AsString() (string, bool) {
	if k.IsString() {
		return k.value.(string), true
	}

	return "", false
}

// AsInt8 returns the value of the Kind as int8.
func (k *Kind) AsInt8() (int8, bool) {
	if k.IsInt8() {
		return k.value.(int8), true
	}

	return 0, false
}

// AsInt16 returns the value of the Kind as int16.
func (k *Kind) AsInt16() (int16, bool) {
	if k.IsInt16() {
		return k.value.(int16), true
	}

	return 0, false
}

// AsInt32 returns the value of the Kind as int32.
func (k *Kind) AsInt32() (int32, bool) {
	if k.IsInt32() {
		return k.value.(int32), true
	}

	return 0, false
}

// AsInt64 returns the value of the Kind as int64.
func (k *Kind) AsInt64() (int64, bool) {
	if k.IsInt64() {
		return k.value.(int64), true
	}

	return 0, false
}

// AsInt returns the value of the Kind as int.
func (k *Kind) AsInt() (int, bool) {
	if k.IsInt() {
		return k.value.(int), true
	}

	return 0, false
}

// AsUint8 returns the value of the Kind as uint8.
func (k *Kind) AsUint8() (uint8, bool) {
	if k.IsUint8() {
		return k.value.(uint8), true
	}

	return 0, false
}

// AsUint16 returns the value of the Kind as uint16.
func (k *Kind) AsUint16() (uint16, bool) {
	if k.IsUint16() {
		return k.value.(uint16), true
	}

	return 0, false
}

// AsUint32 returns the value of the Kind as uint32.
func (k *Kind) AsUint32() (uint32, bool) {
	if k.IsUint32() {
		return k.value.(uint32), true
	}

	return 0, false
}

// AsUint64 returns the value of the Kind as uint64.
func (k *Kind) AsUint64() (uint64, bool) {
	if k.IsUint64() {
		return k.value.(uint64), true
	}

	return 0, false
}

// AsUint returns the value of the Kind as uint.
func (k *Kind) AsUint() (uint, bool) {
	if k.IsUint() {
		return k.value.(uint), true
	}

	return 0, false
}

// AsFloat32 returns the value of the Kind as float32.
func (k *Kind) AsFloat32() (float32, bool) {
	if k.IsFloat32() {
		return k.value.(float32), true
	}

	return 0, false
}

// AsFloat64 returns the value of the Kind as float64.
func (k *Kind) AsFloat64() (float64, bool) {
	if k.IsFloat64() {
		return k.value.(float64), true
	}

	return 0, false
}

// AsComplex64 returns the value of the Kind as complex64.
func (k *Kind) AsComplex64() (complex64, bool) {
	if k.IsComplex64() {
		return k.value.(complex64), true
	}

	return 0, false
}

// AsComplex128 returns the value of the Kind as complex128.
func (k *Kind) AsComplex128() (complex128, bool) {
	if k.IsComplex128() {
		return k.value.(complex128), true
	}

	return 0, false
}

// Of returns a Kind instance that represents the type of the given value.
//
// Example usage:
//
//	kind := kind.Of(42)
//	fmt.Println(kind.IsInt(), kind.Name()) // true "int"
//
//	kind := kind.Of([]int{1, 2, 3})
//	fmt.Println(kind.IsSlice(), kind.IsInt(), kind.Name()) // true true "[]int"
func Of(v interface{}) *Kind {
	k := new(Kind)
	k.value = v

	if v == nil {
		k.tag.Set(tag.Nil)
		k.name = "nil"

		return k
	}

	t := reflect.TypeOf(v)
	k.name = t.String()

	level := 0
	checkComplexTypes(k, t, level)

	return k
}

// checkComplexTypes checks for complex types like slices,
// arrays, pointers, etc.
//
// This function is called recursively.
func checkComplexTypes(k *Kind, t reflect.Type, level int) {
	multiSequence := func() bool {
		return k.tag.Any(tag.SliceOfSlices, tag.SliceOfArrays,
			tag.ArrayOfSlices, tag.ArrayOfArrays)
	}

	switch t.Kind() {
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Slice {
			k.tag.Add(tag.SliceOfSlices)
		} else if t.Elem().Kind() == reflect.Array {
			k.tag.Add(tag.SliceOfArrays)
		} else if !multiSequence() {
			k.tag.Add(tag.Slice)
		}

		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem(), level+1)
	case reflect.Array:
		if t.Elem().Kind() == reflect.Slice {
			k.tag.Add(tag.ArrayOfSlices)
		} else if t.Elem().Kind() == reflect.Array {
			k.tag.Add(tag.ArrayOfArrays)
		} else if !multiSequence() {
			k.tag.Add(tag.Array)
		}

		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem(), level+1)
	case reflect.Ptr:
		k.tag.Add(tag.Pointer)
		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem(), level+1)
	case reflect.Map:
		k.tag.Add(tag.Map)
		k.mapKeyKind = &Kind{name: t.Key().String()}
		k.mapValueKind = &Kind{name: t.Elem().String()}
		checkComplexTypes(k.mapKeyKind, t.Key(), 0)    // another level
		checkComplexTypes(k.mapValueKind, t.Elem(), 0) // another level
	case reflect.Chan:
		k.tag.Add(tag.Chan)
		checkComplexTypes(k, t.Elem(), level+1)
	case reflect.Struct:
		k.tag.Add(tag.Struct)
		// For struct, we stop the recursion,
		// because it could have many different types of fields.

	// We cannot define an interface because an empty interface is nil,
	// and if an object is passed through an interface, its type is a struct.
	//  case reflect.Interface:
	//  	k.isInterface = true
	//  	// For interface, we also stop the recursion,
	//  	// because it could have many different types of methods.
	default:
		switch t.Kind() {
		case reflect.Bool:
			k.tag.Add(tag.Bool)
		case reflect.String:
			k.tag.Add(tag.String)
		case reflect.Int8:
			k.tag.Add(tag.Int8)
		case reflect.Int16:
			k.tag.Add(tag.Int16)
		case reflect.Int32:
			k.tag.Add(tag.Int32)
		case reflect.Int64:
			k.tag.Add(tag.Int64)
		case reflect.Uint8:
			k.tag.Add(tag.Uint8)
		case reflect.Uint16:
			k.tag.Add(tag.Uint16)
		case reflect.Uint32:
			k.tag.Add(tag.Uint32)
		case reflect.Uint64:
			k.tag.Add(tag.Uint64)
		case reflect.Int:
			k.tag.Add(tag.Int)
		case reflect.Uint:
			k.tag.Add(tag.Uint)
		case reflect.Uintptr:
			k.tag.Add(tag.Uintptr)
		case reflect.Float32:
			k.tag.Add(tag.Float32)
		case reflect.Float64:
			k.tag.Add(tag.Float64)
		case reflect.Complex64:
			k.tag.Add(tag.Complex64)
		case reflect.Complex128:
			k.tag.Add(tag.Complex128)

			// We cannot have a default solution since the object is either
			// a simple type or a struct type. This block will never be used.
			//  default:
			//  	k.name = "undefined"
			//  	k.isUndefined = true
		}
	}
}
