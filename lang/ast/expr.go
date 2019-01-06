package ast

import "github.com/lohvht/went/lang/token"

//============================================================================
// This file is generated by tool/ast-generate.go
// This file describes the individual node types as well as their implementations
// All edits should be made to the nodeTemplate in ast-generate.go instead
//============================================================================

// Expr AST nodes
type (
	Expr interface {
		Node
		expr()
	}
	// GrpExpr node
	GrpExpr struct {
		Expression Expr
	}

	// BinExpr node
	BinExpr struct {
		Left  Expr
		Right Expr
		Op    token.Token
	}

	// UnExpr node
	UnExpr struct {
		Operand Expr
		Op      token.Token
	}

	// BasicLit node
	BasicLit struct {
		Text string
	}
)

func (n *GrpExpr) expr()  {}
func (n *BinExpr) expr()  {}
func (n *UnExpr) expr()   {}
func (n *BasicLit) expr() {}

func (n *GrpExpr) accept(v Visitor) interface{}  { return v.visitGrpExpr(n) }
func (n *BinExpr) accept(v Visitor) interface{}  { return v.visitBinExpr(n) }
func (n *UnExpr) accept(v Visitor) interface{}   { return v.visitUnExpr(n) }
func (n *BasicLit) accept(v Visitor) interface{} { return v.visitBasicLit(n) }
