[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/kind)](https://goreportcard.com/report/github.com/goloop/kind) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/kind/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/kind) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)


# kind

The "kind" package in GoLang is designed to facilitate the inspection and categorization of different data types at runtime, using reflection. It provides a unified way of querying the characteristics of a given value, determining whether it's a simple or complex type, and even obtaining its string representation. The package revolves around a central Kind struct, which encapsulates the nature of a particular data type and provides a myriad of methods to query its specifics.

Example:

```go
// Numeric types.
a := aind.Of(42)
fmt.Println(a.IsInt(), a.Name()) // true "int"

// Slices.
b := bind.Of([]int{1, 2, 3})
fmt.Println(b.IsSlice(), b.IsInt(), b.Name()) // true true "[]int"

// Slice of slices.
c := cind.Of([][]int{{1, 2, 3}, {4, 5, 6}})
fmt.Println(c.IsSliceOfSlices(), c.IsInt(), c.Name()) // true true "[][]int"

// Array of slices.
d := dind.Of([3][]int{{1,2, 3}, {4, 5, 6}, {7, 8, 9}})
fmt.Println(d.IsArrayOfSlices(), d.IsInt(), d.Name()) // true true "[3][]int"
```

Key features and components include:

 1. The Kind struct: This structure serves as a container for an exhaustive list of possible data types a given value can have in Go. This includes scalars (like bools, ints, uints, floats, and strings), complex types (like arrays, slices, maps, channels, and structs), and some language-specific types (like pointers, interfaces, and functions). It also has a field to hold the value itself and a name that stores a string representation of the type.

 2. Query methods: The Kind struct offers numerous methods that let you determine if the encapsulated value is of a particular type. Each method corresponds to a particular type (e.g., IsInt(), IsSlice(), IsMap()) and returns true if the value is of that type.

 3. Conversion methods: The struct also provides methods to attempt conversion of the encapsulated value to a specific type. If successful, these methods return the value in the requested type along with a success indicator (e.g., AsBool(), AsFloat32()).

 4. Helper functions: There are additional methods to assist with common tasks, such as determining whether a value is of a complex type (IsComplex()), or whether it is a signed or unsigned type (IsSigned(), IsUnsigned()). There's also an Is function that compares the name of the kind with a provided string, offering another way to determine the type.

 5. Map Key-Value Kind information: If the type in question is a Map, then MapKeyKind() and MapValueKind() provide Kind instances representing the type of key and value respectively.

 6. Of function: This package provides a standalone Of function that takes in a value and returns an instance of Kind that describes the type of the given value. It performs an initial check to see if the value is nil, and then proceeds to dissect the type of the value using reflection, populating the relevant fields in the Kind struct accordingly.

In summary, the "kind" package in GoLang is a robust utility for dealing with dynamic types, making it much easier to inspect, categorize, and convert different types of data at runtime.