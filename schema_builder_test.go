package swagger2

import (
	"encoding/json"
	"github.com/emicklei/swagger2/model"
	"net"
	"reflect"
	"testing"
)

type fakeint int

type Recursive struct {
	Rec    *Recursive
	RecInt int
}

type Annoted struct {
	Name    string  `description:"name" modelDescription:"a test"`
	Happy   bool    `json:"happy"`
	Stati   string  `enum:"off|on" default:"on" modelDescription:"more description" valid:"range(10|20)"`
	ID      string  `unique:"true"`
	FakeInt fakeint `type:"integer" valid:"required,range(3|4)"`
	IP      net.IP  `type:"string"`
	Rec     Recursive `valid:"required"`
	IgnoreMe string `json:"-"`
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
			schema, _ := b.Build(each.value)
			got := doc(schema)
			want := each.result
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		}
	}
}

func TestAnnotedModel(t *testing.T) {
	expectedSchemasJson := `{
	  "Annoted": {
	    "type": "object",
	    "properties": {
	      "FakeInt": {
	        "type": "integer",
	        "minimum": 3,
	        "maximum": 4,
	        "required": true
	      },
	      "ID": {
	        "type": "string"
	      },
	      "IP": {
	        "type": "array"
	      },
	      "Name": {
	        "type": "string"
	      },
	      "Rec": {
	        "$ref": "#/definitions/Recursive",
	        "required": true
	      },
	      "Stati": {
	        "type": "string",
	        "minLength": 10,
	        "maxLength": 20
	      },
	      "happy": {
	        "type": "bool"
	      }
	    }
	  },
	  "Recursive": {
	    "type": "object",
	    "properties": {
	      "Rec": {
	        "$ref": "#/definitions/Recursive"
	      },
	      "RecInt": {
	        "type": "integer"
	      }
	    }
	  }
	}
	`

	val := Annoted{}
	b := NewSchemaBuilder()

	ref, schemas := b.Build(val)

	expectedRef := &model.Schema{}
	expectedSchemas := map[string]*model.Schema{}
	json.Unmarshal([]byte(`{"$ref":"#/definitions/Annoted"}`), expectedRef)
	json.Unmarshal([]byte(expectedSchemasJson), &expectedSchemas)
	if !reflect.DeepEqual(ref, expectedRef) {
		t.Fail()
	}
	if !reflect.DeepEqual(schemas, expectedSchemas) {
		t.Errorf("got %v want %v", doc(schemas), doc(expectedSchemas))
		t.Fail()
	}
}

func doc(schema interface{}) string {
	data, _ := json.Marshal(schema)
	return string(data)
}
