package kind

import (
	"reflect"
	"strings"
)

type flag uint64

const (
	flagUndefined flag = 1 << iota
	flagNil
	flagPointer
	flagArray
	flagSlice
	flagSliceOfSlices
	flagArrayOfSlices
	flagSliceOfArrays
	flagArrayOfArrays
	flagMap
	flagStruct
	flagInterface
	flagFunction
	flagChannel
	flagBool
	flagString
	flagInt8
	flagInt16
	flagInt32
	flagInt64
	flagUint8
	flagUint16
	flagUint32
	flagUint64
	flagInt
	flagUint
	flagUintptr
	flagFloat32
	flagFloat64
	flagComplex64
	flagComplex128
	flagUnsafePointer
	flagNamed
	flagNilable
)

const nilName = "nil"

// Kind describes a Go value's dynamic type and value-specific state. It is
// immutable after construction and safe to read concurrently.
type Kind struct {
	desc   *descriptor
	value  any
	isNil  bool
	isZero bool
}

type descriptor struct {
	name  string
	t     reflect.Type
	flags flag
	key   *descriptor
	elem  *descriptor
	kind  *Kind
}

var nilKind = &Kind{
	desc: &descriptor{
		name:  nilName,
		flags: flagNil,
	},
	isNil:  true,
	isZero: true,
}

// Of returns a Kind that describes the dynamic type of v.
func Of(v any) *Kind {
	if v == nil {
		return nilKind
	}

	t := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	d := descriptorOf(t)

	return &Kind{
		desc:   d,
		value:  v,
		isNil:  isNilValue(rv),
		isZero: rv.IsZero(),
	}
}

// OfType returns a Kind that describes t without binding it to a value.
func OfType(t reflect.Type) *Kind {
	if t == nil {
		return nilKind
	}
	return descriptorOf(t).kind
}

// TypeOf is a generic helper that returns a Kind for T's static type.
func TypeOf[T any]() *Kind {
	return OfType(reflect.TypeFor[T]())
}

// Type returns the reflected type. It returns nil for the nil Kind.
func (k *Kind) Type() reflect.Type {
	if k == nil || k.desc == nil {
		return nil
	}
	return k.desc.t
}

// Kind returns the underlying reflect.Kind. It returns reflect.Invalid for nil.
func (k *Kind) Kind() reflect.Kind {
	if k == nil || k.desc == nil || k.desc.t == nil {
		return reflect.Invalid
	}
	return k.desc.t.Kind()
}

// Name returns the canonical Go type name.
func (k *Kind) Name() string {
	if k == nil || k.desc == nil {
		return nilName
	}
	return k.desc.name
}

// String returns the canonical Go type name.
func (k *Kind) String() string {
	return k.Name()
}

// Value returns the original value passed to Of. For OfType and nil it returns nil.
func (k *Kind) Value() any {
	if k == nil {
		return nil
	}
	return k.value
}

// Elem returns the element Kind for pointers, arrays, slices, maps and channels.
func (k *Kind) Elem() *Kind {
	if k == nil || k.desc == nil || k.desc.elem == nil {
		return nilKind
	}
	return k.desc.elem.kind
}

// Key returns the key Kind for maps.
func (k *Kind) Key() *Kind {
	if k == nil || k.desc == nil || k.desc.key == nil {
		return nilKind
	}
	return k.desc.key.kind
}

// MapKeyKind returns the key Kind for maps.
func (k *Kind) MapKeyKind() *Kind {
	return k.Key()
}

// MapValueKind returns the value Kind for maps.
func (k *Kind) MapValueKind() *Kind {
	return k.Elem()
}

// Is reports whether name matches the canonical Go type name. Spaces are
// ignored, and built-in names are also compared case-insensitively for
// compatibility with older releases.
func (k *Kind) Is(name string) bool {
	got := compactName(k.Name())
	want := compactName(name)
	return got == want || strings.EqualFold(got, want)
}

func compactName(s string) string {
	return strings.Join(strings.Fields(s), "")
}

func (k *Kind) has(f flag) bool {
	if k == nil || k.desc == nil {
		return f == flagNil
	}
	return k.desc.flags&f != 0
}

func (d *descriptor) has(f flag) bool {
	return d != nil && d.flags&f != 0
}
