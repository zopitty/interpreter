package ast

import "github.com/zopitty/interpreter/token"

type Node interface {
	TokenLiteral() string
}

// does not produce any values
// type of Node
type Statement interface {
	Node
	statementNode()
}

// produces a value
// type of Node
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

// LET statement node
type LetStatement struct {
	Token token.Token // the token.LET toke
	Name  *Identitifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// IDENT node
type Identitifier struct {
	Token token.Token
	Value string
}

func (i *Identitifier) expressionNode()      {}
func (i *Identitifier) TokenLiteral() string { return i.Token.Literal }

// Return node
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
