package kind

import (
	"bytes"
	"database/sql/driver"
	flagpkg "flag"
	"log/slog"
	"testing"
)

type envCodec struct{}

func (envCodec) MarshalEnv() (map[string]string, error) {
	return map[string]string{"A": "B"}, nil
}

func (*envCodec) UnmarshalEnv(map[string]string) error {
	return nil
}

type validatable struct{}

func (*validatable) Validate() error { return nil }
func (validatable) Valid() bool      { return true }

type parsable struct{}

func (*parsable) Parse(string) error      { return nil }
func (*parsable) ParseBytes([]byte) error { return nil }
func (*parsable) Set(string) error        { return nil }
func (parsable) String() string           { return "" }

type sqlThing struct{}

func (*sqlThing) Scan(any) error              { return nil }
func (sqlThing) Value() (driver.Value, error) { return "x", nil }

type logThing struct{}

func (logThing) LogValue() slog.Value { return slog.StringValue("x") }

type closerThing struct{}

func (closerThing) Close() error { return nil }

func TestCapabilityHelpers(t *testing.T) {
	if !TypeOf[envCodec]().IsEnvMarshaler() || !TypeOf[envCodec]().IsEnvUnmarshaler() {
		t.Fatal("env capabilities failed")
	}
	if !TypeOf[validatable]().IsValidator() || !TypeOf[validatable]().IsVerifier() {
		t.Fatal("validation capabilities failed")
	}
	p := TypeOf[parsable]()
	if !p.IsStringParser() || !p.IsBytesParser() || !p.IsSetter() || !p.IsFlagValue() {
		t.Fatal("parser/setter capabilities failed")
	}
	if !TypeOf[*sqlThing]().IsScanner() || !TypeOf[sqlThing]().IsValuer() {
		t.Fatal("sql capabilities failed")
	}
	if !TypeOf[*bytes.Buffer]().IsReader() || !TypeOf[*bytes.Buffer]().IsWriter() {
		t.Fatal("io reader/writer capabilities failed")
	}
	if !TypeOf[closerThing]().IsCloser() {
		t.Fatal("io closer capability failed")
	}
	if !TypeOf[*bytes.Buffer]().IsReaderFrom() || !TypeOf[*bytes.Buffer]().IsWriterTo() {
		t.Fatal("io transfer capabilities failed")
	}
	if !TypeOf[logThing]().IsLogValuer() {
		t.Fatal("log valuer capability failed")
	}
	if !CanImplement[flagpkg.Value](TypeOf[parsable]()) {
		t.Fatal("CanImplement should detect pointer receiver")
	}
	if Implements[flagpkg.Value](TypeOf[parsable]()) {
		t.Fatal("Implements should remain strict")
	}
}
