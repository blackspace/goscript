package parser

import (
	"log"
	"io"
	"bufio"
	"goscript/parser/pattern"
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
	var r rune

	if l.HasPreRune {
		l.Buf = l.Buf + string(l.PreRune)
		l.HasPreRune=false
		l.PreRune=0
	} else {
		c, _, err := l.Reader.ReadRune()

		r=c

		if err == io.EOF {
			return 0
		}
	}

	l.Buf = l.Buf + string(r)

	log.Println(l.Buf)

	if p := pattern.FindPattern(l.Buf); p != nil {
		v, p, h := p.BuildFun(l.Buf, l.Reader)

		if h {
			l.PreRune = p
			l.HasPreRune=true
		}

		switch i := v.(type) {
		case ast.Expr:
			lval.Expr = i
			return NUMBER
		case rune:
			return int(i)

		}


	} else {

		log.Fatal("No Pattern to be matched")
	}

	return 0

}


func (l *Lexer)Error(msg string) {
	log.Println("Parsing:"+msg)
}
