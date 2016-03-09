package swagger

import "github.com/emicklei/swagger2/model"

type SchemasBuilder struct {
	schemas map[string]model.Schema
}

func NewSchemasBuilder() SchemasBuilder {
	return SchemasBuilder{schemas: map[string]model.Schema{}}
}
