package model

type Response struct {
	Description string                 `json:"description"`
	Schema      Schema                 `json:"schema"`
	Headers     HeaderMap              `json:"headers"`
	Examples    map[string]interface{} `json:"example"`
	Extensions  interface{}            // TODO
}

func (m ResponseMap) Default(r Response) {
	m.Put("default", r)
}
