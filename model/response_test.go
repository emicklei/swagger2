package model

import "testing"

func TestResponses(t *testing.T) {
	m := ResponseMap{}
	r0 := Response{Description: "Unexpected error", Schema: Schema{Ref: "#/definitions/ErrorModel"}}
	m.Put("default", r0)
	r1 := Response{Description: "a pet to be returned", Schema: Schema{Ref: "#/definitions/Pet"}}
	m.Put("200", r1)
	data, err := m.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
