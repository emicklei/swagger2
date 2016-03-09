package model

import "testing"

func TestResponseMap(t *testing.T) {
	l := &ResponseMap{}
	e := Response{}
	l.Put("a Response", e)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new(ResponseMap)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(string(data))
}
