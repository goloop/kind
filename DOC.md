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
func (k *Kind) Leaf() *Kind
func (k *Kind) Base() *Kind
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
IsEmpty() bool
IsTruthy() bool
```

Containers and composite/reference types:

```go
IsScalar() bool
IsContainer() bool
IsComposite() bool
IsReference() bool
IsComparable() bool
IsOrdered() bool
IsNumeric() bool
IsPointer() bool
IsPointerToStruct() bool
IsArray() bool
IsSlice() bool
IsSliceLike() bool
IsSliceOfSlices() bool
IsArrayOfSlices() bool
IsSliceOfArrays() bool
IsArrayOfArrays() bool
IsMap() bool
IsMapLike() bool
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

## Shape and Indirection

```go
func (k *Kind) ChanDir() reflect.ChanDir
func (k *Kind) Len() int
func (k *Kind) Cap() int
func (k *Kind) Depth() int
func (k *Kind) Deref() *Kind
func (k *Kind) Indirect() *Kind
func (k *Kind) PointerDepth() int
func (k *Kind) IsPointerTo(target reflect.Type) bool
func IsPointerTo[T any](k *Kind) bool
```

`Leaf` / `Base` follow pointer, array, slice, map and channel element types to
the terminal type. For example `[][]*User` has depth `3` and leaf `User`.

`Len` returns the value length for strings, arrays, slices, maps and channels.
For type-only arrays it returns the array length. Otherwise it returns `-1`.
`Cap` returns capacity for arrays, slices and channels, or `-1`.

## Struct Fields

```go
type Field struct {
    Name      string
    Type      *Kind
    Index     []int
    Tag       reflect.StructTag
    Anonymous bool
    Exported  bool
    Offset    uintptr
}

func (k *Kind) Fields() []Field
func (k *Kind) ExportedFields() []Field
func (k *Kind) HasField(name string) bool
func (k *Kind) Field(name string) (Field, bool)
func (k *Kind) HasTag(key string) bool
func (k *Kind) FieldsByTag(key string) []Field
func (f Field) HasTag(key string) bool
func (f Field) TagValue(key string) (string, bool)
```

Field APIs inspect direct struct fields. Returned slices and index values are
defensive copies, so callers can sort or modify them without mutating the cache.

## Assignability and Interfaces

Go does not support generic methods, so generic helpers are top-level
functions rather than `k.Implements[T]()` methods.

```go
func (k *Kind) Implements(target reflect.Type) bool
func (k *Kind) AssignableTo(target reflect.Type) bool
func (k *Kind) ConvertibleTo(target reflect.Type) bool
func Implements[T any](k *Kind) bool
func AssignableTo[T any](k *Kind) bool
func ConvertibleTo[T any](k *Kind) bool
```

For common parser/encoding checks:

```go
IsTextMarshaler() bool
IsTextUnmarshaler() bool
IsJSONMarshaler() bool
IsJSONUnmarshaler() bool
IsError() bool
IsStringer() bool
```

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

```go
type Config struct {
    Port int `env:"PORT" json:"port"`
}

k := kind.TypeOf[Config]()
field, _ := k.Field("Port")
fmt.Println(k.HasTag("env"))          // true
fmt.Println(field.Type.IsInt())       // true
fmt.Println(field.TagValue("json"))   // "port", true
```

## Performance

The package caches descriptors by `reflect.Type` in a `sync.Map`. `Of(value)`
still evaluates value-specific state (`nil`, `zero`) per call. `OfType`,
`TypeOf`, `Elem`, `Key`, `MapKeyKind` and `MapValueKind` reuse type-only `Kind`
values from the cache.
