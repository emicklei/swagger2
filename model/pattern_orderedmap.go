package model

// generated with github.com/emicklei/go-templates/orderedmap

import (
	"bytes"
	"encoding/json"
)

// namedPattern associates a name with a Pattern
type namedPattern struct {
	Name    string
	Pattern Pattern
}

// PatternMap encapsulates a list of namedPattern (association) and maintains the insertion order.
type PatternMap struct {
	List []namedPattern
}

// Put adds or replaces a Pattern by its name
func (l *PatternMap) Put(name string, model Pattern) {
	for i, each := range l.List {
		if each.Name == name {
			// replace
			l.List[i] = namedPattern{name, model}
			return
		}
	}
	// add
	l.List = append(l.List, namedPattern{name, model})
}

// At returns a Pattern by its name, ok is false if absent
func (l PatternMap) At(name string) (m Pattern, ok bool) {
	for _, each := range l.List {
		if each.Name == name {
			return each.Pattern, true
		}
	}
	return m, false
}

// Do enumerates all the headers, each with its assigned name
func (l PatternMap) Do(block func(name string, value Pattern)) {
	for _, each := range l.List {
		block(each.Name, each.Pattern)
	}
}

// MarshalJSON writes the PatternMap as if it was a map[string]Pattern
func (l PatternMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, each := range l.List {
		buf.WriteString("\"")
		buf.WriteString(each.Name)
		buf.WriteString("\": ")
		data, err := json.MarshalIndent(each.Pattern, "", "\t")
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

// UnmarshalJSON reads back a PatternMap. This is an expensive operation.
func (l *PatternMap) UnmarshalJSON(data []byte) error {
	raw := map[string]interface{}{}
	json.NewDecoder(bytes.NewReader(data)).Decode(&raw)
	for k, v := range raw {
		// produces JSON bytes for each value
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		var m Pattern
		json.NewDecoder(bytes.NewReader(data)).Decode(&m)
		l.Put(k, m)
	}
	return nil
}
