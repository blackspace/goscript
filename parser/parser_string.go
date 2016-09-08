package parser

import (
	"goscript/ast"
	"strings"
	"log"
	"bufio"
)

func ParseString(src string) ([]ast.Expr,error)  {
	r:=bufio.NewReader(strings.NewReader(src))

	l:= Lexer{Reader:r}
	yyParse(&l)

	if  ParseResult!=nil {
		return ParseResult,nil
	} else {
		log.Fatal("Parse result is empty")
		return nil,nil
	}


}
