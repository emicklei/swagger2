# swagger2
Go implementation of the Swagger 2.0 specification (https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md)

## Work in progress

[![Build Status](https://travis-ci.org/emicklei/swagger2.png)](https://travis-ci.org/emicklei/swagger2)

This will include a ModelBuilder implementation such as the one in github.com/emicklei/go-restful/swagger

## Generating descriptions from comments

Compile the tool for comment extraction:

```bash
cd desc
go build
```

Generate the documentation map for structs in a go file and print it to stdout:

```bash
./desc -in description_generator_test.go -out -
```
This will generate the following output:
```golang
package main

func (Address) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Address doc",
		"country": "Country doc",
		"postcode": "PostCode doc",
	}
}

func (Person) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Person doc",
		"firstName": "FirstName doc",
		"LastName": "LastName doc",
		"middleName": "MiddleName doc",
		"HeightInPounds": "Field without tag",
	}
}
```

The original structs look like this:

```
//Address doc
type Address struct {
	//Country doc
	Country string `json:"country,omitempty"`
	//PostCode doc
	PostCode int `json:"postcode,omitempty"`
}

//Person doc
type Person struct {
	//FirstName doc
	FirstName string `json:"firstName,omitempty"`
	//LastName doc
	LastName string `json:",omitempty"`
	//MiddleName doc
	MiddleName string `json:"middleName"`

	//Field without tag
	HeightInPounds int
}
```

If `-out` is omitted, the full path and name of the input file is taken and a
`_swagger_generated.go` prefix will be added to it.  So if the input file was
`~/wherever/test.go`, the resulting output file will be
`~/wherever/test_swagger_generated.go`.

(c) 2015-2016, http://ernestmicklei.com. MIT License
