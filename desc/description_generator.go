package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	sourceFile := flag.String("in", os.Getenv("GOFILE"), "golang file containing strucs for swagger")
	targetFile := flag.String("out", "", "target file where description should be written")
	flag.Parse()

	if *sourceFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *targetFile == "" {
		if strings.HasSuffix(*sourceFile, "_test.go") {
			*targetFile = strings.TrimSuffix(*sourceFile, "_test.go") + "_swagger_generated_test.go"
		} else {
			*targetFile = strings.TrimSuffix(*sourceFile, ".go") + "_swagger_generated.go"
		}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *sourceFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	c := Parse(f)

	var file *os.File
	if *targetFile == "-" {
		file = os.Stdout
	} else {
		file, err = os.Create(*targetFile)
		defer file.Close()
		if err != nil {
			panic(err)
		}
	}
	for x := range c {
		fmt.Fprintln(file, x)
	}
}

func Parse(f *ast.File) chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		re := regexp.MustCompile(`^([^,]+).*`)
		firstStruct := true
		for _, decl := range f.Decls {

			gendecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if len(gendecl.Specs) == 0 {
				continue
			}
			typeSpec, ok := gendecl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}
			structDecl, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			if firstStruct {
				firstStruct = false
				c <- "package " + f.Name.Name
			}
			c <- ""
			c <- fmt.Sprintf("func (%s) SwaggerDoc() map[string]string {", typeSpec.Name)

			type Field struct {
				Field, Desc string
			}
			structDoc := filterDoc(gendecl.Doc.Text())

			fields := []*Field{}
			offsetLength := 0
			for _, field := range structDecl.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				fieldName := field.Names[0].Name
				if field.Tag != nil {
					tag := reflect.StructTag(field.Tag.Value)
					jsonTag := tag.Get("`json")
					if len(jsonTag) > 0 {
						matches := re.FindStringSubmatch(jsonTag)
						if len(matches) == 2 && len(matches[1]) > 0 {
							fieldName = matches[1]
						}
					}
				}

				docText := filterDoc(field.Doc.Text())
				if len(docText) > 0 {
					fields = append(fields, &Field{fieldName, strconv.Quote(docText)})
					fieldLength := utf8.RuneCountInString(fieldName)
					if fieldLength > offsetLength {
						offsetLength = fieldLength
					}
				}
			}

			if len(structDoc) != 0 || len(fields) != 0 {
				c <- "\treturn map[string]string{"
				if len(structDoc) > 0 {
					c <- "\t\t" + "\"\": " + strings.Repeat(" ", offsetLength) + strconv.Quote(structDoc) + ","
				}
				if len(fields) > 0 {
					for _, field := range fields {
						c <- "\t\t" + "\"" + field.Field + "\": " + strings.Repeat(" ", offsetLength-len(field.Field)) + field.Desc + ","
					}
				}
				c <- "\t}"
			} else {
				c <- "\treturn map[string]string{}"
			}
			c <- "}"
		}
	}()
	return c
}

func filterDoc(doc string) string {
	buf := ""
	scanner := bufio.NewScanner(strings.NewReader(doc))
	for scanner.Scan() {
		token := scanner.Text()
		trimmed := strings.TrimSpace(token)
		if strings.HasPrefix(trimmed, "TODO") {
			continue
		}
		if strings.HasPrefix(trimmed, "---") {
			break
		}
		buf = buf + scanner.Text() + "\n"
	}
	return strings.TrimSpace(buf)
}
