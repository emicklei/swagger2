package model

import (
	"encoding/json"
	"testing"
)

func TestParameterJSON(t *testing.T) {
	p := Parameter{}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestExternalDocumentObject(t *testing.T) {
	j := `{
  "description": "Find more info here",
  "url": "https://swagger.io"
}`
	var doc ExternalDocumentation
	json.Unmarshal([]byte(j), &doc)
	if got, want := doc.Description, "Find more info here"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if got, want := doc.URL, "https://swagger.io"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInfo(t *testing.T) {
	j := `{
  "title": "Swagger Sample App",
  "description": "This is a sample server Petstore server.",
  "termsOfService": "http://swagger.io/terms/",
  "contact": {
    "name": "API Support",
    "url": "http://www.swagger.io/support",
    "email": "support@swagger.io"
  },
  "license": {
    "name": "Apache 2.0",
    "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
  },
  "version": "1.0.1"
}`
	var doc Info
	json.Unmarshal([]byte(j), &doc)
	if got, want := doc.Contact.Name, "API Support"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if got, want := doc.License.URL, "http://www.apache.org/licenses/LICENSE-2.0.html"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
