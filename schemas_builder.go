package swagger

type SchemasBuilder struct {
	schemas map[string]Schema
}

func NewSchemasBuilder() SchemasBuilder {
	return SchemasBuilder{schemas: map[string]Schema{}}
}
