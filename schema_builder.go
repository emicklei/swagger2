package swagger2

import (
	"reflect"

	"fmt"
	"github.com/emicklei/swagger2/model"
	"strings"
)

type SchemaBuilder struct {
	schemas map[string]*model.Schema
	schema  *model.Schema
}

func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{
		schemas: map[string]*model.Schema{},
		schema:  new(model.Schema),
	}
}

func (s *SchemaBuilder) Build(value interface{}) (*model.Schema, map[string]*model.Schema) {
	if value == nil {
		return nil, nil
	}
	rv := reflect.ValueOf(value)
	s.build(rv.Type())
	// In case we have a struct here, make sure we return a reference
	if rv.Type().Kind() == reflect.Struct {
		return &model.Schema{Ref: ref(rv.Type())}, s.schemas
	} else {
		return s.schema, s.schemas
	}

}

func (s *SchemaBuilder) Schemas(value map[string]*model.Schema) *SchemaBuilder {
	if value == nil {
		return nil
	}
	s.schemas = value
	return s
}

func (s *SchemaBuilder) build(t reflect.Type) *model.Schema {
	kind := t.Kind()

	if jsType := getTypeFromMapping(kind); jsType != "" {
		s.schema.Type = jsType
	}
	switch kind {
	case reflect.Slice:
		s.buildFromSlice(t)
	case reflect.Map:
		//p.buildFromMap(t)
	case reflect.Struct:
		s.buildFromStruct(t)
	case reflect.Ptr:
		s.build(t.Elem())
	}
	return s.schema
}
func (s *SchemaBuilder) buildFromSlice(t reflect.Type) {
	itemType := t.Elem()
	itemSchema := NewSchemaBuilder().Schemas(s.schemas).build(itemType)
	switch itemType.Kind() {
	case reflect.Struct:
		s.schema.Items = &model.Schema{Ref: ref(itemType)}
	default:
		s.schema.Items = itemSchema
	}
}

func (s *SchemaBuilder) buildFromStruct(t reflect.Type) {
	fc := t.NumField()

	// If we already know the struct, we just ned a reference
	if _, ok := s.schemas[t.Name()]; ok {
		s.schema = &model.Schema{Ref: ref(t)}
		return
	}

	s.schemas[t.Name()] = s.schema
	s.schema.Properties = map[string]*model.Schema{}

	for c := 0; c < fc; c++ {
		// find out the field name
		field := t.Field(c)
		fieldName := getFieldName(field)
		if fieldName == "" {
			continue
		}

		fieldType := field.Type
		for fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		// Consider type overrides
		overrideType := getOverrideType(field)
		if overrideType != "" {
			s.schema.Properties[fieldName] = &model.Schema{Type: overrideType}
			model.MapToGoValidator(s.schema.Properties[fieldName], field.Tag.Get("valid"), fieldType)
			continue
		}

		// Create the schema definition of the field and pass in all discovered schemas
		fieldSchema := NewSchemaBuilder().Schemas(s.schemas).build(fieldType)
		switch fieldType.Kind() {
		case reflect.Struct:
			s.schema.Properties[fieldName] = &model.Schema{Ref: ref(fieldType)}
			model.MapToGoValidator(s.schema.Properties[fieldName], field.Tag.Get("valid"), fieldType)
		default:
			model.MapToGoValidator(fieldSchema, field.Tag.Get("valid"), fieldType)
			s.schema.Properties[fieldName] = fieldSchema
		}
	}
}

// github.com/mcuadros/go-jsonschema-generator
var mapping = map[reflect.Kind]string{
	reflect.Bool:    "boolean",
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

func getOverrideType(t reflect.StructField) string {
	return t.Tag.Get("type")
}

func getFieldName(field reflect.StructField) string {
	if field.Anonymous {
		return ""
	}
	jsonTags := field.Tag.Get("json")
	if "-" == jsonTags {
		return ""
	}

	fieldName := ""
	if !strings.HasPrefix(jsonTags, ",") {
		fieldName = jsonTags
	} else {
		fieldName = strings.Split(jsonTags, ",")[0]
	}
	if fieldName == "" {
		fieldName = field.Name
	}
	return fieldName
}

func ref(t reflect.Type) string {
	return fmt.Sprintf("#/definitions/%s", t.Name())
}
