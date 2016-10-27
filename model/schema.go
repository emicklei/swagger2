package model

type ValidationFields interface {
	SetMaximum(int)
	SetMinimum(int)
	SetRequired(bool)
	SetMaxLength(int)
	SetMinLength(int)
	SetMaxItems(int)
	SetMinItems(int)
}

type Schema struct {
	Ref              string        `json:"$ref,omitempty"`
	Format           string        `json:"format,omitempty"`
	Title            string        `json:"title,omitempty"`
	Description      string        `json:"description,omitempty"`
	Default          string        `json:"default,omitempty"`
	MultipleOf       string        `json:"multipleOf,omitempty"`
	Maximum          *int          `json:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty"`
	Minimum          *int          `json:"minimum,omitempty"`
	ExclusiveMinimum *bool         `json:"exclusiveMinimum,omitempty"`
	MaxLength        *int          `json:"maxLength,omitempty"`
	MinLength        *int          `json:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         *int          `json:"maxItems,omitempty"`
	MinItems         *int          `json:"minItems,omitempty"`
	UniqueItems      *bool         `json:"uniqueItems,omitempty"`
	MaxProperties    *int          `json:"maxProperties,omitempty"`
	MinProperties    *int          `json:"minProperties,omitempty"`
	Required         *bool         `json:"required,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	Type             string        `json:"type,omitempty"`
	// definitions were adjusted to the Swagger
	Items                []*Schema          `json:"items,omitempty"`
	AllOf                []*Schema          `json:"allOf,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema            `json:"additionalProperties,omitempty"`
	//PatternProperties    PatternMap `json:"patternProperties,omitempty"`
	//  further schema documentation
	Discriminator string                 `json:"discriminator,omitempty"`
	ReadOnly      *bool                  `json:"readOnly,omitempty"`
	XML           *XMLObject             `json:"xml,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	Example       interface{}            `json:"example,omitempty"`
}

func (s *Schema) SetMaximum(m int) {
	s.Maximum = &m
}

func (s *Schema) SetMinimum(m int) {
	s.Minimum = &m
}

func (s *Schema) SetMaxLength(m int) {
	s.MaxLength = &m
}

func (s *Schema) SetMinLength(m int) {
	s.MinLength = &m
}

func (s *Schema) SetMaxItems(m int) {
	s.MaxItems = &m
}

func (s *Schema) SetMinItems(m int) {
	s.MinItems = &m
}

func (s *Schema) SetRequired(b bool) {
	s.Required = &b
}
