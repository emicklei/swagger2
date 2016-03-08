package model

import "testing"

func TestHeaderList(t *testing.T) {
	l := &HeaderList{}
	h1 := Header{Description: "HTTP standard Authorization"}
	l.Put("Authorization", h1)
	h2 := Header{Description: "HTTP standard Content-Type"}
	l.Put("Content-Type", h2)
	data, err := l.MarshalJSON()
	if err != nil {
		t.Fatal("marshal failed", err)
	}
	back := new(HeaderList)
	err = back.UnmarshalJSON(data)
	if err != nil {
		t.Fatal("unmarshal failed", err)
	}
	if got, want := len(back.List), 2; got != want {
		t.Errorf("missing header(s), got %v want %v", got, want)
	}
	t.Log(string(data))
}
