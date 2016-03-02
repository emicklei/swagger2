package model

type Object struct {
	Swagger string `json:"swagger"`
}

type Info struct {
	// Required. The title of the application.
	Title string
	// A short description of the application. GFM syntax can be used for rich text representation.
	Description string
	// The Terms of Service for the API.
	TermsOfService string
	// The contact information for the exposed API.
	Contact Contact
	// The license information for the exposed API.
	License License
	// Required Provides the version of the application API (not to be confused with the specification version).
	Version string
}

type Contact struct {
	// The identifying name of the contact person/organization.
	Name string
	// The URL pointing to the contact information. MUST be in the format of a URL.
	URL string
	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string
}

type License struct {
	// Required. The license name used for the API.
	Name string
	// A URL to the license used for the API. MUST be in the format of a URL.
	URL string
}

type number interface{}

type Items struct {
	// The type of the parameter. Since the parameter is not located at the request body, it is limited to simple types (that is, not an object).
	Type string `json:"type,omitempty"`
	// See Data Type Formats for further details.
	Format           string        `json:"format,omitempty"`
	Items            []Items       `json:"items,omitempty"`
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

type XMLObject struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}

// ExternalDocumentation allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Tag allows adding meta data to a single tag
type Tag struct {
	Name         string                `json:"name,omitempty"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs"`
}

type Reference struct {
	Ref string `json:"$ref"`
}
