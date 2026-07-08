package kind

import (
	"encoding/json"
	"fmt"
	"testing"
)

type jsonThing struct{}

func (jsonThing) MarshalJSON() ([]byte, error) {
	return []byte(`{}`), nil
}

type jsonThingPtr struct{}

func (*jsonThingPtr) UnmarshalJSON([]byte) error {
	return nil
}

type stringThing struct{}

func (stringThing) String() string {
	return "thing"
}

type errThing struct{}

func (errThing) Error() string {
	return "err"
}

func TestEncodingHelpers(t *testing.T) {
	if !TypeOf[textValue]().IsTextMarshaler() {
		t.Fatal("expected text marshaler")
	}
	if !TypeOf[*textPointer]().IsTextUnmarshaler() {
		t.Fatal("expected text unmarshaler")
	}
	if !TypeOf[jsonThing]().IsJSONMarshaler() || !Implements[json.Marshaler](TypeOf[jsonThing]()) {
		t.Fatal("expected json marshaler")
	}
	if !TypeOf[*jsonThingPtr]().IsJSONUnmarshaler() {
		t.Fatal("expected json unmarshaler")
	}
	if !TypeOf[errThing]().IsError() {
		t.Fatal("expected error")
	}
	if !TypeOf[stringThing]().IsStringer() || !Implements[fmt.Stringer](TypeOf[stringThing]()) {
		t.Fatal("expected stringer")
	}
}
