# kind Reference

`kind` is a cached reflection helper for classifying Go values and static
types. It targets Go 1.24+ and has no third-party dependencies.

## Constructors

```go
func Of(v any) *Kind
func OfType(t reflect.Type) *Kind
func TypeOf[T any]() *Kind
```

- `Of` inspects the dynamic type and value state of `v`.
- `OfType` inspects a `reflect.Type` without a value.
- `TypeOf[T]` inspects a static generic type and is the preferred way to ask
  about interface types.

## Core Methods

```go
func (k *Kind) Type() reflect.Type
func (k *Kind) Kind() reflect.Kind
func (k *Kind) Name() string
func (k *Kind) String() string
func (k *Kind) Value() any
func (k *Kind) Is(name string) bool
func (k *Kind) Elem() *Kind
func (k *Kind) Key() *Kind
func (k *Kind) MapKeyKind() *Kind
func (k *Kind) MapValueKind() *Kind
```

`Elem` works for pointers, arrays, slices, maps and channels. `Key` works for
maps. If there is no such type component, the methods return a nil Kind whose
`Name()` is `"nil"`.

## Predicates

Value state:

```go
IsNil() bool
IsZero() bool
IsNilable() bool
IsNamed() bool
```

Containers and composite/reference types:

```go
IsPointer() bool
IsArray() bool
IsSlice() bool
IsSliceOfSlices() bool
IsArrayOfSlices() bool
IsSliceOfArrays() bool
IsArrayOfArrays() bool
IsMap() bool
IsStruct() bool
IsInterface() bool
IsFunction() bool
IsChannel() bool
IsUnsafePointer() bool
IsComplex() bool
```

Scalars:

```go
IsBool() bool
IsString() bool
IsInt() bool
IsInt8() bool
IsInt16() bool
IsInt32() bool
IsInt64() bool
IsUint() bool
IsUint8() bool
IsUint16() bool
IsUint32() bool
IsUint64() bool
IsUintptr() bool
IsFloat32() bool
IsFloat64() bool
IsComplex64() bool
IsComplex128() bool
IsNumber() bool
IsAnyInt() bool
IsAnyFloat() bool
IsAnyComplex() bool
IsUnsigned() bool
IsSigned() bool
```

Container predicates keep the historical behaviour where the leaf type is also
visible. For example, `kind.Of([]int{}).IsSlice()` and
`kind.Of([]int{}).IsInt()` both return true.

## Conversions

```go
AsBool() (bool, bool)
AsString() (string, bool)
AsInt() (int, bool)
AsInt8() (int8, bool)
AsInt16() (int16, bool)
AsInt32() (int32, bool)
AsInt64() (int64, bool)
AsUint() (uint, bool)
AsUint8() (uint8, bool)
AsUint16() (uint16, bool)
AsUint32() (uint32, bool)
AsUint64() (uint64, bool)
AsUintptr() (uintptr, bool)
AsFloat32() (float32, bool)
AsFloat64() (float64, bool)
AsComplex64() (complex64, bool)
AsComplex128() (complex128, bool)
AsUnsafePointer() (unsafe.Pointer, bool)
```

The `As*` methods return `ok=false` when the dynamic value is not the requested
kind. Named scalar types are converted to their built-in representation instead
of panicking.

## Examples

```go
k := kind.Of(map[string][]int{"one": {1, 2, 3}})

fmt.Println(k.IsMap())                  // true
fmt.Println(k.MapKeyKind().IsString())  // true
fmt.Println(k.MapValueKind().IsSlice()) // true
fmt.Println(k.MapValueKind().IsInt())   // true
```

```go
type Reader interface {
    Read([]byte) (int, error)
}

k := kind.TypeOf[Reader]()
fmt.Println(k.IsInterface()) // true
```

## Performance

The package caches descriptors by `reflect.Type` in a `sync.Map`. `Of(value)`
still evaluates value-specific state (`nil`, `zero`) per call. `OfType`,
`TypeOf`, `Elem`, `Key`, `MapKeyKind` and `MapValueKind` reuse type-only `Kind`
values from the cache.
