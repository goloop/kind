package kind

import (
	"reflect"
	"sync"
)

var typeCache sync.Map // map[reflect.Type]*descriptor
var typeBuildMu sync.Mutex

func descriptorOf(t reflect.Type) *descriptor {
	if t == nil {
		return nilKind.desc
	}
	if d, ok := typeCache.Load(t); ok {
		return d.(*descriptor)
	}

	typeBuildMu.Lock()
	defer typeBuildMu.Unlock()

	if d, ok := typeCache.Load(t); ok {
		return d.(*descriptor)
	}

	d := buildDescriptor(t, make(map[reflect.Type]*descriptor, 8))
	d = canonicalizeDescriptorTree(d, make(map[*descriptor]*descriptor, 8))
	publishDescriptorTree(d, make(map[*descriptor]struct{}, 8))
	return d
}

func canonicalizeDescriptorTree(
	d *descriptor,
	seen map[*descriptor]*descriptor,
) *descriptor {
	if d == nil {
		return nil
	}
	if canonical, ok := seen[d]; ok {
		return canonical
	}
	if d.t != nil {
		if existing, ok := typeCache.Load(d.t); ok {
			canonical := existing.(*descriptor)
			seen[d] = canonical
			return canonical
		}
	}

	seen[d] = d
	d.key = canonicalizeDescriptorTree(d.key, seen)
	d.elem = canonicalizeDescriptorTree(d.elem, seen)
	for i := range d.fields {
		if d.fields[i].Type != nil {
			d.fields[i].Type = canonicalizeDescriptorTree(d.fields[i].Type.desc, seen).kind
		}
	}
	return d
}

func publishDescriptorTree(d *descriptor, seen map[*descriptor]struct{}) {
	if d == nil {
		return
	}
	if _, ok := seen[d]; ok {
		return
	}
	seen[d] = struct{}{}

	if d.t != nil {
		typeCache.LoadOrStore(d.t, d)
	}
	publishDescriptorTree(d.key, seen)
	publishDescriptorTree(d.elem, seen)
	for _, f := range d.fields {
		if f.Type != nil {
			publishDescriptorTree(f.Type.desc, seen)
		}
	}
}

func buildDescriptor(t reflect.Type, seen map[reflect.Type]*descriptor) *descriptor {
	if t == nil {
		return nilKind.desc
	}
	if d, ok := typeCache.Load(t); ok {
		return d.(*descriptor)
	}
	if d, ok := seen[t]; ok {
		return d
	}

	d := &descriptor{name: t.String(), t: t}
	d.kind = &Kind{desc: d}
	seen[t] = d

	if t.Name() != "" {
		d.flags |= flagNamed
	}
	if isNilableType(t.Kind()) {
		d.flags |= flagNilable
	}

	classify(d, t, seen)
	return d
}

func classify(d *descriptor, t reflect.Type, seen map[reflect.Type]*descriptor) {
	switch t.Kind() {
	case reflect.Bool:
		d.flags |= flagBool
	case reflect.String:
		d.flags |= flagString
	case reflect.Int:
		d.flags |= flagInt
	case reflect.Int8:
		d.flags |= flagInt8
	case reflect.Int16:
		d.flags |= flagInt16
	case reflect.Int32:
		d.flags |= flagInt32
	case reflect.Int64:
		d.flags |= flagInt64
	case reflect.Uint:
		d.flags |= flagUint
	case reflect.Uint8:
		d.flags |= flagUint8
	case reflect.Uint16:
		d.flags |= flagUint16
	case reflect.Uint32:
		d.flags |= flagUint32
	case reflect.Uint64:
		d.flags |= flagUint64
	case reflect.Uintptr:
		d.flags |= flagUintptr
	case reflect.Float32:
		d.flags |= flagFloat32
	case reflect.Float64:
		d.flags |= flagFloat64
	case reflect.Complex64:
		d.flags |= flagComplex64
	case reflect.Complex128:
		d.flags |= flagComplex128
	case reflect.Array:
		d.flags |= flagArray
		elem := buildDescriptor(t.Elem(), seen)
		d.elem = elem
		if elem.has(flagSlice) || elem.has(flagSliceOfSlices) || elem.has(flagSliceOfArrays) {
			d.flags |= flagArrayOfSlices
			d.flags &^= flagArray
		}
		if elem.has(flagArray) || elem.has(flagArrayOfArrays) || elem.has(flagArrayOfSlices) {
			d.flags |= flagArrayOfArrays
			d.flags &^= flagArray
		}
		d.flags |= leafFlags(elem.flags)
	case reflect.Slice:
		d.flags |= flagSlice
		elem := buildDescriptor(t.Elem(), seen)
		d.elem = elem
		if elem.has(flagSlice) || elem.has(flagSliceOfSlices) || elem.has(flagSliceOfArrays) {
			d.flags |= flagSliceOfSlices
			d.flags &^= flagSlice
		}
		if elem.has(flagArray) || elem.has(flagArrayOfArrays) || elem.has(flagArrayOfSlices) {
			d.flags |= flagSliceOfArrays
			d.flags &^= flagSlice
		}
		d.flags |= leafFlags(elem.flags)
	case reflect.Pointer:
		d.flags |= flagPointer
		elem := buildDescriptor(t.Elem(), seen)
		d.elem = elem
		d.flags |= leafFlags(elem.flags)
	case reflect.Map:
		d.flags |= flagMap
		d.key = buildDescriptor(t.Key(), seen)
		d.elem = buildDescriptor(t.Elem(), seen)
		d.flags |= leafFlags(d.elem.flags)
	case reflect.Chan:
		d.flags |= flagChannel
		elem := buildDescriptor(t.Elem(), seen)
		d.elem = elem
		d.flags |= leafFlags(elem.flags)
	case reflect.Struct:
		d.flags |= flagStruct
		d.fields = buildFields(t, seen)
	case reflect.Interface:
		d.flags |= flagInterface
	case reflect.Func:
		d.flags |= flagFunction
	case reflect.UnsafePointer:
		d.flags |= flagUnsafePointer
	default:
		d.flags |= flagUndefined
	}
}

func leafFlags(flags flag) flag {
	const leaves = flagBool |
		flagString |
		flagInt8 |
		flagInt16 |
		flagInt32 |
		flagInt64 |
		flagUint8 |
		flagUint16 |
		flagUint32 |
		flagUint64 |
		flagInt |
		flagUint |
		flagUintptr |
		flagFloat32 |
		flagFloat64 |
		flagComplex64 |
		flagComplex128 |
		flagUnsafePointer

	return flags & leaves
}

func isNilableType(k reflect.Kind) bool {
	switch k {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return true
	default:
		return false
	}
}

func isNilValue(v reflect.Value) bool {
	if !v.IsValid() || !isNilableType(v.Kind()) {
		return false
	}
	return v.IsNil()
}
