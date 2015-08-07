package swagger

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

type Paths struct {
	Path string
	// Extensions
}

// TODO custom marshal of Paths

type PathItem struct {
	// Allows for an external definition of this path item.
	// The referenced structure MUST be in the format of a Path Item Object.
	// If there are conflicts between the referenced definition and this Path Item's definition,
	// the behavior is undefined.
	Ref string
	// A definition of a GET operation on this path.
	Get Operation
	// A definition of a PUT operation on this path.
	Put Operation
	// A definition of a POST operation on this path.
	Post Operation
	// A definition of a DELETE operation on this path.
	Delete Operation
	// A definition of a OPTIONS operation on this path.
	Options Operation
	// A definition of a HEAD operation on this path.
	Head Operation
	// A definition of a PATCH operation on this path.
	Patch Operation
	// Extensions
}

type Operation struct {
	// A list of tags for API documentation control.
	// Tags can be used for logical grouping of operations by resources or any other qualifier.
	Tags []string
	// A short summary of what the operation does.
	// For maximum readability in the swagger-ui, this field SHOULD be less than 120 characters.
	Summary string
	// A verbose explanation of the operation behavior. GFM syntax can be used for rich text representation.
	Description string
}
