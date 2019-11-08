package main

import (
	"math"
)

type errSyntax struct {
	msg string
}

func (e *errSyntax) Error() string {
	return e.msg
}

type parser struct {
	s *scanner

	lowestVar, highestVar rune
}

func (p *parser) init(s *scanner) {
	p.s = s
	p.lowestVar = math.MaxInt32
	p.highestVar = math.MinInt32
}

func (p *parser) next() rune {
	c := p.s.c
	p.s.next()
	return c
}

func (p *parser) parsePrimaryExpression() expression {
	c := p.s.c
	p.next()
	switch c {
	case '(':
		e := p.parseOperationExpression()
		if p.s.c != ')' {
			panic(&errSyntax{
				msg: "expected ')'",
			})
		}
		p.next()
		return e
	default:
		if 'a' <= c && c <= 'h' {
			if p.lowestVar > c {
				p.lowestVar = c
			}
			if p.highestVar < c {
				p.highestVar = c
			}
			return &astInputExpression{
				v: c,
			}
		}
		panic(&errSyntax{
			msg: "unexpected '" + string(c) + "'",
		})
	}
}

func (p *parser) parseUnaryExpression() expression {
	switch p.s.c {
	case '!', '-':
		p.next()
		return &astNegatedExpression{
			expr: p.parseUnaryExpression(),
		}
	}
	return p.parsePrimaryExpression()
}

func (p *parser) parseAndExpression() expression {
	left := p.parseUnaryExpression()
	switch p.s.c {
	case '&', '^':
		p.next()
		return &astExpression{
			l:  left,
			op: opAND,
			r:  p.parseAndExpression(),
		}
	}
	return left
}

func (p *parser) parseOrExpression() expression {
	left := p.parseAndExpression()
	switch p.s.c {
	case '|', 'v':
		p.next()
		return &astExpression{
			l:  left,
			op: opOR,
			r:  p.parseOrExpression(),
		}
	}
	return left
}

func (p *parser) parseOperationExpression() expression {
	return p.parseOrExpression()
}

func (p *parser) parse() (expr expression, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case *errSyntax:
				err = e
			default:
				panic(r)
			}
		}
	}()
	e := p.parseOperationExpression()
	if p.s.pos < len(p.s.src) {
		return nil, &errSyntax{
			msg: "could not parse entire formula",
		}
	}
	return e, nil
}
