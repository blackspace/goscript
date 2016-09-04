package parser

import (
	"log"
	"io"
	"goscript/ast"
)

type Lexer struct {
	Reader io.Reader
}

func (l *Lexer)Lex(lval * yySymType) int {
	buf :=make([]byte,1)
	if c,err:=l.Reader.Read(buf);err==nil {
		log.Println(c)
		lval.Expr=&ast.Number{1}
		return NUMBER
	} else if err==io.EOF {
		return 0
	}
	return 0

}


func (l *Lexer)Error(msg string) {
	log.Println("Parsing:"+msg)
}
