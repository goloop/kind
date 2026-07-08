package kind

import (
	"encoding"
	"reflect"
	"testing"
)

type textValue string

func (t textValue) MarshalText() ([]byte, error) {
	return []byte(t), nil
}

type textPointer string

func (t *textPointer) UnmarshalText([]byte) error {
	return nil
}

func TestAssignability(t *testing.T) {
	k := TypeOf[textValue]()
	if !k.Implements(reflect.TypeFor[encoding.TextMarshaler]()) || !Implements[encoding.TextMarshaler](k) {
		t.Fatal("expected TextMarshaler implementation")
	}
	if !k.ConvertibleTo(reflect.TypeFor[string]()) || !ConvertibleTo[string](k) {
		t.Fatal("expected string convertibility")
	}
	if !TypeOf[string]().AssignableTo(reflect.TypeFor[string]()) || !AssignableTo[string](TypeOf[string]()) {
		t.Fatal("expected string assignability")
	}
	if !TypeOf[*textPointer]().Implements(reflect.TypeFor[encoding.TextUnmarshaler]()) {
		t.Fatal("expected pointer TextUnmarshaler implementation")
	}
}
