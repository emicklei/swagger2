package model

import (
	"testing"
)

func TestMinMaxValidator(t *testing.T) {
	item := Schema{}
	minMaxLengthValidator(&item, "length(12|14)", nil)
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
	err := MapToGoValidator(&prop, "length(2|4)", nil)
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
