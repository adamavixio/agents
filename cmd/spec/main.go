package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func main() {
	// Create a new file set
	fset := token.NewFileSet()

	// Create the AST for the file
	file := &ast.File{
		Name: ast.NewIdent("main"), // Package name
		Decls: []ast.Decl{
			// Import declaration if needed
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
				},
			},
			// Type declaration for our struct
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Person"), // Struct name
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{ast.NewIdent("Name")},
										Type:  ast.NewIdent("string"),
										Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"name\"`"},
									},
									{
										Names: []*ast.Ident{ast.NewIdent("Age")},
										Type:  ast.NewIdent("int"),
										Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"age\"`"},
									},
									{
										Names: []*ast.Ident{ast.NewIdent("Address")},
										Type:  ast.NewIdent("string"),
										Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"address\"`"},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Format the AST
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, file); err != nil {
		panic(err)
	}

	// Write to file
	f, err := os.Create("person.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}
