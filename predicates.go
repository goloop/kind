package kind

// IsUndefined reports whether the type could not be classified.
func (k *Kind) IsUndefined() bool { return k.has(flagUndefined) }

// IsNil reports whether the value is nil. It is value-specific and is true for
// nil interfaces and typed nil pointers, slices, maps, channels and funcs.
func (k *Kind) IsNil() bool {
	return k == nil || k.isNil
}

// IsZero reports whether the value is Go's zero value for its type.
func (k *Kind) IsZero() bool {
	return k == nil || k.isZero
}

// IsNilable reports whether values of this type may be nil.
func (k *Kind) IsNilable() bool { return k.has(flagNilable) }

// IsNamed reports whether the type has a defined name.
func (k *Kind) IsNamed() bool { return k.has(flagNamed) }

// IsPointer reports whether the type is a pointer.
func (k *Kind) IsPointer() bool { return k.has(flagPointer) }

// IsArray reports whether the type is an array that is not classified as an
// array of arrays or array of slices.
func (k *Kind) IsArray() bool { return k.has(flagArray) }

// IsSlice reports whether the type is a slice that is not classified as a
// slice of arrays or slice of slices.
func (k *Kind) IsSlice() bool { return k.has(flagSlice) }

// IsSliceOfSlices reports whether the type is a slice whose element chain
// contains another slice before the leaf type.
func (k *Kind) IsSliceOfSlices() bool { return k.has(flagSliceOfSlices) }

// IsArrayOfSlices reports whether the type is an array whose element chain
// contains a slice before the leaf type.
func (k *Kind) IsArrayOfSlices() bool { return k.has(flagArrayOfSlices) }

// IsSliceOfArrays reports whether the type is a slice whose element chain
// contains an array before the leaf type.
func (k *Kind) IsSliceOfArrays() bool { return k.has(flagSliceOfArrays) }

// IsArrayOfArrays reports whether the type is an array whose element chain
// contains another array before the leaf type.
func (k *Kind) IsArrayOfArrays() bool { return k.has(flagArrayOfArrays) }

// IsMap reports whether the type is a map.
func (k *Kind) IsMap() bool { return k.has(flagMap) }

// IsStruct reports whether the type is a struct.
func (k *Kind) IsStruct() bool { return k.has(flagStruct) }

// IsInterface reports whether the type is an interface. Use OfType or TypeOf
// when you need the static interface type; Of sees the dynamic concrete type.
func (k *Kind) IsInterface() bool { return k.has(flagInterface) }

// IsFunction reports whether the type is a function.
func (k *Kind) IsFunction() bool { return k.has(flagFunction) }

// IsChannel reports whether the type is a channel.
func (k *Kind) IsChannel() bool { return k.has(flagChannel) }

// IsBool reports whether the type or its element leaf is bool.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsBool() bool { return k.has(flagBool) }

// IsString reports whether the type or its element leaf is string.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsString() bool { return k.has(flagString) }

// IsInt8 reports whether the type or its element leaf is int8.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsInt8() bool { return k.has(flagInt8) }

// IsInt16 reports whether the type or its element leaf is int16.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsInt16() bool { return k.has(flagInt16) }

// IsInt32 reports whether the type or its element leaf is int32.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsInt32() bool { return k.has(flagInt32) }

// IsInt64 reports whether the type or its element leaf is int64.
// For maps, this follows the value type, not the key type.
func (k *Kind) IsInt64() bool { return k.has(flagInt64) }

// IsUint8 reports whether the type or its element leaf is uint8.
func (k *Kind) IsUint8() bool { return k.has(flagUint8) }

// IsUint16 reports whether the type or its element leaf is uint16.
func (k *Kind) IsUint16() bool { return k.has(flagUint16) }

// IsUint32 reports whether the type or its element leaf is uint32.
func (k *Kind) IsUint32() bool { return k.has(flagUint32) }

// IsUint64 reports whether the type or its element leaf is uint64.
func (k *Kind) IsUint64() bool { return k.has(flagUint64) }

// IsInt reports whether the type or its element leaf is int.
func (k *Kind) IsInt() bool { return k.has(flagInt) }

// IsUint reports whether the type or its element leaf is uint.
func (k *Kind) IsUint() bool { return k.has(flagUint) }

// IsUintptr reports whether the type or its element leaf is uintptr.
func (k *Kind) IsUintptr() bool { return k.has(flagUintptr) }

// IsFloat32 reports whether the type or its element leaf is float32.
func (k *Kind) IsFloat32() bool { return k.has(flagFloat32) }

// IsFloat64 reports whether the type or its element leaf is float64.
func (k *Kind) IsFloat64() bool { return k.has(flagFloat64) }

// IsComplex64 reports whether the type or its element leaf is complex64.
func (k *Kind) IsComplex64() bool { return k.has(flagComplex64) }

// IsComplex128 reports whether the type or its element leaf is complex128.
func (k *Kind) IsComplex128() bool { return k.has(flagComplex128) }

// IsUnsafePointer reports whether the type or its element leaf is unsafe.Pointer.
func (k *Kind) IsUnsafePointer() bool { return k.has(flagUnsafePointer) }

// IsNumber reports whether the type or leaf is any numeric type.
func (k *Kind) IsNumber() bool {
	return k.IsAnyInt() || k.IsAnyFloat() || k.IsAnyComplex()
}

// IsAnyInt reports whether the type or leaf is any integer type, including uintptr.
func (k *Kind) IsAnyInt() bool {
	return k.has(flagInt8 | flagInt16 | flagInt32 | flagInt64 |
		flagUint8 | flagUint16 | flagUint32 | flagUint64 |
		flagInt | flagUint | flagUintptr)
}

// IsAnyFloat reports whether the type or leaf is float32 or float64.
func (k *Kind) IsAnyFloat() bool {
	return k.has(flagFloat32 | flagFloat64)
}

// IsAnyComplex reports whether the type or leaf is complex64 or complex128.
func (k *Kind) IsAnyComplex() bool {
	return k.has(flagComplex64 | flagComplex128)
}

// IsUnsigned reports whether the type or leaf is an unsigned integer,
// including uintptr.
func (k *Kind) IsUnsigned() bool {
	return k.has(flagUint8 | flagUint16 | flagUint32 | flagUint64 | flagUint | flagUintptr)
}

// IsSigned reports whether the type or leaf is signed numeric.
func (k *Kind) IsSigned() bool {
	return k.has(flagInt8 | flagInt16 | flagInt32 | flagInt64 |
		flagInt | flagFloat32 | flagFloat64 | flagComplex64 | flagComplex128)
}

// IsComplex reports whether the type is composite, reference-like, function,
// interface or unsafe pointer. Numeric complex64/complex128 are covered by
// IsAnyComplex.
func (k *Kind) IsComplex() bool {
	return k.has(flagPointer | flagArray | flagSlice | flagSliceOfSlices |
		flagArrayOfSlices | flagSliceOfArrays | flagArrayOfArrays |
		flagMap | flagStruct | flagInterface | flagFunction |
		flagChannel | flagUnsafePointer)
}
