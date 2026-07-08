package kind

import "reflect"

// IsScalar reports whether the type is a scalar value: bool, string or number.
func (k *Kind) IsScalar() bool {
	return k.IsBool() || k.IsString() || k.IsNumber()
}

// IsContainer reports whether the type is an array, slice, map or channel.
func (k *Kind) IsContainer() bool {
	return k.IsSliceLike() || k.IsMapLike() || k.IsChannel()
}

// IsComposite reports whether the type is built from other values: array,
// slice, map or struct.
func (k *Kind) IsComposite() bool {
	if k == nil {
		return false
	}
	switch k.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		return true
	default:
		return false
	}
}

// IsReference reports whether values of this type have reference-like
// semantics or may be nil.
func (k *Kind) IsReference() bool {
	if k == nil {
		return false
	}
	switch k.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return true
	default:
		return false
	}
}

// IsComparable reports whether values of this type are comparable with ==.
func (k *Kind) IsComparable() bool {
	t := k.Type()
	return t != nil && t.Comparable()
}

// IsOrdered reports whether values of this type support the language's
// ordered comparison operators.
func (k *Kind) IsOrdered() bool {
	switch k.Kind() {
	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsNumeric reports whether the type or leaf is any numeric type.
func (k *Kind) IsNumeric() bool {
	return k.IsNumber()
}
