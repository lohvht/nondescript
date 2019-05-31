package ast

import "github.com/lohvht/went/lang/token"

//============================================================================
// This file is generated by tool/ast-generate.go
// This file describes the individual node types as well as their implementations
// All edits should be made to the nodeTemplate in ast-generate.go instead
//============================================================================

// Stmt AST nodes
type (
	Stmt interface {
		Node
		stmt()
	}
	// ExprStmt node
	ExprStmt struct {
		Expression Expr
	}

	// NameDeclStmt node
	NameDeclStmt struct {
		Name token.Token
		Init Expr
	}
)

func (n *ExprStmt) stmt()     {}
func (n *NameDeclStmt) stmt() {}

// Accept marshals the Visitor to the correct Visitor.visitXX method
func (n *ExprStmt) Accept(v Visitor) interface{} { return v.VisitExprStmt(n) }

// Accept marshals the Visitor to the correct Visitor.visitXX method
func (n *NameDeclStmt) Accept(v Visitor) interface{} { return v.VisitNameDeclStmt(n) }
