package model

type Schema struct {
	Ref              string        `json:"$ref,omitempty"`
	Format           string        `json:"format,omitempty"`
	Title            string        `json:"title,omitempty"`
	Description      string        `json:"description,omitempty"`
	Default          string        `json:"default,omitempty"`
	MultipleOf       string        `json:"multipleOf,omitempty"`
	Maximum          int           `json:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          number        `json:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int           `json:"maxLength"`
	MinLength        int           `json:"minLength"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         int           `json:"maxItems"`
	MinItems         int           `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	MaxProperties    int           `json:"maxProperties"`
	MinProperties    int           `json:"minProperties"`
	Required         bool          `json:"required"`
	Enum             []interface{} `json:"enum,omitempty"`
	Type             string        `json:"type,omitempty"`
	// definitions were adjusted to the Swagger
	Items                []interface{} `json:"items,omitempty"`
	AllOf                []interface{} `json:"allOf,omitempty"`
	Properties           []interface{} `json:"properties,omitempty"`
	AdditionalProperties []interface{} `json:"additionalProperties,omitempty"`
	//  further schema documentation
	Discriminator string                `json:"discriminator,omitempty"`
	ReadOnly      bool                  `json:"readOnly"`
	XML           XMLObject             `json:"xml"`
	ExternalDocs  ExternalDocumentation `json:"externalDocs"`
	Example       interface{}           `json:"example"`
}
