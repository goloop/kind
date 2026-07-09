[![deps.dev](https://img.shields.io/badge/deps.dev-insights-4c8dbc)](https://deps.dev/go/github.com%2Fgoloop%2Fkind) [![Go Reference](https://pkg.go.dev/badge/github.com/goloop/kind.svg)](https://pkg.go.dev/github.com/goloop/kind) [![License](https://img.shields.io/badge/license-MIT-brightgreen?style=flat)](https://github.com/goloop/kind/blob/master/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# kind

`kind` is a small cached reflection helper for Go. It classifies values and
types with readable predicates such as `IsSlice`, `IsMap`, `IsInt`,
`IsNilable`, `IsNamed`, `IsZero`, and exposes map key / element information
without repeating reflection code at every call site.

The package is intentionally narrow: it does not replace `reflect`, but wraps
the common type-inspection questions that appear in parsers, validators,
configuration loaders and diagnostics.

## Features

- Cached descriptors per `reflect.Type`.
- Value-aware `IsNil` and `IsZero`, including typed nil pointers, slices, maps,
  channels and funcs.
- Static type inspection through `TypeOf[T]()` and `OfType(reflect.Type)`.
- Container helpers: `Elem`, `Key`, `Leaf`, `Base`, `Depth`, `Len`, `Cap`.
- Category predicates: `IsScalar`, `IsContainer`, `IsComposite`,
  `IsReference`, `IsComparable`, `IsOrdered`, `IsNumeric`.
- Struct field/tag helpers for parser and config packages.
- Assignability and interface helpers, including generic top-level functions
  such as `kind.Implements[encoding.TextMarshaler](k)`.
- Capability helpers answer what a type can do: text/binary/JSON/env
  marshaling, validation, parsing, flag values, SQL scanner/valuer, IO and
  slog value support.
- Numeric and scalar predicates, including named types and `uintptr`.
- For maps, scalar leaf predicates describe the value type; inspect keys with
  `Key` or `MapKeyKind`.
- Safe `As*` helpers for scalar values; named scalar types convert without
  panicking.
- Zero third-party dependencies.

## Installation

```shell
go get github.com/goloop/kind
```

Requires Go 1.24 or newer.

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/goloop/kind"
)

func main() {
    k := kind.Of(map[string][]int{"one": {1, 2, 3}})

    fmt.Println(k.IsMap())                   // true
    fmt.Println(k.MapKeyKind().IsString())   // true
    fmt.Println(k.MapValueKind().IsSlice())  // true
    fmt.Println(k.MapValueKind().IsInt())    // true
    fmt.Println(k.IsString())                 // false; map key does not leak into leaf flags
    fmt.Println(k.Depth())                    // 2
    fmt.Println(k.Leaf().Name())              // int
}
```

Static interface types should be inspected with `TypeOf`:

```go
type Reader interface {
    Read([]byte) (int, error)
}

k := kind.TypeOf[Reader]()
fmt.Println(k.IsInterface()) // true
```

Capability checks are based on interfaces and method sets, not package names:

```go
type Request struct{}

func (*Request) Validate() error { return nil }

k := kind.TypeOf[Request]()
fmt.Println(k.IsValidator()) // true
```

Struct tags are cached with the type descriptor:

```go
type Config struct {
    Port int `env:"PORT" json:"port"`
}

k := kind.TypeOf[Config]()
field, _ := k.Field("Port")

fmt.Println(k.HasTag("env"))        // true
fmt.Println(field.Type.IsInt())     // true
fmt.Println(field.Tag.Get("json"))  // port
```

## Notes

`Of(value)` sees the dynamic type stored in an interface, exactly like
`reflect.TypeOf`. If you need a static interface type, use `TypeOf[T]()` or
`OfType`.

## Contributing

Before submitting changes, run:

```shell
go test ./...
go vet ./...
gofmt -l .
```

## License

`kind` is released under the MIT License. See [LICENSE](LICENSE).
