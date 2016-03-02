package model

// SecurityDefinition is a declaration of the security schemes available to be used in the specification.
type SecurityDefinition struct {
	// A single security scheme definition, mapping a "name" to the scheme it defines.
	Schemes map[string]SecurityScheme
}

// SecurityScheme allows the definition of a security scheme that can be used by the operations.
type SecurityScheme struct {
	//  The type of the security scheme. Valid values are "basic", "apiKey" or "oauth2".
	Type string `json:"type,omitempty"`
	// A short description for security scheme.
	Description string `json:"description,omitempty"`
	// The name of the header or query parameter to be used.
	Name string `json:"name,omitempty"`
	// The location of the API key. Valid values are "query" or "header".
	In string `json:"in,omitempty"`
	// The flow used by the OAuth2 security scheme. Valid values are "implicit", "password", "application" or "accessCode".
	Flow string `json:"flow,omitempty"`
	// The authorization URL to be used for this flow. This SHOULD be in the form of a URL.
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	// The token URL to be used for this flow. This SHOULD be in the form of a URL.
	TokenUrl string `json:"tokenUrl,omitempty"`
	// Maps between a name of a scope to a short description of it (as the value of the property).
	Scopes map[string]string `json:"scopes,omitempty"`
}

// SecurityDefinitionList encapsulates a list of SecurityDefinition (association)
type SecurityDefinitionList struct {
	List []SecurityDefinition
}
