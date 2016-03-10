package model

type Object struct {
	Swagger string `json:"swagger"`
}

type Info struct {
	// Required. The title of the application.
	Title string `json:"title,omitempty"`
	// A short description of the application. GFM syntax can be used for rich text representation.
	Description string `json:"description,omitempty"`
	// The Terms of Service for the API.
	TermsOfService string `json:"termsOfService,omitempty"`
	// The contact information for the exposed API.
	Contact Contact `json:"contact"`
	// The license information for the exposed API.
	License License `json:"license"`
	// Required Provides the version of the application API (not to be confused with the specification version).
	Version string `json:"version,omitempty"`
}

type Contact struct {
	// The identifying name of the contact person/organization.
	Name string `json:"name,omitempty"`
	// The URL pointing to the contact information. MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
	// The email address of the contact person/organization. MUST be in the format of an email address.
	Email string `json:"email,omitempty"`
}

type License struct {
	// Required. The license name used for the API.
	Name string `json:"name,omitempty"`
	// A URL to the license used for the API. MUST be in the format of a URL.
	URL string `json:"url,omitempty"`
}

type number interface{}

type XMLObject struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute *bool  `json:"attribute,omitempty"`
	Wrapped   *bool  `json:"wrapped,omitempty"`
}

// ExternalDocumentation allows referencing an external resource for extended documentation.
type ExternalDocumentation struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
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
