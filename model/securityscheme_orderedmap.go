package model

// generated with github.com/emicklei/go-templates/orderedmap

import (
	"bytes"
	"encoding/json"
)

// namedSecurityScheme associates a name with a SecurityScheme
type namedSecurityScheme struct {
	Name   string
	SecurityScheme SecurityScheme
}

// SecuritySchemeMap encapsulates a list of namedSecurityScheme (association) and maintains the insertion order.
type SecuritySchemeMap struct {
	List []namedSecurityScheme
}

// Put adds or replaces a SecurityScheme by its name
func (l *SecuritySchemeMap) Put(name string, model SecurityScheme) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = namedSecurityScheme{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, namedSecurityScheme{name, model})
}

// At returns a SecurityScheme by its name, ok is false if absent
func (l SecuritySchemeMap) At(name string) (m SecurityScheme, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.SecurityScheme, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l SecuritySchemeMap) Do(block func(name string, value SecurityScheme)) {
	for _, each := range l.List {
		block(each.Name, each.SecurityScheme)
	}
}

// MarshalJSON writes the SecuritySchemeMap as if it was a map[string]SecurityScheme
func (l SecuritySchemeMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.SecurityScheme,"","\t")
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

// UnmarshalJSON reads back a SecuritySchemeMap. This is an expensive operation.
func (l *SecuritySchemeMap) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m SecurityScheme
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
