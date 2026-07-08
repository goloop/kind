package kind

import "reflect"

// IsPointerToStruct reports whether the type is a pointer chain ending in a struct.
func (k *Kind) IsPointerToStruct() bool {
	return k.IsPointer() && k.Indirect().IsStruct()
}

// IsPointerTo reports whether the type is a pointer chain ending in target.
func (k *Kind) IsPointerTo(target reflect.Type) bool {
	if target == nil || !k.IsPointer() {
		return false
	}
	return k.Indirect().Type() == target
}

// IsPointerTo reports whether k is a pointer chain ending in T.
func IsPointerTo[T any](k *Kind) bool {
	return k.IsPointerTo(reflect.TypeFor[T]())
}

// Deref returns the element type when k is a pointer. For non-pointers it
// returns k itself.
func (k *Kind) Deref() *Kind {
	if k == nil {
		return nilKind
	}
	if k.Kind() != reflect.Pointer {
		return k
	}
	return k.Elem()
}

// Indirect follows pointer types until a non-pointer type is reached.
func (k *Kind) Indirect() *Kind {
	if k == nil {
		return nilKind
	}
	seen := make(map[*descriptor]struct{}, 4)
	for k.Kind() == reflect.Pointer {
		if _, ok := seen[k.desc]; ok {
			return k
		}
		seen[k.desc] = struct{}{}
		k = k.Elem()
	}
	return k
}

// PointerDepth returns the number of consecutive pointer indirections.
func (k *Kind) PointerDepth() int {
	depth := 0
	seen := make(map[*descriptor]struct{}, 4)
	for k != nil && k.Kind() == reflect.Pointer {
		if _, ok := seen[k.desc]; ok {
			return depth
		}
		seen[k.desc] = struct{}{}
		depth++
		k = k.Elem()
	}
	return depth
}
