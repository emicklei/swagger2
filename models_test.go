package swagger

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
