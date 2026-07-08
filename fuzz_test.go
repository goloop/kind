package kind

import "testing"

func FuzzOfScalars(f *testing.F) {
	f.Add("", int64(0), uint64(0), float64(0), true)
	f.Add("hello", int64(-42), uint64(42), float64(3.14), false)

	f.Fuzz(func(t *testing.T, s string, i int64, u uint64, fl float64, b bool) {
		if k := Of(s); !k.IsString() {
			t.Fatalf("string classified as %s", k)
		}
		if k := Of(i); !k.IsInt64() || !k.IsSigned() {
			t.Fatalf("int64 classified as %s", k)
		}
		if k := Of(u); !k.IsUint64() || !k.IsUnsigned() {
			t.Fatalf("uint64 classified as %s", k)
		}
		if k := Of(fl); !k.IsFloat64() || !k.IsSigned() {
			t.Fatalf("float64 classified as %s", k)
		}
		if k := Of(b); !k.IsBool() {
			t.Fatalf("bool classified as %s", k)
		}
	})
}
