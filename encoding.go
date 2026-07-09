package kind

import (
	"encoding"
	"encoding/json"
	"fmt"
)

// IsTextMarshaler reports whether the type implements encoding.TextMarshaler.
func (k *Kind) IsTextMarshaler() bool {
	return CanImplement[encoding.TextMarshaler](k)
}

// IsTextUnmarshaler reports whether the type implements encoding.TextUnmarshaler.
func (k *Kind) IsTextUnmarshaler() bool {
	return CanImplement[encoding.TextUnmarshaler](k)
}

// IsBinaryMarshaler reports whether the type implements encoding.BinaryMarshaler.
func (k *Kind) IsBinaryMarshaler() bool {
	return CanImplement[encoding.BinaryMarshaler](k)
}

// IsBinaryUnmarshaler reports whether the type implements encoding.BinaryUnmarshaler.
func (k *Kind) IsBinaryUnmarshaler() bool {
	return CanImplement[encoding.BinaryUnmarshaler](k)
}

// IsJSONMarshaler reports whether the type implements json.Marshaler.
func (k *Kind) IsJSONMarshaler() bool {
	return CanImplement[json.Marshaler](k)
}

// IsJSONUnmarshaler reports whether the type implements json.Unmarshaler.
func (k *Kind) IsJSONUnmarshaler() bool {
	return CanImplement[json.Unmarshaler](k)
}

// IsError reports whether the type implements error.
func (k *Kind) IsError() bool {
	return CanImplement[error](k)
}

// IsStringer reports whether the type implements fmt.Stringer.
func (k *Kind) IsStringer() bool {
	return CanImplement[fmt.Stringer](k)
}
