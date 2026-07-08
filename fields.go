package kind

import "reflect"

// Field describes a struct field.
type Field struct {
	Name      string
	Type      *Kind
	Index     []int
	Tag       reflect.StructTag
	Anonymous bool
	Exported  bool
	Offset    uintptr
}

func buildFields(t reflect.Type, seen map[reflect.Type]*descriptor) []Field {
	fields := make([]Field, t.NumField())
	for i := range t.NumField() {
		sf := t.Field(i)
		index := make([]int, len(sf.Index))
		copy(index, sf.Index)
		fields[i] = Field{
			Name:      sf.Name,
			Type:      buildDescriptor(sf.Type, seen).kind,
			Index:     index,
			Tag:       sf.Tag,
			Anonymous: sf.Anonymous,
			Exported:  sf.IsExported(),
			Offset:    sf.Offset,
		}
	}
	return fields
}

// Fields returns all direct struct fields. The returned slice and Index values
// are copies and may be modified by the caller.
func (k *Kind) Fields() []Field {
	if k == nil || !k.IsStruct() || len(k.desc.fields) == 0 {
		return nil
	}
	return cloneFields(k.desc.fields)
}

// ExportedFields returns direct exported struct fields.
func (k *Kind) ExportedFields() []Field {
	if k == nil || !k.IsStruct() || len(k.desc.fields) == 0 {
		return nil
	}

	out := make([]Field, 0, len(k.desc.fields))
	for _, f := range k.desc.fields {
		if f.Exported {
			out = append(out, cloneField(f))
		}
	}
	return out
}

// HasField reports whether a direct struct field with name exists.
func (k *Kind) HasField(name string) bool {
	_, ok := k.Field(name)
	return ok
}

// Field returns a direct struct field by name.
func (k *Kind) Field(name string) (Field, bool) {
	if k == nil || !k.IsStruct() {
		return Field{}, false
	}
	for _, f := range k.desc.fields {
		if f.Name == name {
			return cloneField(f), true
		}
	}
	return Field{}, false
}

// HasTag reports whether any direct struct field has a non-empty tag value for key.
func (k *Kind) HasTag(key string) bool {
	if k == nil || !k.IsStruct() {
		return false
	}
	for _, f := range k.desc.fields {
		if _, ok := f.Tag.Lookup(key); ok {
			return true
		}
	}
	return false
}

// FieldsByTag returns direct struct fields that define tag key.
func (k *Kind) FieldsByTag(key string) []Field {
	if k == nil || !k.IsStruct() {
		return nil
	}

	out := make([]Field, 0)
	for _, f := range k.desc.fields {
		if _, ok := f.Tag.Lookup(key); ok {
			out = append(out, cloneField(f))
		}
	}
	return out
}

// HasTag reports whether the field has a non-empty tag value for key.
func (f Field) HasTag(key string) bool {
	_, ok := f.Tag.Lookup(key)
	return ok
}

// TagValue returns the field tag value for key.
func (f Field) TagValue(key string) (string, bool) {
	return f.Tag.Lookup(key)
}

func cloneFields(fields []Field) []Field {
	out := make([]Field, len(fields))
	for i, f := range fields {
		out[i] = cloneField(f)
	}
	return out
}

func cloneField(f Field) Field {
	if len(f.Index) != 0 {
		index := make([]int, len(f.Index))
		copy(index, f.Index)
		f.Index = index
	}
	return f
}
