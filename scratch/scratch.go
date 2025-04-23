package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func main() {
	// Sample code
	const sourceCode = `
package main

func example() {
    x := 10
    y := 20
}
`
	// Create FileSet
	fset := token.NewFileSet()

	// Parse source
	file, err := parser.ParseFile(fset, "", sourceCode, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Modify the AST: Add a print statement after every assignment
	ast.Inspect(file, func(n ast.Node) bool {
		if assign, ok := n.(*ast.AssignStmt); ok {
			// Create a print statement
			printStmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("fmt"),
						Sel: ast.NewIdent("Println"),
					},
					Args: []ast.Expr{
						assign.Lhs[0], // Print the assigned variable
					},
				},
			}

			// Add import if needed
			addImport(file, "fmt")

			// TODO: Insert printStmt after the assignment
			// Note: Actually inserting the statement requires more complex AST manipulation
		}
		return true
	})

	// Print the modified AST
	printer.Fprint(os.Stdout, fset, file)
}

func addImport(file *ast.File, importPath string) {
	// Check if import already exists
	for _, imp := range file.Imports {
		if imp.Path.Value == `"`+importPath+`"` {
			return
		}
	}

	// Add new import
	file.Imports = append(file.Imports, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"` + importPath + `"`,
		},
	})
}

func TestASTParser(t *testing.T) {
	const src = `
package test
func example() {
    x := 42
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	// Count declarations
	var funcCount int
	ast.Inspect(file, func(n ast.Node) bool {
		if _, ok := n.(*ast.FuncDecl); ok {
			funcCount++
		}
		return true
	})

	if funcCount != 1 {
		t.Errorf("expected 1 function, got %d", funcCount)
	}
}
