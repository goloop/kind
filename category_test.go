package kind

import "testing"

func TestCategories(t *testing.T) {
	if !Of(1).IsScalar() || !Of("x").IsScalar() || !Of(1.5).IsNumeric() {
		t.Fatal("scalar/numeric classification failed")
	}
	if Of(complex(1, 2)).IsOrdered() {
		t.Fatal("complex values are not ordered")
	}
	if !Of(1).IsOrdered() || !Of("x").IsOrdered() {
		t.Fatal("ordered classification failed")
	}
	if !Of([]int{}).IsContainer() || !Of([1]int{}).IsContainer() ||
		!Of(map[string]int{}).IsContainer() || !Of(make(chan int)).IsContainer() {
		t.Fatal("container classification failed")
	}
	if !Of(struct{ A int }{}).IsComposite() || Of(new(int)).IsComposite() {
		t.Fatal("composite classification failed")
	}
	if !Of(new(int)).IsReference() || !Of([]int{}).IsReference() || Of(1).IsReference() {
		t.Fatal("reference classification failed")
	}
	if !Of([2]int{}).IsComparable() || Of([]int{}).IsComparable() {
		t.Fatal("comparable classification failed")
	}
}
