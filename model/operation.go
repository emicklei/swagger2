package model

// Operation describes a single API operation on a path.
type Operation struct {
	// A list of tags for API documentation control.
	// Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags []string
	// A short summary of what the operation does.
	// For maximum readability in the swagger-ui, this field SHOULD be less than 120 characters.
	Summary string
	// A verbose explanation of the operation behavior. GFM syntax can be used for rich text representation.
	Description  string
	ExternalDocs ExternalDocumentation
	OperationId  string                 `json:"operationId"`
	Consumes     []string               `json:"consumes"`
	Produces     []string               `json:"produces"`
	Parameters   []Parameter            `json:"parameters"`
	Responses    []Responses            `json:"responses"`
	Schemes      []string               `json:"schemes"`
	Deprecated   bool                   `json:"deprecated"`
	Security     SecurityDefinitionList `json:"security"`
}
