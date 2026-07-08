package kind

import "testing"

var benchKind *Kind

func BenchmarkOfInt(b *testing.B) {
	for b.Loop() {
		benchKind = Of(42)
	}
}

func BenchmarkOfSliceCached(b *testing.B) {
	v := []int{1, 2, 3}
	Of(v)
	b.ResetTimer()
	for b.Loop() {
		benchKind = Of(v)
	}
}

func BenchmarkOfMapCached(b *testing.B) {
	v := map[string][]int{"one": {1, 2, 3}}
	Of(v)
	b.ResetTimer()
	for b.Loop() {
		benchKind = Of(v)
	}
}

func BenchmarkTypeOf(b *testing.B) {
	for b.Loop() {
		benchKind = TypeOf[map[string][]int]()
	}
}
