package kind

import (
	"reflect"
	"testing"
	"unsafe"
)

type userID int
type namedString string
type recursiveSlice []recursiveSlice

func TestOfSimpleTypes(t *testing.T) {
	tests := []struct {
		name  string
		input any
		check func(*Kind) bool
	}{
		{"bool", true, (*Kind).IsBool},
		{"string", "test", (*Kind).IsString},
		{"int8", int8(1), (*Kind).IsInt8},
		{"int16", int16(1), (*Kind).IsInt16},
		{"int32", int32(1), (*Kind).IsInt32},
		{"int64", int64(1), (*Kind).IsInt64},
		{"uint8", uint8(1), (*Kind).IsUint8},
		{"uint16", uint16(1), (*Kind).IsUint16},
		{"uint32", uint32(1), (*Kind).IsUint32},
		{"uint64", uint64(1), (*Kind).IsUint64},
		{"int", 1, (*Kind).IsInt},
		{"uint", uint(1), (*Kind).IsUint},
		{"uintptr", uintptr(1), (*Kind).IsUintptr},
		{"float32", float32(1), (*Kind).IsFloat32},
		{"float64", float64(1), (*Kind).IsFloat64},
		{"complex64", complex64(1), (*Kind).IsComplex64},
		{"complex128", complex128(1), (*Kind).IsComplex128},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Of(tt.input)
			if !tt.check(k) {
				t.Fatalf("%s was not classified correctly: %#v", tt.name, k)
			}
			if k.IsNil() || k.IsZero() {
				t.Fatalf("%s unexpectedly nil/zero", tt.name)
			}
		})
	}
}

func TestOfCompositeTypes(t *testing.T) {
	tests := []struct {
		name  string
		input any
		check func(*Kind) bool
		leaf  func(*Kind) bool
	}{
		{"array", [5]int{}, (*Kind).IsArray, (*Kind).IsInt},
		{"pointer", new(int), (*Kind).IsPointer, (*Kind).IsInt},
		{"slice", []int{1, 2, 3}, (*Kind).IsSlice, (*Kind).IsInt},
		{"slice of slices", [][]int{{1}}, (*Kind).IsSliceOfSlices, (*Kind).IsInt},
		{"slice of arrays", [][2]int{{1, 2}}, (*Kind).IsSliceOfArrays, (*Kind).IsInt},
		{"array of slices", [2][]int{{1}}, (*Kind).IsArrayOfSlices, (*Kind).IsInt},
		{"array of arrays", [2][2]int{}, (*Kind).IsArrayOfArrays, (*Kind).IsInt},
		{"map", map[string]int{"one": 1}, (*Kind).IsMap, (*Kind).IsInt},
		{"channel", make(chan int), (*Kind).IsChannel, (*Kind).IsInt},
		{"struct", struct{ A int }{}, (*Kind).IsStruct, nil},
		{"function", func() {}, (*Kind).IsFunction, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Of(tt.input)
			if !tt.check(k) {
				t.Fatalf("%s was not classified correctly: %s", tt.name, k)
			}
			if tt.leaf != nil && !tt.leaf(k) {
				t.Fatalf("%s leaf was not classified correctly: %s", tt.name, k)
			}
			if !k.IsComplex() {
				t.Fatalf("%s should be complex", tt.name)
			}
		})
	}
}

func TestNilAndZero(t *testing.T) {
	var p *int
	var s []int
	var m map[string]int
	var ch chan int
	var fn func()

	for name, input := range map[string]any{
		"nil":     nil,
		"pointer": p,
		"slice":   s,
		"map":     m,
		"channel": ch,
		"func":    fn,
	} {
		t.Run(name, func(t *testing.T) {
			k := Of(input)
			if !k.IsNil() {
				t.Fatalf("expected nil for %s", name)
			}
			if !k.IsZero() {
				t.Fatalf("expected zero for %s", name)
			}
		})
	}

	if Of([]int{}).IsNil() {
		t.Fatal("empty non-nil slice must not be nil")
	}
	if !Of(0).IsZero() {
		t.Fatal("zero int must be zero")
	}
}

func TestMapKeyAndElem(t *testing.T) {
	k := Of(map[string][]int{"one": {1}})
	if !k.IsMap() {
		t.Fatal("expected map")
	}
	if key := k.MapKeyKind(); !key.IsString() || key.Name() != "string" {
		t.Fatalf("unexpected key kind: %s", key)
	}
	if elem := k.MapValueKind(); !elem.IsSlice() || !elem.IsInt() || elem.Name() != "[]int" {
		t.Fatalf("unexpected value kind: %s", elem)
	}
}

func TestOfTypeAndInterface(t *testing.T) {
	type reader interface{ Read([]byte) (int, error) }

	k := TypeOf[reader]()
	if !k.IsInterface() {
		t.Fatalf("expected interface, got %s", k)
	}
	if k.Type().Kind() != reflect.Interface {
		t.Fatalf("unexpected reflect kind: %s", k.Kind())
	}
}

func TestNamedTypesDoNotPanic(t *testing.T) {
	k := Of(userID(42))
	if !k.IsInt() || !k.IsNamed() {
		t.Fatalf("unexpected named int classification: %s", k)
	}
	if got, ok := k.AsInt(); !ok || got != 42 {
		t.Fatalf("AsInt() = %d, %v; want 42, true", got, ok)
	}

	s := Of(namedString("hello"))
	if got, ok := s.AsString(); !ok || got != "hello" {
		t.Fatalf("AsString() = %q, %v; want hello, true", got, ok)
	}
}

func TestUintptrAndUnsafePointer(t *testing.T) {
	k := Of(uintptr(1))
	if !k.IsUintptr() || !k.IsUnsigned() || !k.IsAnyInt() || !k.IsNumber() {
		t.Fatalf("unexpected uintptr classification: %#v", k)
	}
	if got, ok := k.AsUintptr(); !ok || got != 1 {
		t.Fatalf("AsUintptr() = %d, %v; want 1, true", got, ok)
	}

	var p unsafe.Pointer
	u := Of(p)
	if !u.IsUnsafePointer() || !u.IsNil() {
		t.Fatalf("unexpected unsafe pointer classification: %#v", u)
	}
	if got, ok := u.AsUnsafePointer(); !ok || got != nil {
		t.Fatalf("AsUnsafePointer() = %v, %v; want nil, true", got, ok)
	}
}

func TestRecursiveTypesDoNotLoop(t *testing.T) {
	k := Of(recursiveSlice{})
	if !k.IsSliceOfSlices() || !k.IsNamed() {
		t.Fatalf("unexpected recursive classification: %s", k)
	}
}

func TestIsName(t *testing.T) {
	k := Of([]int{})
	if !k.Is("[]int") || !k.Is("[] int") || !k.Is("[]INT") {
		t.Fatal("Is should match compact names")
	}
}

func TestCacheReusesDescriptor(t *testing.T) {
	a := Of(map[string]int{})
	b := Of(map[string]int{"x": 1})
	if a.desc != b.desc {
		t.Fatal("expected descriptor cache reuse")
	}
	if descriptorOf(reflect.TypeFor[int]()) != a.MapValueKind().desc {
		t.Fatal("expected nested descriptor cache reuse")
	}
}
