package model

import "testing"

func TestSecurityDefinitionMap(t *testing.T) {
	l := &SecurityDefinitionMap{}
	e := SecurityDefinition{}
	l.Put("a SecurityDefinition", e)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new(SecurityDefinitionMap)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(string(data))
}
