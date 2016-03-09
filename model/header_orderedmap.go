package model

// generated with github.com/emicklei/go-templates/orderedmap

import (
	"bytes"
	"encoding/json"
)

// namedHeader associates a name with a Header
type namedHeader struct {
	Name   string
	Header Header
}

// HeaderMap encapsulates a list of namedHeader (association) and maintains the insertion order.
type HeaderMap struct {
	List []namedHeader
}

// Put adds or replaces a Header by its name
func (l *HeaderMap) Put(name string, model Header) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = namedHeader{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, namedHeader{name, model})
}

// At returns a Header by its name, ok is false if absent
func (l HeaderMap) At(name string) (m Header, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.Header, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l HeaderMap) Do(block func(name string, value Header)) {
	for _, each := range l.List {
		block(each.Name, each.Header)
	}
}

// MarshalJSON writes the HeaderMap as if it was a map[string]Header
func (l HeaderMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.Header,"","\t")
		if err != nil {
			return buf.Bytes(), err
		}
		buf.Write(data)
		if i < len(l.List)-1 {
			buf.WriteString(",\n")
		}
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

// UnmarshalJSON reads back a HeaderMap. This is an expensive operation.
func (l *HeaderMap) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m Header
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
