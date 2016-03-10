package swagger2

import (
	"encoding/json"
	"net"
	"testing"
)

type fakeint int

type Annoted struct {
	Name    string  `description:"name" modelDescription:"a test"`
	Happy   bool    `json:"happy"`
	Stati   string  `enum:"off|on" default:"on" modelDescription:"more description"`
	ID      string  `unique:"true"`
	FakeInt fakeint `type:"integer"`
	IP      net.IP  `type:"string"`
}

func TestSchemaPrimitives(t *testing.T) {
	for _, each := range []struct {
		value  interface{}
		result string
	}{
		{"string", `{"type":"string"}`},
		{42, `{"type":"integer"}`},
		{int8(42), `{"type":"integer"}`},
		{int16(42), `{"type":"integer"}`},
		{uint8(42), `{"type":"integer"}`},
		{uint16(42), `{"type":"integer"}`},
		{int32(42), `{"type":"integer"}`},
		{int64(42), `{"type":"integer"}`},
		{uint32(42), `{"type":"integer"}`},
		{uint64(42), `{"type":"integer"}`},
		{false, `{"type":"bool"}`},
		{nil, `null`},
	} {
		b := NewSchemaBuilder()
		{
			got := doc(b.Build(each.value))
			want := each.result
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}
}

func doc(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
