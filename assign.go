package kind

import "reflect"

// Implements reports whether the type implements target. target may be either
// an interface type or a pointer to an interface type.
func (k *Kind) Implements(target reflect.Type) bool {
	t := k.Type()
	target = interfaceType(target)
	return t != nil && target != nil && t.Implements(target)
}

// AssignableTo reports whether values of this type are assignable to target.
func (k *Kind) AssignableTo(target reflect.Type) bool {
	t := k.Type()
	return t != nil && target != nil && t.AssignableTo(target)
}

// ConvertibleTo reports whether values of this type are convertible to target.
func (k *Kind) ConvertibleTo(target reflect.Type) bool {
	t := k.Type()
	return t != nil && target != nil && t.ConvertibleTo(target)
}

// CanImplement reports whether the type or, for non-pointer types, a pointer
// to the type implements target. It is useful for addressable values whose
// parsing/validation methods usually have pointer receivers.
func (k *Kind) CanImplement(target reflect.Type) bool {
	t := k.Type()
	target = interfaceType(target)
	if t == nil || target == nil {
		return false
	}
	if t.Implements(target) {
		return true
	}
	return t.Kind() != reflect.Pointer && reflect.PointerTo(t).Implements(target)
}

// Implements reports whether k implements interface T.
func Implements[T any](k *Kind) bool {
	return k.Implements(reflect.TypeFor[T]())
}

// AssignableTo reports whether k is assignable to T.
func AssignableTo[T any](k *Kind) bool {
	return k.AssignableTo(reflect.TypeFor[T]())
}

// ConvertibleTo reports whether k is convertible to T.
func ConvertibleTo[T any](k *Kind) bool {
	return k.ConvertibleTo(reflect.TypeFor[T]())
}

// CanImplement reports whether k or *k implements interface T.
func CanImplement[T any](k *Kind) bool {
	return k.CanImplement(reflect.TypeFor[T]())
}

func interfaceType(t reflect.Type) reflect.Type {
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Interface {
		return t
	}
	if t.Kind() == reflect.Pointer && t.Elem().Kind() == reflect.Interface {
		return t.Elem()
	}
	return nil
}
