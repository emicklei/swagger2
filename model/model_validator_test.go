package model

import (
	"testing"
	"reflect"
)

func TestMinMaxValidator(t *testing.T) {
	item := Schema{}
	str := ""
	minMaxLengthValidator(nil, &item, "length(12|14)", reflect.TypeOf(str), "")
	if *item.MinLength != 12 {
		t.Fail()
	}
	if *item.MaxLength != 14 {
		t.Fail()
	}
}

func TestTokenizer(t *testing.T) {
	input := "length(1|2),matches(.*,),alnum,matches(.*,),1234"
	expected := []string{
		"length(1|2)",
		"matches(.*,)",
		"alnum",
		"matches(.*,)",
		"1234",
	}
	tokens := tokenize(input)
	for _, exp := range expected {
		if exp != (<-tokens).value {
			t.Fail()
		}
	}
}

func TestMinMaxValidatorMapping(t *testing.T) {
	prop := Schema{}
	str := ""
	err := MapToGoValidator(nil, &prop, "length(2|4)", reflect.TypeOf(str), "")
	if err != nil {
		t.Fail()
	}
	if *prop.MaxLength != 4 {
		t.Fail()
	}
	if *prop.MinLength != 2 {
		t.Fail()
	}
}
