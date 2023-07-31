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
)

// Kind is a struct that represents detailed information about the type of an instance.
type Kind struct {
	name            string      // name of the type
	value           interface{} // original value
	mapKeyKind      *Kind       // representing the key type of a map
	mapValueKind    *Kind       // representing the value type of a map
	isMap           bool        // value is a map type
	isUndefined     bool        // type is undefined
	isNil           bool        // value is nil
	isPointer       bool        // value is a pointer type
	isArray         bool        // value is an array type
	isSlice         bool        // value is a slice type
	isSliceOfSlices bool        // value is a slice of slices ([][]int)
	isArrayOfSlices bool        // value is an array of slices ([5][]int)
	isSliceOfArrays bool        // value is a slice of arrays ([][5]int)
	isArrayOfArrays bool        // value is an array of arrays ([5][5]int)
	isStruct        bool        // value is a struct type
	isInterface     bool        // value is an interface type
	isFunction      bool        // value is a function type
	isChannel       bool        // value is a channel type
	isBool          bool        // value is of bool type
	isString        bool        // value is of string type
	isInt8          bool        // value is of int8 type
	isInt16         bool        // value is of int16 type
	isInt32         bool        // value is of int32 type
	isInt64         bool        // value is of int64 type
	isUint8         bool        // value is of uint8 type
	isUint16        bool        // value is of uint16 type
	isUint32        bool        // value is of uint32 type
	isUint64        bool        // value is of uint64 type
	isInt           bool        // value is of int type
	isUint          bool        // value is of uint type
	isUintptr       bool        // value is of uintptr type
	isFloat32       bool        // value is of float32 type
	isFloat64       bool        // value is of float64 type
	isComplex64     bool        // value is of complex64 type
	isComplex128    bool        // value is of complex128 type
}

// IsComplex returns true if the Kind instance represents a complex type.
// Types like int, uint, string, etc are simple types. That is, they can be
// determined by one indicator, for example, IsInt(), IsUint(), IsString, etc..
// Types like struct, map, slice are complex types, for example:
// map[string]int == IsMap() and IsString() and IsInt() - have several attr.
func (k *Kind) IsComplex() bool {
	fields := []bool{
		k.isUndefined, k.isNil, k.isPointer, k.isArray, k.isSlice,
		k.isSliceOfSlices, k.isArrayOfSlices, k.isSliceOfArrays,
		k.isArrayOfArrays, k.isMap, k.isStruct, k.isInterface,
		k.isFunction, k.isChannel, k.isBool, k.isString,
		k.isInt8, k.isInt16, k.isInt32, k.isInt64,
		k.isUint8, k.isUint16, k.isUint32, k.isUint64,
		k.isInt, k.isUint, k.isUintptr,
		k.isFloat32, k.isFloat64,
		k.isComplex64, k.isComplex128,
	}

	// Count the number of active attributes.
	count := 0
	for _, field := range fields {
		if field {
			count++
		}

		if count > 2 {
			break
		}
	}

	return count > 2
}

// MapKeyKind returns the Kind instance of the map key.
func (k *Kind) MapKeyKind() *Kind {
	if k.isMap && k.mapKeyKind != nil {
		return k.mapKeyKind
	}

	return &Kind{name: "nil", isNil: true}
}

// MapValueKind returns the Kind instance of the map key.
func (k *Kind) MapValueKind() *Kind {
	if k.isMap && k.mapValueKind != nil {
		return k.mapValueKind
	}

	return &Kind{name: "nil", isNil: true}
}

// Name returns the name of the Kind instance.
func (k *Kind) Name() string {
	return k.name
}

// IsUndefined returns true if the Kind instance represents an undefined type.
func (k *Kind) IsUndefined() bool {
	return k.isUndefined
}

// IsNil returns true if the Kind instance represents a nil type.
func (k *Kind) IsNil() bool {
	return k.isNil
}

// IsPointer returns true if the Kind instance represents a pointer type.
func (k *Kind) IsPointer() bool {
	return k.isPointer
}

// IsArray returns true if the Kind instance represents an array type.
func (k *Kind) IsArray() bool {
	return k.isArray
}

// IsSlice returns true if the Kind instance represents a slice type.
func (k *Kind) IsSlice() bool {
	return k.isSlice
}

// IsSliceOfSlices returns true if the Kind instance represents
// a slice of slices type.
func (k *Kind) IsSliceOfSlices() bool {
	return k.isSliceOfSlices
}

// IsArrayOfSlices returns true if the Kind instance represents
// an array of slices type.
func (k *Kind) IsArrayOfSlices() bool {
	return k.isArrayOfSlices
}

// IsSliceOfArrays returns true if the Kind instance represents
// a slice of arrays type.
func (k *Kind) IsSliceOfArrays() bool {
	return k.isSliceOfArrays
}

// IsArrayOfArrays returns true if the Kind instance represents
// an array of arrays type.
func (k *Kind) IsArrayOfArrays() bool {
	return k.isArrayOfArrays
}

// IsMap returns true if the Kind instance represents a map type.
func (k *Kind) IsMap() bool {
	return k.isMap
}

// IsStruct returns true if the Kind instance represents a struct type.
func (k *Kind) IsStruct() bool {
	return k.isStruct
}

// IsInterface returns true if the Kind instance represents an interface type.
func (k *Kind) IsInterface() bool {
	return k.isInterface
}

// IsFunction returns true if the Kind instance represents a function type.
func (k *Kind) IsFunction() bool {
	return k.isFunction
}

// IsChannel returns true if the Kind instance represents a channel type.
func (k *Kind) IsChannel() bool {
	return k.isChannel
}

// IsBool returns true if the Kind instance represents a bool type.
func (k *Kind) IsBool() bool {
	return k.isBool
}

// IsString returns true if the Kind instance represents a string type.
func (k *Kind) IsString() bool {
	return k.isString
}

// IsInt8 returns true if the Kind instance represents an int8 type.
func (k *Kind) IsInt8() bool {
	return k.isInt8
}

// IsInt16 returns true if the Kind instance represents an int16 type.
func (k *Kind) IsInt16() bool {
	return k.isInt16
}

// IsInt32 returns true if the Kind instance represents an int32 type.
func (k *Kind) IsInt32() bool {
	return k.isInt32
}

// IsInt64 returns true if the Kind instance represents an int64 type.
func (k *Kind) IsInt64() bool {
	return k.isInt64
}

// IsUint8 returns true if the Kind instance represents a uint8 type.
func (k *Kind) IsUint8() bool {
	return k.isUint8
}

// IsUint16 returns true if the Kind instance represents a uint16 type.
func (k *Kind) IsUint16() bool {
	return k.isUint16
}

// IsUint32 returns true if the Kind instance represents a uint32 type.
func (k *Kind) IsUint32() bool {
	return k.isUint32
}

// IsUint64 returns true if the Kind instance represents a uint64 type.
func (k *Kind) IsUint64() bool {
	return k.isUint64
}

// IsInt returns true if the Kind instance represents an int type.
func (k *Kind) IsInt() bool {
	return k.isInt
}

// IsUint returns true if the Kind instance represents a uint type.
func (k *Kind) IsUint() bool {
	return k.isUint
}

// IsUintptr returns true if the Kind instance represents a uintptr type.
func (k *Kind) IsUintptr() bool {
	return k.isUintptr
}

// IsFloat32 returns true if the Kind instance represents a float32 type.
func (k *Kind) IsFloat32() bool {
	return k.isFloat32
}

// IsFloat64 returns true if the Kind instance represents a float64 type.
func (k *Kind) IsFloat64() bool {
	return k.isFloat64
}

// IsComplex64 returns true if the Kind instance represents a complex64 type.
func (k *Kind) IsComplex64() bool {
	return k.isComplex64
}

// IsComplex128 returns true if the Kind instance represents a complex128 type.
func (k *Kind) IsComplex128() bool {
	return k.isComplex128
}

// IsNumber returns true if the Kind instance represents a number type.
func (k *Kind) IsNumber() bool {
	return k.IsAnyInt() || k.IsAnyFloat() || k.IsAnyComplex()
}

// IsAnyInt returns true if the Kind instance represents an integer type.
func (k *Kind) IsAnyInt() bool {
	return k.isInt8 || k.isInt16 || k.isInt32 || k.isInt64 || k.isInt ||
		k.isUint8 || k.isUint16 || k.isUint32 || k.isUint64 || k.isUint
}

// IsAnyFloat returns true if the Kind instance represents a float type.
func (k *Kind) IsAnyFloat() bool {
	return k.isFloat32 || k.isFloat64
}

// IsAnyComplex returns true if the Kind instance represents a complex type.
func (k *Kind) IsAnyComplex() bool {
	return k.isComplex64 || k.isComplex128
}

// IsUnsigned returns true if the Kind instance represents an unsigned type.
func (k *Kind) IsUnsigned() bool {
	return k.isUint8 || k.isUint16 || k.isUint32 || k.isUint64 || k.isUint
}

// IsSigned returns true if the Kind instance represents a signed type.
func (k *Kind) IsSigned() bool {
	return k.isInt8 || k.isInt16 || k.isInt32 || k.isInt64 || k.isInt ||
		k.isFloat32 || k.isFloat64 || k.isComplex64 || k.isComplex128
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
	return k.name == strings.Replace(strings.ToLower(name), " ", "", -1)
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
		k.isNil = true
		k.name = "nil"
		return k
	}

	t := reflect.TypeOf(v)
	k.name = t.String()

	checkComplexTypes(k, t)

	return k
}

// checkComplexTypes checks for complex types like slices,
// arrays, pointers, etc.
//
// This function is called recursively.
func checkComplexTypes(k *Kind, t reflect.Type) {
	multiSequence := func() bool {
		return k.isSliceOfSlices || k.isSliceOfArrays ||
			k.isArrayOfSlices || k.isArrayOfArrays
	}

	switch t.Kind() {
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Slice {
			k.isSliceOfSlices = true
		} else if t.Elem().Kind() == reflect.Array {
			k.isSliceOfArrays = true
		} else if !multiSequence() {
			k.isSlice = true
		}

		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem())
	case reflect.Array:
		if t.Elem().Kind() == reflect.Slice {
			k.isArrayOfSlices = true
		} else if t.Elem().Kind() == reflect.Array {
			k.isArrayOfArrays = true
		} else if !multiSequence() {
			k.isArray = true
		}

		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem())
	case reflect.Ptr:
		k.isPointer = true

		// Continue checking for more complex types.
		checkComplexTypes(k, t.Elem())
	case reflect.Map:
		k.isMap = true
		k.mapKeyKind = &Kind{name: t.Key().String()}
		k.mapValueKind = &Kind{name: t.Elem().String()}
		checkComplexTypes(k.mapKeyKind, t.Key())
		checkComplexTypes(k.mapValueKind, t.Elem())

		// fmt.Println("!!!!", k.mapKeyKind, t.Key())
		//k.mapKeyKind = Of(t.Key())
		//k.mapValueKind = Of(t.Elem())
	case reflect.Chan:
		k.isChannel = true
		checkComplexTypes(k, t.Elem())
	case reflect.Struct:
		k.isStruct = true
		// For struct, we stop the recursion,
		// because it could have many different types of fields.
	case reflect.Interface:
		k.isInterface = true
		// For interface, we also stop the recursion,
		// because it could have many different types of methods
	default:
		switch t.Kind() {
		case reflect.Bool:
			k.isBool = true
		case reflect.String:
			k.isString = true
		case reflect.Int8:
			k.isInt8 = true
		case reflect.Int16:
			k.isInt16 = true
		case reflect.Int32:
			k.isInt32 = true
		case reflect.Int64:
			k.isInt64 = true
		case reflect.Uint8:
			k.isUint8 = true
		case reflect.Uint16:
			k.isUint16 = true
		case reflect.Uint32:
			k.isUint32 = true
		case reflect.Uint64:
			k.isUint64 = true
		case reflect.Int:
			k.isInt = true
		case reflect.Uint:
			k.isUint = true
		case reflect.Uintptr:
			k.isUintptr = true
		case reflect.Float32:
			k.isFloat32 = true
		case reflect.Float64:
			k.isFloat64 = true
		case reflect.Complex64:
			k.isComplex64 = true
		case reflect.Complex128:
			k.isComplex128 = true
		default:
			k.name = "undefined"
			k.isUndefined = true
		}
	}
}
