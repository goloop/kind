package kind

import "testing"

type taggedStruct struct {
	ID     int    `json:"id" env:"ID"`
	Name   string `json:"name"`
	hidden bool
}

func TestStructFields(t *testing.T) {
	k := TypeOf[taggedStruct]()
	fields := k.Fields()
	if len(fields) != 3 {
		t.Fatalf("Fields len = %d, want 3", len(fields))
	}
	exported := k.ExportedFields()
	if len(exported) != 2 {
		t.Fatalf("ExportedFields len = %d, want 2", len(exported))
	}
	if !k.HasField("ID") || k.HasField("Missing") {
		t.Fatal("HasField failed")
	}
	f, ok := k.Field("ID")
	if !ok || !f.Type.IsInt() || !f.HasTag("json") {
		t.Fatalf("unexpected ID field: %#v, %v", f, ok)
	}
	if got, ok := f.TagValue("env"); !ok || got != "ID" {
		t.Fatalf("TagValue(env) = %q, %v; want ID, true", got, ok)
	}
	if !k.HasTag("env") || len(k.FieldsByTag("json")) != 2 {
		t.Fatal("tag helpers failed")
	}

	fields[0].Index[0] = 99
	f, _ = k.Field("ID")
	if f.Index[0] == 99 {
		t.Fatal("Fields should return defensive copies")
	}
}
