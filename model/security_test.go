package model

import (
	"encoding/json"
	"testing"
)

func TestSecurityScheme(t *testing.T) {
	j := `{
  "type": "oauth2",
  "authorizationUrl": "http://swagger.io/api/oauth/dialog",
  "flow": "implicit",
  "scopes": {
    "write:pets": "modify pets in your account",
    "read:pets": "read your pets"
  }
}`
	var doc SecurityScheme
	json.Unmarshal([]byte(j), &doc)
	if got, want := doc.Type, "oauth2"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if got, want := doc.Scopes["write:pets"], "modify pets in your account"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
