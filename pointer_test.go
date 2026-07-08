package kind

import (
	"testing"
)

type pointerUser struct {
	Name string
}

func TestPointerHelpers(t *testing.T) {
	k := TypeOf[***pointerUser]()
	if !k.IsPointerToStruct() {
		t.Fatal("expected pointer to struct")
	}
	if k.PointerDepth() != 3 {
		t.Fatalf("PointerDepth() = %d, want 3", k.PointerDepth())
	}
	if !IsPointerTo[pointerUser](k) {
		t.Fatal("expected pointer to pointerUser")
	}
	if k.Deref().PointerDepth() != 2 {
		t.Fatalf("Deref pointer depth = %d, want 2", k.Deref().PointerDepth())
	}
	if !k.Indirect().IsStruct() {
		t.Fatal("Indirect should end at struct")
	}
}
