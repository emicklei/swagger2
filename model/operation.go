package model

// Operation describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control.
	// Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags []string `json:"summary,tags"`
	// A short summary of what the operation does.
	// For maximum readability in the swagger-ui, this field SHOULD be less than 120 characters.
	Summary string `json:"summary,omitempty"`
	// A verbose explanation of the operation behavior. GFM syntax can be used for rich text representation.
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationId  string                `json:"operationId,omitempty"`
	Consumes     []string              `json:"consumes,omitempty"`
	Produces     []string              `json:"produces,omitempty"`
	Parameters   []Parameter           `json:"parameters,omitempty"`
	Responses    ResponseMap           `json:"responses,omitempty"`
	Schemes      []string              `json:"schemes,omitempty"`
	Deprecated   bool                  `json:"deprecated"`
	Security     SecurityDefinitionMap `json:"security"`
}
