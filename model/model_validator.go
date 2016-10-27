package model

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidatorMapper func(ValidationFields, string, reflect.Type)

func nopValidator(_ ValidationFields, _ string, _ reflect.Type) {
}

func minMaxLengthValidator(item ValidationFields, validator string, t reflect.Type) {
	re := regexp.MustCompile(`length\(([0-9]+)\|([0-9]+)\)`)
	groups := re.FindStringSubmatch(validator)
	min, _ := strconv.Atoi(groups[1])
	max, _ := strconv.Atoi(groups[2])
	if t.Kind() == reflect.String {
		item.SetMinLength(min)
		item.SetMaxLength(max)
	}
}

func requiredValidator(item ValidationFields, validator string, t reflect.Type) {
	if validator == "required" {
		item.SetRequired(true)
	} else {
		item.SetRequired(false)
	}
}

func rangeValidator(item ValidationFields, validator string, t reflect.Type) {
	re := regexp.MustCompile(`range\(([0-9]+)\|([0-9]+)\)`)
	groups := re.FindStringSubmatch(validator)
	min, _ := strconv.Atoi(groups[1])
	max, _ := strconv.Atoi(groups[2])

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.String:
		item.SetMinLength(min)
		item.SetMaxLength(max)
	case reflect.Slice, reflect.Array:
		item.SetMinItems(min)
		item.SetMaxItems(max)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		item.SetMinimum(min)
		item.SetMaximum(max)
	default:
	}
}

var (
	validatorRegistry = map[string]ValidatorMapper{
		"required":       requiredValidator,
		"range":          rangeValidator,
		"length":         minMaxLengthValidator,
		"matches":        nopValidator,
		"alpha":          nopValidator,
		"alphanum":       nopValidator,
		"ascii":          nopValidator,
		"base64":         nopValidator,
		"creditcard":     nopValidator,
		"datauri":        nopValidator,
		"dialstring":     nopValidator,
		"dns":            nopValidator,
		"email":          nopValidator,
		"float":          nopValidator,
		"fullwidth":      nopValidator,
		"halfwidth":      nopValidator,
		"hexadecimal":    nopValidator,
		"hexcolor":       nopValidator,
		"host":           nopValidator,
		"int":            nopValidator,
		"ip":             nopValidator,
		"ipv4":           nopValidator,
		"ipv6":           nopValidator,
		"isbn10":         nopValidator,
		"isbn13":         nopValidator,
		"json":           nopValidator,
		"latitude":       nopValidator,
		"longitude":      nopValidator,
		"lowercase":      nopValidator,
		"mac":            nopValidator,
		"multibyte":      nopValidator,
		"null":           nopValidator,
		"numeric":        nopValidator,
		"port":           nopValidator,
		"printableascii": nopValidator,
		"requri":         nopValidator,
		"requrl":         nopValidator,
		"rgbcolor":       nopValidator,
		"ssn":            nopValidator,
		"semver":         nopValidator,
		"uppercase":      nopValidator,
		"url":            nopValidator,
		"utfdigit":       nopValidator,
		"utfletter":      nopValidator,
		"utfletternum":   nopValidator,
		"utfnumeric":     nopValidator,
		"uuid":           nopValidator,
		"uuidv3":         nopValidator,
		"uuidv4":         nopValidator,
		"uuidv5":         nopValidator,
	}
)

type Token struct {
	value string
	err   error
}

func tokenize(tagValue string) chan *Token {
	c := make(chan *Token)
	go func() {
		for tagValue != "" {
			if strings.HasPrefix(tagValue, "matches(") {
				index := strings.Index(tagValue, ")")
				if index == -1 {
					c <- &Token{"", errors.New("Invalid sequence")}
				}
				c <- &Token{tagValue[:index+1], nil}
				if len(tagValue)-1 == index+1 {
					break
				}
				tagValue = tagValue[index+2:]
			} else {
				index := strings.Index(tagValue, ",")
				if index == -1 {
					c <- &Token{tagValue[:], nil}
					break
				}
				c <- &Token{tagValue[:index], nil}
				tagValue = tagValue[index+1:]
			}
		}
		close(c)
	}()
	return c
}

func MapToGoValidator(validationFields ValidationFields, valid string, reflectType reflect.Type) error {
	c := tokenize(valid)

	for token := range c {
		if token.err != nil {
			return token.err
		}
		val, ok := validatorRegistry[strings.Split(token.value, "(")[0]]
		if ok {
			val(validationFields, token.value, reflectType)
		}
	}
	return nil
}

func newBoolPtr(val bool) *bool {
	return &val
}
