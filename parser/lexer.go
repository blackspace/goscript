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

		switch i := v.(type) {
		case ast.Expr:
			lval.Expr = i
			return p.GetToken()
		case rune:
			return int(i)
		case string:
			lval.String=i
			return p.GetToken()

		}


	} else {

		log.Fatal("No Pattern to be matched")
	}

	return 0

}


func (l *Lexer)Error(msg string) {
	log.Println("Parsing:"+msg)
}
