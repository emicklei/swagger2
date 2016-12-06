package model

// generated with github.com/emicklei/go-templates/orderedmap

import (
	"bytes"
	"encoding/json"
)

// namedResponse associates a name with a Response
type namedResponse struct {
	Name     string
	Response Response
}

// ResponseMap encapsulates a list of namedResponse (association) and maintains the insertion order.
type ResponseMap struct {
	List []namedResponse
}

// Put adds or replaces a Response by its name
func (l *ResponseMap) Put(name string, model Response) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = namedResponse{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, namedResponse{name, model})
}

// At returns a Response by its name, ok is false if absent
func (l ResponseMap) At(name string) (m Response, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.Response, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l ResponseMap) Do(block func(name string, value Response)) {
	for _, each := range l.List {
		block(each.Name, each.Response)
	}
}

// MarshalJSON writes the ResponseMap as if it was a map[string]Response
func (l ResponseMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.Response, "", "\t")
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

// UnmarshalJSON reads back a ResponseMap. This is an expensive operation.
func (l *ResponseMap) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m Response
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
