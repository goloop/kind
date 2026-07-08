package kind

import (
	"reflect"
	"unsafe"
)

func (k *Kind) convert(to reflect.Kind) (reflect.Value, bool) {
	if k == nil || k.value == nil {
		return reflect.Value{}, false
	}
	v := reflect.ValueOf(k.value)
	if v.Kind() != to {
		return reflect.Value{}, false
	}
	target := builtinType(to)
	if target == nil || !v.Type().ConvertibleTo(target) {
		return reflect.Value{}, false
	}
	return v.Convert(target), true
}

func builtinType(k reflect.Kind) reflect.Type {
	switch k {
	case reflect.Bool:
		return reflect.TypeFor[bool]()
	case reflect.String:
		return reflect.TypeFor[string]()
	case reflect.Int:
		return reflect.TypeFor[int]()
	case reflect.Int8:
		return reflect.TypeFor[int8]()
	case reflect.Int16:
		return reflect.TypeFor[int16]()
	case reflect.Int32:
		return reflect.TypeFor[int32]()
	case reflect.Int64:
		return reflect.TypeFor[int64]()
	case reflect.Uint:
		return reflect.TypeFor[uint]()
	case reflect.Uint8:
		return reflect.TypeFor[uint8]()
	case reflect.Uint16:
		return reflect.TypeFor[uint16]()
	case reflect.Uint32:
		return reflect.TypeFor[uint32]()
	case reflect.Uint64:
		return reflect.TypeFor[uint64]()
	case reflect.Uintptr:
		return reflect.TypeFor[uintptr]()
	case reflect.Float32:
		return reflect.TypeFor[float32]()
	case reflect.Float64:
		return reflect.TypeFor[float64]()
	case reflect.Complex64:
		return reflect.TypeFor[complex64]()
	case reflect.Complex128:
		return reflect.TypeFor[complex128]()
	case reflect.UnsafePointer:
		return reflect.TypeFor[unsafe.Pointer]()
	default:
		return nil
	}
}

// AsBool returns the value as bool when the dynamic type's kind is bool.
func (k *Kind) AsBool() (bool, bool) {
	v, ok := k.convert(reflect.Bool)
	if !ok {
		return false, false
	}
	return v.Bool(), true
}

// AsString returns the value as string when the dynamic type's kind is string.
func (k *Kind) AsString() (string, bool) {
	v, ok := k.convert(reflect.String)
	if !ok {
		return "", false
	}
	return v.String(), true
}

// AsInt returns the value as int when the dynamic type's kind is int.
func (k *Kind) AsInt() (int, bool) {
	v, ok := k.convert(reflect.Int)
	if !ok {
		return 0, false
	}
	return int(v.Int()), true
}

// AsInt8 returns the value as int8 when the dynamic type's kind is int8.
func (k *Kind) AsInt8() (int8, bool) {
	v, ok := k.convert(reflect.Int8)
	if !ok {
		return 0, false
	}
	return int8(v.Int()), true
}

// AsInt16 returns the value as int16 when the dynamic type's kind is int16.
func (k *Kind) AsInt16() (int16, bool) {
	v, ok := k.convert(reflect.Int16)
	if !ok {
		return 0, false
	}
	return int16(v.Int()), true
}

// AsInt32 returns the value as int32 when the dynamic type's kind is int32.
func (k *Kind) AsInt32() (int32, bool) {
	v, ok := k.convert(reflect.Int32)
	if !ok {
		return 0, false
	}
	return int32(v.Int()), true
}

// AsInt64 returns the value as int64 when the dynamic type's kind is int64.
func (k *Kind) AsInt64() (int64, bool) {
	v, ok := k.convert(reflect.Int64)
	if !ok {
		return 0, false
	}
	return v.Int(), true
}

// AsUint returns the value as uint when the dynamic type's kind is uint.
func (k *Kind) AsUint() (uint, bool) {
	v, ok := k.convert(reflect.Uint)
	if !ok {
		return 0, false
	}
	return uint(v.Uint()), true
}

// AsUint8 returns the value as uint8 when the dynamic type's kind is uint8.
func (k *Kind) AsUint8() (uint8, bool) {
	v, ok := k.convert(reflect.Uint8)
	if !ok {
		return 0, false
	}
	return uint8(v.Uint()), true
}

// AsUint16 returns the value as uint16 when the dynamic type's kind is uint16.
func (k *Kind) AsUint16() (uint16, bool) {
	v, ok := k.convert(reflect.Uint16)
	if !ok {
		return 0, false
	}
	return uint16(v.Uint()), true
}

// AsUint32 returns the value as uint32 when the dynamic type's kind is uint32.
func (k *Kind) AsUint32() (uint32, bool) {
	v, ok := k.convert(reflect.Uint32)
	if !ok {
		return 0, false
	}
	return uint32(v.Uint()), true
}

// AsUint64 returns the value as uint64 when the dynamic type's kind is uint64.
func (k *Kind) AsUint64() (uint64, bool) {
	v, ok := k.convert(reflect.Uint64)
	if !ok {
		return 0, false
	}
	return v.Uint(), true
}

// AsUintptr returns the value as uintptr when the dynamic type's kind is uintptr.
func (k *Kind) AsUintptr() (uintptr, bool) {
	v, ok := k.convert(reflect.Uintptr)
	if !ok {
		return 0, false
	}
	return uintptr(v.Uint()), true
}

// AsFloat32 returns the value as float32 when the dynamic type's kind is float32.
func (k *Kind) AsFloat32() (float32, bool) {
	v, ok := k.convert(reflect.Float32)
	if !ok {
		return 0, false
	}
	return float32(v.Float()), true
}

// AsFloat64 returns the value as float64 when the dynamic type's kind is float64.
func (k *Kind) AsFloat64() (float64, bool) {
	v, ok := k.convert(reflect.Float64)
	if !ok {
		return 0, false
	}
	return v.Float(), true
}

// AsComplex64 returns the value as complex64 when the dynamic type's kind is complex64.
func (k *Kind) AsComplex64() (complex64, bool) {
	v, ok := k.convert(reflect.Complex64)
	if !ok {
		return 0, false
	}
	return complex64(v.Complex()), true
}

// AsComplex128 returns the value as complex128 when the dynamic type's kind is complex128.
func (k *Kind) AsComplex128() (complex128, bool) {
	v, ok := k.convert(reflect.Complex128)
	if !ok {
		return 0, false
	}
	return v.Complex(), true
}

// AsUnsafePointer returns the value as unsafe.Pointer when the dynamic type's
// kind is unsafe.Pointer.
func (k *Kind) AsUnsafePointer() (unsafe.Pointer, bool) {
	v, ok := k.convert(reflect.UnsafePointer)
	if !ok {
		return nil, false
	}
	return unsafe.Pointer(v.Pointer()), true
}
