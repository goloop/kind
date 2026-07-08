package kind

import "testing"

func TestValueState(t *testing.T) {
	if !Of("").IsEmpty() || !Of(0).IsEmpty() || !Of([]int{}).IsEmpty() {
		t.Fatal("empty values not detected")
	}
	if Of("x").IsEmpty() || !Of("x").IsTruthy() {
		t.Fatal("truthy string not detected")
	}
	if TypeOf[string]().IsEmpty() {
		t.Fatal("type-only string should not be empty")
	}
}
