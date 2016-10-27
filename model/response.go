package model

type Response struct {
	Description string                 `json:"description"`
	Schema      *Schema                 `json:"schema,omitempty"`
	Headers     *HeaderMap              `json:"headers,omitempty"`
	Examples    map[string]interface{} `json:"example,omitempty"`
	//Extensions  interface{}            // TODO
}

func (m ResponseMap) Default(r Response) {
	m.Put("default", r)
}
