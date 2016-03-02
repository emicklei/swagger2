package model

type Paths struct {
	Path string
	// Extensions
	// TODO
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
	// A list of parameters that are applicable for all the operations described under this path
	Parameters []Parameter
}
