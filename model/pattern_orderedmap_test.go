package model

import "testing"

func TestPatternMap(t *testing.T) {
	l := &PatternMap{}
	e := "hier"
	l.Put("a Pattern", e)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new(PatternMap)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 1; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(string(data))
}
