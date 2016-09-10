package parser

import (
	"log"
	"io"
	"bufio"
	"goscript/ast"
)

type Lexer struct {
	Reader  *bufio.Reader
	Buf string
	PreRune rune
	HasPreRune bool

}

func (l *Lexer)Lex(lval * yySymType) int {
	defer func() {l.Buf=""}()
t:
	if l.HasPreRune {
		l.Buf = l.Buf + string(l.PreRune)
		l.HasPreRune=false
		l.PreRune=0
	} else {
		r, _, err := l.Reader.ReadRune()

		if err == io.EOF {
			return 0
		}

		l.Buf = l.Buf + string(r)
	}

	if p := FindPattern(l.Buf); p != nil {
		v, pr, h := p.BuildFun(l.Buf, l.Reader)

		if h {
			l.PreRune = pr
			l.HasPreRune=true
		}

		switch t:=p.GetToken();t {
		case NUMBER:
			lval.Expr = v.(*ast.Number)
			return NUMBER
		case BOOL:
			lval.Expr = v.(*ast.Bool)
			return BOOL
		case STRING:
			lval.Expr = v.(*ast.String)
			return STRING
		case '+':
			r, _, err := l.Reader.ReadRune()

			if err == io.EOF {
				return '+'
			}

			if r=='+' {
				return DOUBLEADD
			} else {
				l.HasPreRune=true
				l.PreRune=r
				l.Buf=""
				return '+'
			}
		case '-':
			r, _, err := l.Reader.ReadRune()

			if err == io.EOF {
				return '-'
			}

			if r=='-' {
				return DOUBLESUB
			} else {
				l.HasPreRune=true
				l.PreRune=r
				return '-'
			}
		case '*','/','=':
			return int(v.(rune))
		case '{','}':
			return int(v.(rune))
		case AND:
			return AND
		case OR:
			return OR
		case WORD:
			w :=v.(string)
			switch t,v:=ParseWord(w);t {
			case VARIABLE:
				lval.Expr=v.(*ast.VarExpr)
				return t
			case BOOL:
				lval.Expr=v.(*ast.Bool)
				return t
			case IF:
				return IF
			default:
				return t

			}
		case BLANKSPACE:
			l.Buf=""
			goto t
		case LFCR:
			l.Buf=""
			goto t
		}

	} else {

		goto t
	}

	return 0

}


func (l *Lexer)Error(msg string) {
	log.Println("Parsing:"+msg)
}
