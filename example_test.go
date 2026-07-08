package kind_test

import (
	"fmt"

	"github.com/goloop/kind"
)

func ExampleOf() {
	k := kind.Of([]int{1, 2, 3})

	fmt.Println(k.IsSlice())
	fmt.Println(k.IsInt())
	fmt.Println(k.Name())

	// Output:
	// true
	// true
	// []int
}

func ExampleTypeOf() {
	type Reader interface {
		Read([]byte) (int, error)
	}

	k := kind.TypeOf[Reader]()

	fmt.Println(k.IsInterface())
	fmt.Println(k.Kind())

	// Output:
	// true
	// interface
}

func ExampleKind_Fields() {
	type Config struct {
		Port int `env:"PORT" json:"port"`
	}

	k := kind.TypeOf[Config]()
	field, _ := k.Field("Port")
	jsonTag, _ := field.TagValue("json")

	fmt.Println(k.HasTag("env"))
	fmt.Println(field.Type.IsInt())
	fmt.Println(jsonTag)

	// Output:
	// true
	// true
	// port
}
