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
