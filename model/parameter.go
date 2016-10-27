package model

type Parameter struct {
	// The name of the parameter. Parameter names are case sensitive.
	Name string `json:"name"`
	// The location of the parameter. Possible values are "query", "header", "path", "formData" or "body".
	In string `json:"in"`
	// A brief description of the parameter.
	Description string `json:"description,omitempty"`
	// Determines whether this parameter is mandatory.
	Required bool `json:"required"`
	// If in is body
	Schema *Schema `json:"schema"`

	// If not in body uses fields below
	// The type of the parameter. Since the parameter is not located at the request body, it is limited to simple types (that is, not an object).
	Type string `json:"type,omitempty"`
	// See Data Type Formats for further details.
	Format string `json:"format,omitempty"`
	// Sets the ability to pass empty-valued parameters.
	AllowEmptyValue bool
	// Required if type is "array". Describes the type of items in the array.
	//Items []Item
	// Determines the format of the array if type array is used.
	CollectionFormat string        `json:"collectionFormat,omitempty"`
	Default          interface{}   `json:"default,omitempty"`
	Maximum          number        `json:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          number        `json:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        int           `json:"maxLength"`
	MinLength        int           `json:"minLength"`
	Pattern          string        `json:"pattern,omitempty"`
	MaxItems         int           `json:"maxItems"`
	MinItems         int           `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum,omitempty"`
	MultipleOf       number        `json:"multipleOf,omitempty"`
	// Extensions
	// TODO
}
