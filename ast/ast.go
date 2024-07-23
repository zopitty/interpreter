package ast

import (
	"bytes"

	"github.com/zopitty/interpreter/token"
)

// base node interface
type Node interface {
	TokenLiteral() string
	String() string // for debugging and comapring them with other AST nodes
}

// does not produce any values
// type of Node, all statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// produces a value
// type of Node, all expression  nodes implement this
type Expression interface {
	Node
	expressionNode()
}

// Root Node of AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LET statement node
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identitifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// IDENT node
type Identitifier struct {
	Token token.Token
	Value string
}

func (i *Identitifier) expressionNode()      {}
func (i *Identitifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identitifier) String() string       { return i.Value }

// Return statement node
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Expression statement (wrapper), only for an expression
// let x = 5; -> let statement
// x + 10; -> expression statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
