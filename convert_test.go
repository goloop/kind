package kind

import (
	"reflect"
	"testing"
	"unsafe"
)

// Named types exercise the convert path (defined type -> builtin) for each
// As* accessor, not just the plain builtin kinds.
type (
	myBool    bool
	myString  string
	myInt     int
	myInt8    int8
	myInt16   int16
	myInt32   int32
	myInt64   int64
	myUint    uint
	myUint8   uint8
	myUint16  uint16
	myUint32  uint32
	myUint64  uint64
	myUintptr uintptr
	myF32     float32
	myF64     float64
	myC64     complex64
	myC128    complex128
)

func TestAsAccessorsSuccess(t *testing.T) {
	if v, ok := Of(myBool(true)).AsBool(); !ok || v != true {
		t.Errorf("AsBool = %v %v", v, ok)
	}
	if v, ok := Of(myString("hi")).AsString(); !ok || v != "hi" {
		t.Errorf("AsString = %v %v", v, ok)
	}
	if v, ok := Of(myInt(1)).AsInt(); !ok || v != 1 {
		t.Errorf("AsInt = %v %v", v, ok)
	}
	if v, ok := Of(myInt8(2)).AsInt8(); !ok || v != 2 {
		t.Errorf("AsInt8 = %v %v", v, ok)
	}
	if v, ok := Of(myInt16(3)).AsInt16(); !ok || v != 3 {
		t.Errorf("AsInt16 = %v %v", v, ok)
	}
	if v, ok := Of(myInt32(4)).AsInt32(); !ok || v != 4 {
		t.Errorf("AsInt32 = %v %v", v, ok)
	}
	if v, ok := Of(myInt64(5)).AsInt64(); !ok || v != 5 {
		t.Errorf("AsInt64 = %v %v", v, ok)
	}
	if v, ok := Of(myUint(6)).AsUint(); !ok || v != 6 {
		t.Errorf("AsUint = %v %v", v, ok)
	}
	if v, ok := Of(myUint8(7)).AsUint8(); !ok || v != 7 {
		t.Errorf("AsUint8 = %v %v", v, ok)
	}
	if v, ok := Of(myUint16(8)).AsUint16(); !ok || v != 8 {
		t.Errorf("AsUint16 = %v %v", v, ok)
	}
	if v, ok := Of(myUint32(9)).AsUint32(); !ok || v != 9 {
		t.Errorf("AsUint32 = %v %v", v, ok)
	}
	if v, ok := Of(myUint64(10)).AsUint64(); !ok || v != 10 {
		t.Errorf("AsUint64 = %v %v", v, ok)
	}
	if v, ok := Of(myUintptr(11)).AsUintptr(); !ok || v != 11 {
		t.Errorf("AsUintptr = %v %v", v, ok)
	}
	if v, ok := Of(myF32(1.5)).AsFloat32(); !ok || v != 1.5 {
		t.Errorf("AsFloat32 = %v %v", v, ok)
	}
	if v, ok := Of(myF64(2.5)).AsFloat64(); !ok || v != 2.5 {
		t.Errorf("AsFloat64 = %v %v", v, ok)
	}
	if v, ok := Of(myC64(1 + 2i)).AsComplex64(); !ok || v != 1+2i {
		t.Errorf("AsComplex64 = %v %v", v, ok)
	}
	if v, ok := Of(myC128(3 + 4i)).AsComplex128(); !ok || v != 3+4i {
		t.Errorf("AsComplex128 = %v %v", v, ok)
	}

	x := 42
	if v, ok := Of(unsafe.Pointer(&x)).AsUnsafePointer(); !ok || v != unsafe.Pointer(&x) {
		t.Errorf("AsUnsafePointer = %v %v", v, ok)
	}
}

// A mismatched dynamic kind must return the zero value and false, and a nil
// Kind must never panic.
func TestAsAccessorsMismatchAndNil(t *testing.T) {
	s := Of("not a number")
	if _, ok := s.AsInt(); ok {
		t.Error("AsInt on string returned ok")
	}
	if _, ok := s.AsFloat64(); ok {
		t.Error("AsFloat64 on string returned ok")
	}
	if _, ok := s.AsBool(); ok {
		t.Error("AsBool on string returned ok")
	}

	var nilK *Kind
	if _, ok := nilK.AsInt(); ok {
		t.Error("AsInt on nil Kind returned ok")
	}
	if _, ok := nilK.AsString(); ok {
		t.Error("AsString on nil Kind returned ok")
	}
}

// NOTE-01: IsEmpty does not dereference pointers - a non-nil pointer to a zero
// value is not empty.
func TestIsEmptyDoesNotDerefPointer(t *testing.T) {
	type box struct{ n int }
	if Of(&box{}).IsEmpty() {
		t.Error("IsEmpty(&box{}) = true, want false (pointers are not dereferenced)")
	}
	if !Of((*box)(nil)).IsEmpty() {
		t.Error("IsEmpty((*box)(nil)) = false, want true")
	}
}

// NOTE-02: scalar predicates are leaf-aware; Kind gives the strict check.
func TestLeafAwareVsStrictKind(t *testing.T) {
	k := Of([]int{1, 2, 3})
	if !k.IsInt() {
		t.Error("IsInt([]int) = false, want true (leaf-aware)")
	}
	if k.Kind() == reflect.Int {
		t.Error("Kind([]int) == reflect.Int, want reflect.Slice (strict)")
	}
	if k.Kind() != reflect.Slice {
		t.Errorf("Kind([]int) = %v, want reflect.Slice", k.Kind())
	}
}

// Cheap coverage for accessors and predicates with no prior direct test.
func TestMiscAccessors(t *testing.T) {
	k := Of(123)
	if k.String() != "int" {
		t.Errorf("String = %q", k.String())
	}
	if v, ok := k.Value().(int); !ok || v != 123 {
		t.Errorf("Value = %v %v", v, ok)
	}
	if k.IsUndefined() {
		t.Error("IsUndefined(int) = true")
	}
	if k.IsNilable() {
		t.Error("IsNilable(int) = true")
	}
	if !Of(1 + 2i).IsAnyComplex() {
		t.Error("IsAnyComplex(complex128) = false")
	}
	// Exercise the binary-marshaler capability checks (result type-dependent).
	_ = k.IsBinaryMarshaler()
	_ = k.IsBinaryUnmarshaler()
}
