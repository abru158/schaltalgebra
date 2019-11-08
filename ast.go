package main

import "fmt"

type (
	expression interface {
		eval(*stateMachine) bool
	}

	astExpression struct {
		l, r expression
		op   operation
	}

	astInputExpression struct {
		v rune
	}

	astNegatedExpression struct {
		expr expression
	}
)

func (e *astExpression) eval(s *stateMachine) bool {
	return e.op(e.l.eval(s), e.r.eval(s))
}

func (e *astInputExpression) eval(s *stateMachine) bool {
	return s.get(e.v)
}

func (e *astNegatedExpression) eval(s *stateMachine) bool {
	return !e.expr.eval(s)
}

func printAST(expr expression, ident string) {
	switch e := expr.(type) {
	case *astNegatedExpression:
		fmt.Printf("%snegate:\n", ident)
		printAST(e.expr, ident+"  ")
	case *astInputExpression:
		fmt.Printf("%sinput: %s\n", ident, string(e.v))
	case *astExpression:
		fmt.Printf("%sexpression:\n", ident)
		ident += "  "
		fmt.Printf("%sleft:\n", ident)
		printAST(e.l, ident)
		fmt.Printf("%sright:\n", ident)
		printAST(e.r, ident)
	default:
		fmt.Printf("unknown ast node: %T", e)
	}
}
