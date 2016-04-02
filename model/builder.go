package model

// idea is to generate builders for all types to allow fluent programming style to create a full Swagger object.

type Swagger_Builder struct{ n *Swagger }

func NewSwagger() Swagger_Builder                             { return Swagger_Builder{new(Swagger)} }
func (b Swagger_Builder) Info(n Info_Builder) Swagger_Builder { b.n.Info = n.Build(); return b }
func (b Swagger_Builder) Host(s string) Swagger_Builder       { b.n.Host = s; return b }
func (b Swagger_Builder) Build() Swagger                      { return *b.n }

type Info_Builder struct{ n *Info }

func NewInfo() Info_Builder                        { return Info_Builder{new(Info)} }
func (b Info_Builder) Title(s string) Info_Builder { b.n.Title = s; return b }
func (b Info_Builder) Build() Info                 { return *b.n }
