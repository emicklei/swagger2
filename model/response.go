package model

// Responses is an object to hold responses to be reused across operations.
type Responses struct {
	Default Response `json:"default"`
	// A single response definition, mapping a "name" to the response it defines.
	ResponsesMap map[string]Response
}

type Response struct {
	Description string                 `json:"description"`
	Schema      Schema                 `json:"schema"`
	Headers     map[string]Header      `json:"headers"`
	Examples    map[string]interface{} `json:"example"`
}
