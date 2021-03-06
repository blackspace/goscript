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

	if p := FindPattern(l.Buf); p == nil {
		goto t
	} else 	{
		v, pr, h := p.BuildFun(l.Buf, l.Reader)

		if h {
			l.PreRune = pr
			l.HasPreRune=true
		}

		switch t:=p.GetToken();t {
		case POUNDCOMMENT:
			l.Buf=""
			goto t
		case MULTILINECOMMENT:
			l.Buf=""
			goto t
		case INT:
			lval.Expr = v.(*ast.Int)
			return INT
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
		case  '/':
			r, _, err := l.Reader.ReadRune()
			if err == io.EOF {
				return '/'
			}

			if r=='/' {
				l.Buf = l.Buf + string(r)

				if p := FindPattern(l.Buf); p != nil {
					_, pr, h := p.BuildFun(l.Buf, l.Reader)

					if h {
						l.PreRune = pr
						l.HasPreRune=true
					}

					l.Buf=""
					goto t
				}
			} else if r=='*' {
				l.Buf = l.Buf + string(r)
				goto t
			} else {
				l.HasPreRune=true
				l.PreRune=r
				return '/'
			}
		case '=':
			r, _, err := l.Reader.ReadRune()
			if err == io.EOF {
				return '='
			}

			if r=='=' {
				return DOUBLEEQUAL
			}  else {
				l.HasPreRune=true
				l.PreRune=r
				return '='
			}
		case '*':
			return p.GetToken()
		case '{','}':
			return p.GetToken()
		case '(',')':
			return p.GetToken()
		case '[',']':
			return p.GetToken()
		case ';':
			return p.GetToken()
		case ',':
			return p.GetToken()
		case '.':
			return p.GetToken()
		case ':':
			return p.GetToken()
		case '@':
			r, _, err := l.Reader.ReadRune()

			if err == io.EOF {
				return '@'
			}

			if r=='@' {
				return DOUBLEADD
			} else {
				l.HasPreRune=true
				l.PreRune=r
				return '@'
			}
		case '>':
			r, _, err := l.Reader.ReadRune()

			if err == io.EOF {
				return '>'
			}

			if r=='=' {
				return GREATEREQUAL
			} else {
				l.HasPreRune=true
				l.PreRune=r
				return '>'
			}
		case '<':
			r, _, err := l.Reader.ReadRune()

			if err == io.EOF {
				return '<'
			}

			if r=='=' {
				return LESSEQUAL
			} else {
				l.HasPreRune=true
				l.PreRune=r
				return '<'
			}
		case AND:
			return AND
		case OR:
			return OR
		case WORD:
			w :=v.(string)
			t,v,ok:=ParseWord(w)

			if ok {
				switch t {
				case BOOL:
					lval.Expr = v.(*ast.Bool)
					return BOOL
				default:
					return t

				}
			} else {
				lval.String = w
				return WORD
			}

		case BLANKSPACE:
			l.Buf=""
			goto t
		case LFCR:
			l.Buf=""
			goto t
		}

	}

	return 0

}


func (l *Lexer)Error(msg string) {
	log.Println("Parsing:"+msg)
}
