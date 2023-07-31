package kind

import (
	"fmt"
	"reflect"
	"testing"
)

// deepEqualKind compares two Kind structs and returns true if they are equal.
// If they are not equal, it returns false and a string describing the
// differences.
func deepEqualKind(a, b *Kind, prefixes ...string) (bool, string) {
	// If the prefix is not empty.
	prefix := ""
	if len(prefixes) > 0 {
		prefix = prefixes[0]
	}

	// If object is nil, return false.
	if a == nil && b != nil {
		return false, "nil != not nil"
	} else if a != nil && b == nil {
		return false, "not nil != nil"
	} else if a == nil && b == nil {
		return true, ""
	}

	equal := true
	diffMap := make(map[string][2]string)
	appendDiff := func(field string, aValue, bValue interface{}) {
		if !reflect.DeepEqual(aValue, bValue) {
			equal = false
			diffMap[field] = [2]string{
				fmt.Sprintf("%v", aValue),
				fmt.Sprintf("%v", bValue),
			}
		}
	}

	appendDiff("name", a.name, b.name)
	appendDiff("isUndefined", a.isUndefined, b.isUndefined)
	appendDiff("isNil", a.isNil, b.isNil)
	appendDiff("isPointer", a.isPointer, b.isPointer)
	appendDiff("isArray", a.isArray, b.isArray)
	appendDiff("isSlice", a.isSlice, b.isSlice)
	appendDiff("isSliceOfSlices", a.isSliceOfSlices, b.isSliceOfSlices)
	appendDiff("isArrayOfSlices", a.isArrayOfSlices, b.isArrayOfSlices)
	appendDiff("isSliceOfArrays", a.isSliceOfArrays, b.isSliceOfArrays)
	appendDiff("isArrayOfArrays", a.isArrayOfArrays, b.isArrayOfArrays)
	appendDiff("isStruct", a.isStruct, b.isStruct)
	appendDiff("isInterface", a.isInterface, b.isInterface)
	appendDiff("isFunction", a.isFunction, b.isFunction)
	appendDiff("isChannel", a.isChannel, b.isChannel)
	appendDiff("isBool", a.isBool, b.isBool)
	appendDiff("isString", a.isString, b.isString)
	appendDiff("isInt8", a.isInt8, b.isInt8)
	appendDiff("isInt16", a.isInt16, b.isInt16)
	appendDiff("isInt32", a.isInt32, b.isInt32)
	appendDiff("isInt64", a.isInt64, b.isInt64)
	appendDiff("isUint8", a.isUint8, b.isUint8)
	appendDiff("isUint16", a.isUint16, b.isUint16)
	appendDiff("isUint32", a.isUint32, b.isUint32)
	appendDiff("isUint64", a.isUint64, b.isUint64)
	appendDiff("isInt", a.isInt, b.isInt)
	appendDiff("isUint", a.isUint, b.isUint)
	appendDiff("isUintptr", a.isUintptr, b.isUintptr)
	appendDiff("isFloat32", a.isFloat32, b.isFloat32)
	appendDiff("isFloat64", a.isFloat64, b.isFloat64)
	appendDiff("isComplex64", a.isComplex64, b.isComplex64)
	appendDiff("isComplex128", a.isComplex128, b.isComplex128)

	// Map has a special case for key and value types.
	appendDiff("isMap", a.isMap, b.isMap)
	mapDiff := ""
	if a.isMap && b.isMap {

		if ok, r := deepEqualKind(a.mapKeyKind, b.mapKeyKind, "\t"); !ok {
			mapDiff += fmt.Sprintf("mapKeyKind:\n%s", r)
		}

		if ok, r := deepEqualKind(a.mapValueKind, b.mapValueKind, "\t"); !ok {
			mapDiff += fmt.Sprintf("mapValueKind:\n%s", r)
		}

		if mapDiff != "" {
			equal = false
			diffMap["isMap"] = [2]string{"true", "true"}
		}
	}

	// Generate the diff string.
	result := ""
	for k, v := range diffMap {
		result += fmt.Sprintf("%s%s: %s != %s\n", prefix, k, v[0], v[1])
		if k == "isMap" {
			result += mapDiff
		}
	}

	return equal, result
}

// TestOfSimpleTypes tests the kind.Of function for simple types.
func TestOfSimpleTypes(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		kind  *Kind
	}{
		{
			name:  "bool",
			input: true,
			kind:  &Kind{name: "bool", isBool: true},
		},
		{
			name:  "string",
			input: "test",
			kind:  &Kind{name: "string", isString: true},
		},
		{
			name:  "int8",
			input: int8(1),
			kind:  &Kind{name: "int8", isInt8: true},
		},
		{
			name:  "int16",
			input: int16(1),
			kind:  &Kind{name: "int16", isInt16: true},
		},
		{
			name:  "int32",
			input: int32(1),
			kind:  &Kind{name: "int32", isInt32: true},
		},
		{
			name:  "int64",
			input: int64(1),
			kind:  &Kind{name: "int64", isInt64: true},
		},
		{
			name:  "uint8",
			input: uint8(1),
			kind:  &Kind{name: "uint8", isUint8: true},
		},
		{
			name:  "uint16",
			input: uint16(1),
			kind:  &Kind{name: "uint16", isUint16: true},
		},
		{
			name:  "uint32",
			input: uint32(1),
			kind:  &Kind{name: "uint32", isUint32: true},
		},
		{
			name:  "uint64",
			input: uint64(1),
			kind:  &Kind{name: "uint64", isUint64: true},
		},
		{
			name:  "int",
			input: 1,
			kind:  &Kind{name: "int", isInt: true},
		},
		{
			name:  "uint",
			input: uint(1),
			kind:  &Kind{name: "uint", isUint: true},
		},
		{
			name:  "uintptr",
			input: uintptr(1),
			kind:  &Kind{name: "uintptr", isUintptr: true},
		},
		{
			name:  "float32",
			input: float32(1),
			kind:  &Kind{name: "float32", isFloat32: true},
		},
		{
			name:  "float64",
			input: float64(1),
			kind:  &Kind{name: "float64", isFloat64: true},
		},
		{
			name:  "complex64",
			input: complex64(1),
			kind:  &Kind{name: "complex64", isComplex64: true},
		},
		{
			name:  "complex128",
			input: complex128(1),
			kind:  &Kind{name: "complex128", isComplex128: true},
		},
		{
			name:  "nil",
			input: nil,
			kind:  &Kind{name: "nil", isNil: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kind := Of(tt.input)
			if kind.Name() != tt.kind.Name() {
				t.Errorf("Expected type name %s, but got %s",
					tt.kind.Name(), kind.Name())
			}

			if ok, result := deepEqualKind(tt.kind, kind); !ok {
				t.Errorf("DeepEqual expected kind %+v, but got %+v:\n%s",
					tt.kind, kind, result)
			}
		})
	}
}

// TestOfComplexTypes tests the kind.Of function for complex types.
func TestOfComplexTypes(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		kind  *Kind
	}{
		{
			name:  "array",
			input: [5]int{1, 2, 3, 4, 5},
			kind: &Kind{
				name:    "[5]int",
				isArray: true,
				isInt:   true,
			},
		},
		{
			name:  "pointer",
			input: new(int),
			kind: &Kind{
				name:      "*int",
				isPointer: true,
				isInt:     true,
			},
		},
		{
			name:  "slice",
			input: []int{1, 2, 3},
			kind: &Kind{
				name:    "[]int",
				isSlice: true,
				isInt:   true,
			},
		},
		{
			name:  "slice of slices",
			input: [][]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name:            "[][]int",
				isSlice:         false, // because it is slice of slices
				isSliceOfSlices: true,
				isInt:           true,
			},
		},
		{
			name:  "slice of arrays",
			input: [][2]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name:            "[][2]int",
				isSlice:         false, // because it is slice of arrays
				isSliceOfArrays: true,
				isInt:           true,
			},
		},
		{
			name:  "array of slices",
			input: [2][]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name:            "[2][]int",
				isArray:         false, // because it is slice of arrays
				isArrayOfSlices: true,
				isInt:           true,
			},
		},
		{
			name:  "array of arrays",
			input: [2][2]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name:            "[2][2]int",
				isArray:         false, // because it is slice of arrays
				isArrayOfArrays: true,
				isInt:           true,
			},
		},
		{
			name:  "map of int",
			input: map[string]int{"one": 1, "two": 2},
			kind: &Kind{
				name:         "map[string]int",
				isMap:        true,
				mapKeyKind:   &Kind{name: "string", isString: true},
				mapValueKind: &Kind{name: "int", isInt: true},
			},
		},
		{
			name:  "map of slice",
			input: map[string][]int{"one": {1}, "two": {1, 2}},
			kind: &Kind{
				name:         "map[string][]int",
				isMap:        true,
				mapKeyKind:   &Kind{name: "string", isString: true},
				mapValueKind: &Kind{name: "[]int", isInt: true, isSlice: true},
			},
		},
		{
			name:  "channel",
			input: make(chan int),
			kind: &Kind{
				name:      "chan int",
				isChannel: true,
				isInt:     true,
			},
		},
		{
			name:  "struct",
			input: struct{ a int }{a: 1},
			kind: &Kind{
				name:     "struct { a int }",
				isStruct: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kind := Of(tt.input)
			if kind.Name() != tt.kind.Name() {
				t.Errorf("Expected type name %s, but got %s",
					tt.kind.Name(), kind.Name())
			}

			if ok, result := deepEqualKind(tt.kind, kind); !ok {
				t.Errorf("DeepEqual expected kind %+v, but got %+v:\n%s",
					tt.kind, kind, result)
			}
		})
	}
}
