package model

// generated with github.com/emicklei/go-templates/orderedmap

import (
	"bytes"
	"encoding/json"
)

// namedSecurityDefinition associates a name with a SecurityDefinition
type namedSecurityDefinition struct {
	Name   string
	SecurityDefinition SecurityDefinition
}

// SecurityDefinitionMap encapsulates a list of namedSecurityDefinition (association) and maintains the insertion order.
type SecurityDefinitionMap struct {
	List []namedSecurityDefinition
}

// Put adds or replaces a SecurityDefinition by its name
func (l *SecurityDefinitionMap) Put(name string, model SecurityDefinition) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = namedSecurityDefinition{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, namedSecurityDefinition{name, model})
}

// At returns a SecurityDefinition by its name, ok is false if absent
func (l SecurityDefinitionMap) At(name string) (m SecurityDefinition, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.SecurityDefinition, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l SecurityDefinitionMap) Do(block func(name string, value SecurityDefinition)) {
	for _, each := range l.List {
		block(each.Name, each.SecurityDefinition)
	}
}

// MarshalJSON writes the SecurityDefinitionMap as if it was a map[string]SecurityDefinition
func (l SecurityDefinitionMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.SecurityDefinition,"","\t")
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

// UnmarshalJSON reads back a SecurityDefinitionMap. This is an expensive operation.
func (l *SecurityDefinitionMap) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m SecurityDefinition
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
