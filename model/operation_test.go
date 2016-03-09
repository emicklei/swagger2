package model

import (
	"encoding/json"
	"testing"
)

func TestOperation(t *testing.T) {
	m := Operation{}

	m.Tags = []string{"pet"}
	m.Summary = "Updates a pet in the store with form data"
	m.OperationId = "updatePetWithForm"
	m.Consumes = []string{"application/x-www-form-urlencoded"}
	m.Produces = []string{"application/json", "application/xml"}
	m.Parameters = []Parameter{
		Parameter{
			Name:        "petId",
			In:          "path",
			Description: "ID of pet that needs to be updated",
			Required:    true,
			Type:        "string",
		},
	}
	d := Response{
		Description: "Pet updated",
	}
	i := Response{
		Description: "Invalid input",
	}
	m.Responses.Default(d)
	m.Responses.Put("400", i)

	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
