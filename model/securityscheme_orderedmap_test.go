package model

import "testing"

func TestSecuritySchemeMap(t *testing.T) {
	l := &SecuritySchemeMap{}
	e := SecurityScheme{}
	l.Put("a SecurityScheme", e)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new(SecuritySchemeMap)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(string(data))
}
