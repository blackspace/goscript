package parser

import (
	"goscript/ast"
	"strings"
	"log"
	"bufio"
)

func ParseString(src string) (ast.Expr,error)  {
	r:=bufio.NewReader(strings.NewReader(src))

	l:= Lexer{Reader:r}
	yyParse(&l)

	if  ResultExpr!=nil {
		return ResultExpr,nil
	} else {
		log.Fatal("Express is nil")
		return nil,nil
	}


}
