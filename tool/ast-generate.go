package main

import (
	"bufio"
	"bytes"
	"flag"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type templateData interface {
	templateData()
}

// nodeType models the string data needed for the Node interface
type nodeType struct {
<<<<<<< HEAD
	DirName  string // parent directory name
	BaseName string // name of the node type
	Decls    []nodeImpl
=======
	DirName         string // parent directory name
	BaseName        string // name of the node type
	VisitReturnType string // type to be returned by the visitor
	Decls           []nodeImpl
>>>>>>> lox-impl-temp
}

func (nt nodeType) templateData() {}

// nodeImpl models the string data needed for Node implementations
type nodeImpl struct {
	Name   string  // name of the node implementation
	Fields []field // a mapping of fieldnames to types
}

// field is represents a name-type pair
type field struct{ name, typ string }

func (f field) String() string { return f.name + " " + f.typ }

type astIntInfo struct {
	DirName string
	Types   []nodeType
}

func (nt astIntInfo) templateData() {}

func main() {
	var outdir string
	flag.StringVar(&outdir, "outdir", "", "The output directory where the generated files will be saved to. (Required)")
	flag.Parse()

	if outdir == "" {
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		os.Mkdir(outdir, 0755)
	}
	expr := nodeType{
<<<<<<< HEAD
		DirName:  outdir,
		BaseName: "Expr",
		Decls: []nodeImpl{
=======
		DirName:         outdir,
		BaseName:        "Expr",
		VisitReturnType: "interface{}",
		Decls: []nodeImpl{
			nodeImpl{Name: "NameExpr", Fields: []field{
				field{"Name", "token.Token"},
			}},
>>>>>>> lox-impl-temp
			nodeImpl{Name: "GrpExpr", Fields: []field{
				field{"LeftRound", "token.Token"}, field{"Expression", "Expr"},
				field{"RightRound", "token.Token"},
			}},
			nodeImpl{Name: "BinExpr", Fields: []field{
				field{"Left", "Expr"}, field{"Op", "token.Token"}, field{"Right", "Expr"},
			}},
			nodeImpl{Name: "UnExpr", Fields: []field{
				field{"Op", "token.Token"}, field{"Operand", "Expr"},
			}},
			nodeImpl{Name: "BasicLit", Fields: []field{
<<<<<<< HEAD
				field{"Value", "interface{}"}, field{"Token", "token.Token"},
			}},
		},
	}
	types := []nodeType{expr}
=======
				field{"Text", "string"}, field{"Typ", "token.Type"}, field{"Token", "token.Token"},
				field{"Value", "interface{}"},
			}},
		},
	}
	stmt := nodeType{
		DirName:         outdir,
		BaseName:        "Stmt",
		VisitReturnType: "",
		Decls: []nodeImpl{
			nodeImpl{Name: "ExprStmt", Fields: []field{field{"Expression", "Expr"}}},
			nodeImpl{Name: "NameDeclStmt", Fields: []field{
				field{"Name", "token.Token"}, field{"Init", "Expr"},
			}},
		},
	}
	types := []nodeType{stmt, expr}
>>>>>>> lox-impl-temp
	for _, typ := range types {
		generateAndFormatFile(typ, typ.DirName, typ.BaseName, nodeTemplate)
	}
	ast := astIntInfo{DirName: outdir, Types: types}
	generateAndFormatFile(ast, ast.DirName, "interface", astInterfaceTemplate)
}

// generateAndFormatFile generates a file given the templateData it needs,
// directory name to store, basename (filename) and its templateText as a string
func generateAndFormatFile(td templateData, dirname, baseName, templateText string) {
	t := generateTemplate(baseName, templateText)
	var src bytes.Buffer
	t.Execute(&src, td)
	_, err := format.Source(src.Bytes())
	if err != nil {
		panic(err) // TODO: HANDLE ERROR properly
	}
	f, err := os.Create(filepath.Join(dirname, strings.ToLower(baseName)+".go"))
	if err != nil {
		panic(err) // TODO: HANDLE ERROR properly
	}
	fw := bufio.NewWriter(f)
	defer f.Close()
	goimports := exec.Command("goimports")
	goimports.Stdin = &src
	goimports.Stdout = fw
	err = goimports.Run()
	if err != nil {
		panic(err) // TODO: HANDLE ERROR properly
	}
	fw.Flush()
}

// generateTemplate generates a *Template object with some common string
// manipulating functions baked into its FuncMap
func generateTemplate(name string, templateText string) *template.Template {
	funcMap := template.FuncMap{
		"ToUpper":      strings.ToUpper,
		"ToLower":      strings.ToLower,
		"JoinString":   strings.Join,
		"FilePathBase": filepath.Base,
		"FieldJoin": func(m []field, kvSep string, sep string) string {
			var buf bytes.Buffer
			first := true
			for _, field := range m {
				k, v := field.name, field.typ
				if !first {
					buf.WriteString(sep)
				}
				first = false
				buf.WriteString(k)
				buf.WriteString(kvSep)
				buf.WriteString(v)
			}
			return buf.String()
		},
<<<<<<< HEAD
=======
		"RETURN": func(visitReturnType string) string {
			if visitReturnType == "" {
				return visitReturnType
			}
			return "return"
		},
>>>>>>> lox-impl-temp
	}
	return template.Must(template.New(name).Funcs(funcMap).Parse(templateText))
}

<<<<<<< HEAD
var javatemp = `
package com.craftinginterpreters.lox;

import java.util.List;

abstract class {{.BaseName}} {

	interface Visitor<R> {
		{{- range $i, $nodeImpl := .Decls}}
		R visit{{$nodeImpl.Name}}{{$.BaseName}}({{$nodeImpl.Name}} {{$.BaseName | ToLower}});
		{{- end}}
	}
	{{range $i, $nodeImpl := .Decls}}
	static class {{$nodeImpl.Name}} extends {{$.BaseName}} {
		{{- range $fn, $ft := $nodeImpl.Fields}}
		final {{$fn}};
		{{- end}}

		{{$nodeImpl.Name}}({{FieldJoin $nodeImpl.Fields " " ", "}}) {
			{{- range $fn, $ft := $nodeImpl.Fields}}
			this.{{$fn}} = {{$fn}};
			{{- end}}
		}

		// Visitor pattern
		<R> R accept(Visitor<R> visitor) {
			return visitor.visit{{$nodeImpl.Name}}{{$.BaseName}}(this);
		}
	}
	{{end}}
	abstract <R> R accept(Visitor<R> visitor);
}
`

=======
>>>>>>> lox-impl-temp
var nodeTemplate = `
package {{.DirName | FilePathBase}}

import "github.com/lohvht/went/lang/token"

//============================================================================
// This file is generated by tool/ast-generate.go
// This file describes the individual node types as well as their implementations
// All edits should be made to the nodeTemplate in ast-generate.go instead
//============================================================================

// {{.BaseName}} AST nodes
type (
	{{.BaseName}} interface {
		Node
		{{.BaseName | ToLower}}()
	}

	{{- range $i, $nodeImpl := .Decls}}
	// {{$nodeImpl.Name}} node
	{{$nodeImpl.Name}} struct {
		{{FieldJoin $nodeImpl.Fields " " "\n"}}
	}
	{{end}}
)

{{- range $i, $nodeImpl := .Decls}}
func (n *{{$nodeImpl.Name}}) {{$.BaseName | ToLower}}() {}
{{- end}}

{{range $i, $nodeImpl := .Decls}}
<<<<<<< HEAD
func (n *{{$nodeImpl.Name}}) accept(v Visitor) interface{} { return v.visit{{$nodeImpl.Name}}(n) }
=======
// Accept marshals the Visitor to the correct Visitor.visitXX method
func (n *{{$nodeImpl.Name}}) Accept(v Visitor) interface{} { return v.Visit{{$nodeImpl.Name}}(n) }
>>>>>>> lox-impl-temp
{{- end}}
`

var astInterfaceTemplate = `
package {{.DirName | FilePathBase}}

import "github.com/lohvht/went/lang/token"

//============================================================================
// This file is generated by tool/ast-generate.go
// This file contains the interfaces used for AST nodes and tree walkers
// All edits should be made to the astInterfaceTemplate in ast-generate.go instead
//============================================================================

// Node is the common interface for all AST nodes
type Node interface {
<<<<<<< HEAD
	accept(Visitor) interface{}
}

// Visitor is the interface used to implement visitor pattern for the AST
=======
	// Accept marshals the Visitor to the correct Visitor.visitMethod
	Accept(Visitor) interface{}
}

// Visitor is the interface used to implement visitor pattern for the AST
// Interpreter nodes with their visit methods are supposed to return nil as
// statements should resolve to a value
>>>>>>> lox-impl-temp
type Visitor interface {
	{{- range $i, $type := $.Types}}
	// visit {{$type.BaseName}} node functions
	{{- range $j, $nodeImpl := $type.Decls}}
<<<<<<< HEAD
	visit{{$nodeImpl.Name}}(*{{$nodeImpl.Name}}) interface{}
	{{- end}}
	{{- end}}
=======
	Visit{{$nodeImpl.Name}}(*{{$nodeImpl.Name}}) interface{}
	{{- end}}
	{{end}}
>>>>>>> lox-impl-temp
}
`
