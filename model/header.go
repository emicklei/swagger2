package model

type Header struct {
	Description      string        `json:"description,omitempty"`
	Type             string        `json:"type,omitempty"`
	Format           string        `json:"format,omitempty"`
	Items            []interface{} `json:"items,omitempty"`
	CollectionFormat string        `json:"collectionFormat,omitempty"`
	Default          string        `json:"default,omitempty"`
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
	Enum             []interface{} `json:"enum,omitempty"`
	MultipleOf       string        `json:"multipleOf,omitempty"`
}
