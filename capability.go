package kind

import (
	"database/sql"
	"database/sql/driver"
	flagpkg "flag"
	"io"
	"log/slog"
)

// EnvMarshaler is the shape used by values that can encode themselves as
// environment key/value pairs.
type EnvMarshaler interface {
	MarshalEnv() (map[string]string, error)
}

// EnvUnmarshaler is the shape used by values that can decode themselves from
// environment key/value pairs.
type EnvUnmarshaler interface {
	UnmarshalEnv(map[string]string) error
}

// Validator is the common shape for values that validate their own state.
type Validator interface {
	Validate() error
}

// Verifier is the common shape for values that report whether they are valid.
type Verifier interface {
	Valid() bool
}

// StringParser is the common shape for values that parse themselves from a string.
type StringParser interface {
	Parse(string) error
}

// BytesParser is the common shape for values that parse themselves from bytes.
type BytesParser interface {
	ParseBytes([]byte) error
}

// Setter is the common shape for values that can be set from a string.
type Setter interface {
	Set(string) error
}

// IsEnvMarshaler reports whether the type can marshal itself to environment values.
func (k *Kind) IsEnvMarshaler() bool {
	return CanImplement[EnvMarshaler](k)
}

// IsEnvUnmarshaler reports whether the type can unmarshal itself from environment values.
func (k *Kind) IsEnvUnmarshaler() bool {
	return CanImplement[EnvUnmarshaler](k)
}

// IsValidator reports whether the type can validate itself with Validate.
func (k *Kind) IsValidator() bool {
	return CanImplement[Validator](k)
}

// IsVerifier reports whether the type can report validity with Valid.
func (k *Kind) IsVerifier() bool {
	return CanImplement[Verifier](k)
}

// IsStringParser reports whether the type can parse itself from a string.
func (k *Kind) IsStringParser() bool {
	return CanImplement[StringParser](k)
}

// IsBytesParser reports whether the type can parse itself from bytes.
func (k *Kind) IsBytesParser() bool {
	return CanImplement[BytesParser](k)
}

// IsSetter reports whether the type can be set from a string.
func (k *Kind) IsSetter() bool {
	return CanImplement[Setter](k)
}

// IsFlagValue reports whether the type implements flag.Value.
func (k *Kind) IsFlagValue() bool {
	return CanImplement[flagpkg.Value](k)
}

// IsScanner reports whether the type implements sql.Scanner.
func (k *Kind) IsScanner() bool {
	return CanImplement[sql.Scanner](k)
}

// IsValuer reports whether the type implements driver.Valuer.
func (k *Kind) IsValuer() bool {
	return CanImplement[driver.Valuer](k)
}

// IsReader reports whether the type implements io.Reader.
func (k *Kind) IsReader() bool {
	return CanImplement[io.Reader](k)
}

// IsWriter reports whether the type implements io.Writer.
func (k *Kind) IsWriter() bool {
	return CanImplement[io.Writer](k)
}

// IsCloser reports whether the type implements io.Closer.
func (k *Kind) IsCloser() bool {
	return CanImplement[io.Closer](k)
}

// IsReaderFrom reports whether the type implements io.ReaderFrom.
func (k *Kind) IsReaderFrom() bool {
	return CanImplement[io.ReaderFrom](k)
}

// IsWriterTo reports whether the type implements io.WriterTo.
func (k *Kind) IsWriterTo() bool {
	return CanImplement[io.WriterTo](k)
}

// IsLogValuer reports whether the type implements slog.LogValuer.
func (k *Kind) IsLogValuer() bool {
	return CanImplement[slog.LogValuer](k)
}
