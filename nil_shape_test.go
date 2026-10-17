package kind

import "testing"

func TestNilKindShape(t *testing.T) {
	var k *Kind
	if k.Depth() != 0 || !k.Leaf().IsNil() || !k.Base().IsNil() {
		t.Fatal("unexpected nil shape")
	}
}
