package parser

import (
	"goscript/ast"
	"strings"
	"log"
	"bufio"
	"errors"
)

func ParseString(src string) ([]ast.Expr,error)  {
	r:=bufio.NewReader(strings.NewReader(src))

	l:= Lexer{Reader:r}


	if yyParse(&l)!=0{
		return nil,errors.New("Parsing the codes wrongly")
	} else {
		if  ParseResult!=nil {
			return ParseResult,nil
		} else {
			log.Fatal("Parse result is empty")
			return nil,nil
		}
	}



}
