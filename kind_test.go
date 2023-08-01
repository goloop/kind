package kind

import (
	"fmt"
	"testing"

	"github.com/goloop/kind/tag"
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
	appendDiff := func(field string, aValue, bValue bool) {
		if aValue != bValue {
			equal = false
			diffMap[field] = [2]string{
				fmt.Sprintf("%v", aValue),
				fmt.Sprintf("%v", bValue),
			}
		}
	}

	if a.name != b.name {
		equal = false
		diffMap["name"] = [2]string{
			fmt.Sprintf("%v", a.name),
			fmt.Sprintf("%v", b.name),
		}
	}

	appendDiff("isNil", a.IsNil(), b.IsNil())
	appendDiff("isPointer", a.IsPointer(), b.IsPointer())
	appendDiff("isArray", a.IsArray(), b.IsArray())
	appendDiff("isSlice", a.IsSlice(), b.IsSlice())
	appendDiff("isSliceOfSlices", a.IsSliceOfSlices(), b.IsSliceOfSlices())
	appendDiff("isArrayOfSlices", a.IsArrayOfSlices(), b.IsArrayOfSlices())
	appendDiff("isSliceOfArrays", a.IsSliceOfArrays(), b.IsSliceOfArrays())
	appendDiff("isArrayOfArrays", a.IsArrayOfArrays(), b.IsArrayOfArrays())
	appendDiff("isStruct", a.IsStruct(), b.IsStruct())
	appendDiff("isFunction", a.IsFunction(), b.IsFunction())
	appendDiff("isChannel", a.IsChannel(), b.IsChannel())
	appendDiff("isBool", a.IsBool(), b.IsBool())
	appendDiff("isString", a.IsString(), b.IsString())
	appendDiff("isInt8", a.IsInt8(), b.IsInt8())
	appendDiff("isInt16", a.IsInt16(), b.IsInt16())
	appendDiff("isInt32", a.IsInt32(), b.IsInt32())
	appendDiff("isInt64", a.IsInt64(), b.IsInt64())
	appendDiff("isUint8", a.IsUint8(), b.IsUint8())
	appendDiff("isUint16", a.IsUint16(), b.IsUint16())
	appendDiff("isUint32", a.IsUint32(), b.IsUint32())
	appendDiff("isUint64", a.IsUint64(), b.IsUint64())
	appendDiff("isInt", a.IsInt(), b.IsInt())
	appendDiff("isUint", a.IsUint(), b.IsUint())
	appendDiff("isUintptr", a.IsUintptr(), b.IsUintptr())
	appendDiff("isFloat32", a.IsFloat32(), b.IsFloat32())
	appendDiff("isFloat64", a.IsFloat64(), b.IsFloat64())
	appendDiff("isComplex64", a.IsComplex64(), b.IsComplex64())
	appendDiff("isComplex128", a.IsComplex128(), b.IsComplex128())

	// Map has a special case for key and value types.
	appendDiff("isMap", a.IsMap(), b.IsMap())
	mapDiff := ""
	if a.IsMap() && b.IsMap() {

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
			kind:  &Kind{name: "bool", tag: tag.Bool},
		},
		{
			name:  "string",
			input: "test",
			kind:  &Kind{name: "string", tag: tag.String},
		},
		{
			name:  "int8",
			input: int8(1),
			kind:  &Kind{name: "int8", tag: tag.Int8},
		},
		{
			name:  "int16",
			input: int16(1),
			kind:  &Kind{name: "int16", tag: tag.Int16},
		},
		{
			name:  "int32",
			input: int32(1),
			kind:  &Kind{name: "int32", tag: tag.Int32},
		},
		{
			name:  "int64",
			input: int64(1),
			kind:  &Kind{name: "int64", tag: tag.Int64},
		},
		{
			name:  "uint8",
			input: uint8(1),
			kind:  &Kind{name: "uint8", tag: tag.Uint8},
		},
		{
			name:  "uint16",
			input: uint16(1),
			kind:  &Kind{name: "uint16", tag: tag.Uint16},
		},
		{
			name:  "uint32",
			input: uint32(1),
			kind:  &Kind{name: "uint32", tag: tag.Uint32},
		},
		{
			name:  "uint64",
			input: uint64(1),
			kind:  &Kind{name: "uint64", tag: tag.Uint64},
		},
		{
			name:  "int",
			input: 1,
			kind:  &Kind{name: "int", tag: tag.Int},
		},
		{
			name:  "uint",
			input: uint(1),
			kind:  &Kind{name: "uint", tag: tag.Uint},
		},
		{
			name:  "uintptr",
			input: uintptr(1),
			kind:  &Kind{name: "uintptr", tag: tag.Uintptr},
		},
		{
			name:  "float32",
			input: float32(1),
			kind:  &Kind{name: "float32", tag: tag.Float32},
		},
		{
			name:  "float64",
			input: float64(1),
			kind:  &Kind{name: "float64", tag: tag.Float64},
		},
		{
			name:  "complex64",
			input: complex64(1),
			kind:  &Kind{name: "complex64", tag: tag.Complex64},
		},
		{
			name:  "complex128",
			input: complex128(1),
			kind:  &Kind{name: "complex128", tag: tag.Complex128},
		},
		{
			name:  "nil",
			input: nil,
			kind:  &Kind{name: "nil", tag: tag.Nil},
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
				name: "[5]int",
				tag:  tag.Array | tag.Int,
			},
		},
		{
			name:  "pointer",
			input: new(int),
			kind: &Kind{
				name: "*int",
				tag:  tag.Pointer | tag.Int,
			},
		},
		{
			name:  "slice",
			input: []int{1, 2, 3},
			kind: &Kind{
				name: "[]int",
				tag:  tag.Slice | tag.Int,
			},
		},
		{
			name:  "slice of slices",
			input: [][]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name: "[][]int",
				tag:  tag.SliceOfSlices | tag.Int,
			},
		},
		{
			name:  "slice of arrays",
			input: [][2]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name: "[][2]int",
				tag:  tag.SliceOfArrays | tag.Int,
			},
		},
		{
			name:  "array of slices",
			input: [2][]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name: "[2][]int",
				tag:  tag.ArrayOfSlices | tag.Int,
			},
		},
		{
			name:  "array of arrays",
			input: [2][2]int{{1, 2}, {3, 4}},
			kind: &Kind{
				name: "[2][2]int",
				tag:  tag.ArrayOfArrays | tag.Int,
			},
		},
		{
			name:  "map of int",
			input: map[string]int{"one": 1, "two": 2},
			kind: &Kind{
				name:         "map[string]int",
				tag:          tag.Map,
				mapKeyKind:   &Kind{name: "string", tag: tag.String},
				mapValueKind: &Kind{name: "int", tag: tag.Int},
			},
		},
		{
			name:  "map of slice",
			input: map[string][]int{"one": {1}, "two": {1, 2}},
			kind: &Kind{
				name:         "map[string][]int",
				tag:          tag.Map,
				mapKeyKind:   &Kind{name: "string", tag: tag.String},
				mapValueKind: &Kind{name: "[]int", tag: tag.Slice | tag.Int},
			},
		},
		{
			name:  "channel",
			input: make(chan int),
			kind: &Kind{
				name: "chan int",
				tag:  tag.Chan | tag.Int,
			},
		},
		{
			name:  "struct",
			input: struct{ a int }{a: 1},
			kind: &Kind{
				name: "struct { a int }",
				tag:  tag.Struct,
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
