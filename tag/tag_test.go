package tag

import (
	"testing"
)

// TestIsSingle tests the IsSingle method.
func TestIsSingle(t *testing.T) {
	tests := []struct {
		name  string
		input Tag
		want  bool
	}{
		{"one", Pointer, true},
		{"two", Pointer + Array, false},
		{"three", 0, false},
	}

	for _, tt := range tests {
		if got := tt.input.IsSingle(); got != tt.want {
			t.Errorf("IsSingle(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestIsComplex tests the IsComplex method.
func TestIsComplex(t *testing.T) {
	tests := []struct {
		name  string
		input Tag
		want  bool
	}{
		{"one", Pointer, false},
		{"two", Pointer + Array, true},
		{"three", 0, false},
	}

	for _, tt := range tests {
		if got := tt.input.IsComplex(); got != tt.want {
			t.Errorf("IsComplex(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestContains tests the Contains method.
func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		tag    Tag
		flag   Tag
		want   bool
		hasErr bool
	}{
		{"one", Pointer | Array, Pointer, true, false},
		{"two", Pointer | Array, Array, true, false},
		{"three", Pointer, Array, false, false},
		{"four", Pointer, 0, false, true},
		{"five", Pointer | Array, Pointer | Array, true, false},
		{"six", Pointer | Array, Pointer | Array | Int, false, false},
	}

	for _, tt := range tests {
		got, err := tt.tag.Contains(tt.flag)
		if err != nil && !tt.hasErr {
			t.Errorf("Unexpected error: %v", err)
		}

		if err == nil && tt.hasErr {
			t.Errorf("Expected error, but got none")
		}

		if got != tt.want {
			t.Errorf("Contains(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestIsValid tests the IsValid method.
func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input Tag
		want  bool
	}{
		{"one", Pointer, true},
		{"two", Pointer + Array, true},
		{"three", Pointer + Array + overflowTagValue, false},
		{"four", overflowTagValue, false},
	}

	for _, tt := range tests {
		if got := tt.input.IsValid(); got != tt.want {
			t.Errorf("IsValid(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestSet tests the Set method.
func TestSet(t *testing.T) {
	tests := []struct {
		name    string
		initial Tag
		flags   []Tag
		want    Tag
		hasErr  bool
	}{
		{"one", Pointer, []Tag{Array, Slice}, Array + Slice, false},
		{"two", Pointer, []Tag{Array, Slice, overflowTagValue}, Pointer, true},
		{"three", Pointer, []Tag{}, 0, false},
		{"four", 0, []Tag{Pointer, Array}, Pointer + Array, false},
		{"five", 0, []Tag{Pointer, Array, overflowTagValue}, 0, true},
	}

	for _, tt := range tests {
		got, err := tt.initial.Set(tt.flags...)
		if (err != nil) != tt.hasErr {
			t.Errorf("Set(%s) unexpected error status: got %v", tt.name, err)
		}

		if got != tt.want {
			t.Errorf("Set(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestAdd tests the Add method.
func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		initial Tag
		flags   []Tag
		want    Tag
		hasErr  bool
	}{
		{"one", Pointer, []Tag{Array, Slice}, Pointer + Array + Slice, false},
		{"two", Pointer, []Tag{Array, Slice, overflowTagValue}, Pointer, true},
		{"three", Pointer, []Tag{}, Pointer, false},
		{"four", 0, []Tag{Pointer, Array}, Pointer + Array, false},
		{"five", 0, []Tag{Pointer, Array, overflowTagValue}, 0, true},
	}

	for _, tt := range tests {
		got, err := tt.initial.Add(tt.flags...)
		if (err != nil) != tt.hasErr {
			t.Errorf("Add(%s) unexpected error status: got %v", tt.name, err)
		}

		if got != tt.want {
			t.Errorf("Add(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestDelete tests the Delete method.
func TestDelete(t *testing.T) {
	tests := []struct {
		name    string
		initial Tag
		flags   []Tag
		want    Tag
		hasErr  bool
	}{
		{
			"one",
			Pointer + Array + Slice,
			[]Tag{Array},
			Pointer + Slice,
			false,
		},
		{
			"two",
			Pointer + Array + Slice,
			[]Tag{Array, overflowTagValue},
			Pointer + Array + Slice,
			true,
		},
		{
			"three",
			Pointer + Array + Slice,
			[]Tag{},
			Pointer + Array + Slice,
			false,
		},
		{
			"four",
			Pointer + Array + Slice,
			[]Tag{Pointer, Array},
			Slice,
			false,
		},
		{
			"five",
			Pointer + Array + Slice,
			[]Tag{Pointer, Array, overflowTagValue},
			Pointer + Array + Slice,
			true,
		},
	}

	for _, tt := range tests {
		got, err := tt.initial.Delete(tt.flags...)
		if (err != nil) != tt.hasErr {
			t.Errorf("Delete(%s) unexpected error status: got %v",
				tt.name, err)
		}

		if got != tt.want {
			t.Errorf("Delete(%s) = %v, want %v",
				tt.name, got, tt.want)
		}
	}
}

// TestAll tests the All method.
func TestAll(t *testing.T) {
	tests := []struct {
		name  string
		tag   Tag
		flags []Tag
		want  bool
	}{
		{"all set", Pointer | Array, []Tag{Pointer, Array}, true},
		{"not all set", Pointer | Array, []Tag{Pointer, Slice}, false},
		{"empty flags", Pointer | Array, []Tag{}, true},
		{"empty tag", 0, []Tag{Pointer, Array}, false},
		{"one flag set", Pointer | Array, []Tag{Pointer}, true},
	}

	for _, tt := range tests {
		got := tt.tag.All(tt.flags...)
		if got != tt.want {
			t.Errorf("All(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestAny tests the Any method.
func TestAny(t *testing.T) {
	tests := []struct {
		name  string
		tag   Tag
		flags []Tag
		want  bool
	}{
		{"any set", Pointer | Array, []Tag{Pointer, Slice}, true},
		{"none set", Pointer, []Tag{Array, Slice}, false},
		{"empty flags", Pointer | Array, []Tag{}, false},
		{"empty tag", 0, []Tag{Pointer, Array}, false},
		{"one flag set", Pointer | Array, []Tag{Pointer}, true},
	}

	for _, tt := range tests {
		got := tt.tag.Any(tt.flags...)
		if got != tt.want {
			t.Errorf("Any(%s) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestFlags tests the Is* methods.
func TestFlags(t *testing.T) {
	tests := []struct {
		name     string
		tag      Tag
		isMethod func(Tag) bool
		want     bool
	}{
		{
			"IsNil",
			Nil,
			func(t Tag) bool { return t.IsNil() },
			true,
		},
		{
			"IsPointer",
			Pointer,
			func(t Tag) bool { return t.IsPointer() },
			true,
		},
		{
			"IsArray",
			Array,
			func(t Tag) bool { return t.IsArray() },
			true,
		},
		{
			"IsArrayOfArrays",
			ArrayOfArrays,
			func(t Tag) bool { return t.IsArrayOfArrays() },
			true,
		},
		{
			"IsArrayOfSlices",
			ArrayOfSlices,
			func(t Tag) bool { return t.IsArrayOfSlices() },
			true,
		},
		{
			"IsSlice",
			Slice,
			func(t Tag) bool { return t.IsSlice() },
			true,
		},
		{
			"IsSliceOfSlices",
			SliceOfSlices,
			func(t Tag) bool { return t.IsSliceOfSlices() },
			true,
		},
		{
			"IsSliceOfArrays",
			SliceOfArrays,
			func(t Tag) bool { return t.IsSliceOfArrays() },
			true,
		},
		{
			"IsStruct",
			Struct,
			func(t Tag) bool { return t.IsStruct() },
			true,
		},
		{
			"IsMap",
			Map,
			func(t Tag) bool { return t.IsMap() },
			true,
		},
		{
			"IsFunction",
			Func,
			func(t Tag) bool { return t.IsFunction() },
			true,
		},
		{
			"IsChannel",
			Chan,
			func(t Tag) bool { return t.IsChannel() },
			true,
		},
		{
			"IsBool",
			Bool,
			func(t Tag) bool { return t.IsBool() },
			true,
		},
		{
			"IsInt",
			Int,
			func(t Tag) bool { return t.IsInt() },
			true,
		},
		{
			"IsInt8",
			Int8,
			func(t Tag) bool { return t.IsInt8() },
			true,
		},
		{
			"IsInt16",
			Int16,
			func(t Tag) bool { return t.IsInt16() },
			true,
		},
		{
			"IsInt32",
			Int32,
			func(t Tag) bool { return t.IsInt32() },
			true,
		},
		{
			"IsInt64",
			Int64,
			func(t Tag) bool { return t.IsInt64() },
			true,
		},
		{
			"IsUint",
			Uint,
			func(t Tag) bool { return t.IsUint() },
			true,
		},
		{
			"IsUint8",
			Uint8,
			func(t Tag) bool { return t.IsUint8() },
			true,
		},
		{
			"IsUint16",
			Uint16,
			func(t Tag) bool { return t.IsUint16() },
			true,
		},
		{
			"IsUint32",
			Uint32,
			func(t Tag) bool { return t.IsUint32() },
			true,
		},
		{
			"IsUint64",
			Uint64,
			func(t Tag) bool { return t.IsUint64() },
			true,
		},
		{
			"IsUintptr",
			Uintptr,
			func(t Tag) bool { return t.IsUintptr() },
			true,
		},
		{
			"IsFloat32",
			Float32,
			func(t Tag) bool { return t.IsFloat32() },
			true,
		},
		{
			"IsFloat64",
			Float64,
			func(t Tag) bool { return t.IsFloat64() },
			true,
		},
		{
			"IsComplex64",
			Complex64,
			func(t Tag) bool { return t.IsComplex64() },
			true,
		},
		{
			"IsComplex128",
			Complex128,
			func(t Tag) bool { return t.IsComplex128() },
			true,
		},
		{
			"IsString",
			String,
			func(t Tag) bool { return t.IsString() },
			true,
		},
		{
			"IsUnsafePointer",
			UnsafePointer,
			func(t Tag) bool { return t.IsUnsafePointer() },
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.isMethod(tt.tag)
			if got != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
