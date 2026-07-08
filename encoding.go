package kind

import (
	"encoding"
	"encoding/json"
	"fmt"
)

// IsTextMarshaler reports whether the type implements encoding.TextMarshaler.
func (k *Kind) IsTextMarshaler() bool {
	return Implements[encoding.TextMarshaler](k)
}

// IsTextUnmarshaler reports whether the type implements encoding.TextUnmarshaler.
func (k *Kind) IsTextUnmarshaler() bool {
	return Implements[encoding.TextUnmarshaler](k)
}

// IsJSONMarshaler reports whether the type implements json.Marshaler.
func (k *Kind) IsJSONMarshaler() bool {
	return Implements[json.Marshaler](k)
}

// IsJSONUnmarshaler reports whether the type implements json.Unmarshaler.
func (k *Kind) IsJSONUnmarshaler() bool {
	return Implements[json.Unmarshaler](k)
}

// IsError reports whether the type implements error.
func (k *Kind) IsError() bool {
	return Implements[error](k)
}

// IsStringer reports whether the type implements fmt.Stringer.
func (k *Kind) IsStringer() bool {
	return Implements[fmt.Stringer](k)
}
