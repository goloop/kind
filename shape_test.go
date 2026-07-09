package kind

import (
	"reflect"
	"testing"
)

type leafUser struct {
	Name string
}

func TestShape(t *testing.T) {
	k := TypeOf[[][]*leafUser]()
	if !k.IsSliceLike() || k.IsMapLike() {
		t.Fatal("slice-like/map-like classification failed")
	}
	if k.Depth() != 3 {
		t.Fatalf("Depth() = %d, want 3", k.Depth())
	}
	if leaf := k.Leaf(); leaf.Type() != reflect.TypeFor[leafUser]() {
		t.Fatalf("Leaf() = %v, want leafUser", leaf.Type())
	}
	if k.Base() != k.Leaf() {
		t.Fatal("Base should alias Leaf")
	}

	m := TypeOf[map[string][]int]()
	if !m.IsMapLike() || m.Depth() != 2 || !m.Leaf().IsInt() || m.IsString() {
		t.Fatalf("unexpected map shape: depth=%d leaf=%s", m.Depth(), m.Leaf())
	}

	if got := TypeOf[[4]int]().Len(); got != 4 {
		t.Fatalf("array Len() = %d, want 4", got)
	}
	if got := Of([]int{1, 2, 3}).Len(); got != 3 {
		t.Fatalf("slice Len() = %d, want 3", got)
	}
	if got := Of(make([]int, 2, 8)).Cap(); got != 8 {
		t.Fatalf("slice Cap() = %d, want 8", got)
	}
	if got := TypeOf[chan<- int]().ChanDir(); got != reflect.SendDir {
		t.Fatalf("ChanDir() = %v, want SendDir", got)
	}
}
