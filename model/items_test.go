package model

import (
	"encoding/json"
	"testing"
)

func TestItems(t *testing.T) {
	m := Items{}
	m.Items = []Items{Items{}}
	data, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
