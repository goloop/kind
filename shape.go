package kind

import "reflect"

// IsSliceLike reports whether the type is a slice or array.
func (k *Kind) IsSliceLike() bool {
	if k == nil {
		return false
	}
	switch k.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

// IsMapLike reports whether the type is a map.
func (k *Kind) IsMapLike() bool {
	return k.Kind() == reflect.Map
}

// ChanDir returns the channel direction. It returns 0 for non-channel types.
func (k *Kind) ChanDir() reflect.ChanDir {
	t := k.Type()
	if t == nil || t.Kind() != reflect.Chan {
		return 0
	}
	return t.ChanDir()
}

// Len returns the value length for strings, arrays, slices, maps and channels.
// For type-only array Kinds it returns the array length. It returns -1 when a
// length is not available.
func (k *Kind) Len() int {
	if k == nil {
		return -1
	}
	if k.value != nil {
		v := reflect.ValueOf(k.value)
		switch v.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			return v.Len()
		default:
			return -1
		}
	}

	t := k.Type()
	if t != nil && t.Kind() == reflect.Array {
		return t.Len()
	}
	return -1
}

// Cap returns the value capacity for arrays, slices and channels. It returns
// -1 when a capacity is not available.
func (k *Kind) Cap() int {
	if k == nil || k.value == nil {
		if t := k.Type(); t != nil && t.Kind() == reflect.Array {
			return t.Len()
		}
		return -1
	}

	v := reflect.ValueOf(k.value)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Slice:
		return v.Cap()
	default:
		return -1
	}
}

// Depth returns how many pointer/container indirections lead to the leaf type.
func (k *Kind) Depth() int {
	depth, _ := walkLeaf(k.desc)
	return depth
}

// Leaf returns the terminal type after following pointers, arrays, slices,
// maps and channels through their element type. It is safe for recursive types.
func (k *Kind) Leaf() *Kind {
	_, leaf := walkLeaf(k.desc)
	if leaf == nil {
		return nilKind
	}
	return leaf.kind
}

// Base is an alias for Leaf.
func (k *Kind) Base() *Kind {
	return k.Leaf()
}

func walkLeaf(d *descriptor) (int, *descriptor) {
	if d == nil {
		return 0, nilKind.desc
	}

	seen := make(map[*descriptor]struct{}, 8)
	depth := 0
	for d != nil {
		if _, ok := seen[d]; ok {
			return depth, d
		}
		seen[d] = struct{}{}

		switch d.t.Kind() {
		case reflect.Pointer, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
			if d.elem == nil {
				return depth, d
			}
			depth++
			d = d.elem
		default:
			return depth, d
		}
	}

	return depth, nilKind.desc
}
