package kind

import "reflect"

// IsEmpty reports whether the value is nil, zero, or has length zero for
// strings, arrays, slices, maps and channels. For type-only Kinds it reports
// true only for the nil Kind.
func (k *Kind) IsEmpty() bool {
	if k == nil || k.IsNil() {
		return true
	}
	if k.value == nil {
		return false
	}

	v := reflect.ValueOf(k.value)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	default:
		return v.IsZero()
	}
}

// IsTruthy reports whether the value is non-empty and non-zero.
func (k *Kind) IsTruthy() bool {
	return !k.IsEmpty()
}
