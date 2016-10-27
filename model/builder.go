package model

// idea is to generate builders for all types to allow fluent programming style to create a full Swagger object.

type Swagger_Builder struct{ n *Swagger }

func NewSwagger() *Swagger_Builder                             {
	builder := &Swagger_Builder{new(Swagger)}
	builder.n.Paths = map[string]*PathItem{}
	builder.n.Definitions = map[string]*Schema{}
	builder.n.Swagger = "2.0"
	return builder
}
func (b *Swagger_Builder) Info(n Info_Builder) *Swagger_Builder { b.n.Info = n.Build(); return b }
func (b *Swagger_Builder) Host(s string) *Swagger_Builder       { b.n.Host = s; return b }
func (b *Swagger_Builder) Build() *Swagger                      { return b.n }
func (b *Swagger_Builder) Path(path string, pathItem *PathItem) *Swagger_Builder {
	b.n.Paths[path] = pathItem
	return b
}
func (b *Swagger_Builder) Definitions(definitions map[string]*Schema) *Swagger_Builder {
	for key, value := range definitions {
		b.n.Definitions[key] = value
	}
	return b
}

type Info_Builder struct{ n *Info }

func NewInfo() Info_Builder                        {
	builder := Info_Builder{new(Info)}
	builder.n.Version = "1.0.0"
	return builder
}
func (b Info_Builder) Title(s string) Info_Builder { b.n.Title = s; return b }
func (b Info_Builder) Build() Info                 { return *b.n }
