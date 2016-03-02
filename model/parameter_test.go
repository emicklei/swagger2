package model

import (
	"encoding/json"
	"testing"
)

func TestBodyParameter(t *testing.T) {
	j := `{
  "name": "token",
  "in": "header",
  "description": "token to be passed as a header",
  "required": true,
  "type": "array",
  "items": {
    "type": "integer",
    "format": "int64"
  },
  "collectionFormat": "csv"
}`
	var doc Parameter
	json.Unmarshal([]byte(j), &doc)
	if got, want := doc.Name, "token"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if got, want := doc.CollectionFormat, "csv"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
