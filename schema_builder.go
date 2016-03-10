package swagger2

import (
	"reflect"

	"github.com/emicklei/swagger2/model"
)

type SchemaBuilder struct {
	schemas map[string]*model.Schema
	schema  *model.Schema
}

func NewSchemaBuilder() SchemaBuilder {
	return SchemaBuilder{
		schemas: map[string]*model.Schema{},
		schema:  new(model.Schema),
	}
}

func (s *SchemaBuilder) Build(value interface{}) *model.Schema {
	if value == nil {
		return nil
	}
	rv := reflect.ValueOf(value)
	s.build(rv.Type())
	return s.schema
}

func (s *SchemaBuilder) build(t reflect.Type) {
	kind := t.Kind()

	if jsType := getTypeFromMapping(kind); jsType != "" {
		s.schema.Type = jsType
	}
	switch kind {
	case reflect.Slice:
		//p.buildFromSlice(t)
	case reflect.Map:
		//p.buildFromMap(t)
	case reflect.Struct:
		//p.buildFromStruct(t)
	case reflect.Ptr:
		//p.build(t.Elem())
	}
}

// github.com/mcuadros/go-jsonschema-generator
var mapping = map[reflect.Kind]string{
	reflect.Bool:    "bool",
	reflect.Int:     "integer",
	reflect.Int8:    "integer",
	reflect.Int16:   "integer",
	reflect.Int32:   "integer",
	reflect.Int64:   "integer",
	reflect.Uint:    "integer",
	reflect.Uint8:   "integer",
	reflect.Uint16:  "integer",
	reflect.Uint32:  "integer",
	reflect.Uint64:  "integer",
	reflect.Float32: "number",
	reflect.Float64: "number",
	reflect.String:  "string",
	reflect.Slice:   "array",
	reflect.Struct:  "object",
	reflect.Map:     "object",
}

func getTypeFromMapping(k reflect.Kind) string {
	if t, ok := mapping[k]; ok {
		return t
	}

	return ""
}
