//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	srcFile = "appid.go"
	dstFile = "names.go"
)

func main() {
	// set parser position
	fset := token.NewFileSet()

	// parse appid.go
	node, err := parser.ParseFile(fset, srcFile, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Cant parse %s: %v", srcFile, err)
	}

	var buf bytes.Buffer
	buf.WriteString(`// Code generated by "generate.go"; DO NOT EDIT.
package appid

// AppName contains name for each AppID.
var AppName = map[AppID]string{
`)

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for i, name := range valSpec.Names {
				if len(valSpec.Values) <= i {
					continue
				}

				// Get comment
				comment := ""
				if valSpec.Comment != nil && len(valSpec.Comment.List) > i {
					comment = valSpec.Comment.List[i].Text
					comment = strings.TrimPrefix(comment, "//")
					comment = strings.TrimSpace(comment)
				} else {
					comment = splitCamelCase(name.Name)
				}

				buf.WriteString(fmt.Sprintf("\t%s: \"%s\",\n", name.Name, comment))
			}
		}
	}

	buf.WriteString("}")

	// Format code
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Formatting error: %v", err)
	}

	// write
	err = os.WriteFile(dstFile, formatted, 0644)
	if err != nil {
		log.Fatalf("Fail write %s: %v", dstFile, err)
	}

	fmt.Printf("Success generated %s.\n", dstFile)
}

func splitCamelCase(s string) string {
	var words []string
	runes := []rune(s)

	start := 0
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) && (i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}
	words = append(words, string(runes[start:]))

	return strings.Join(words, " ")
}