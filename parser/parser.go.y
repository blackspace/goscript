%{

package parser

import (
    "goscript/ast"
)

var ParseResult []ast.Expr

%}


%union {
    Expr ast.Expr
    Expres []ast.Expr
}

%type <Expr> expr
%type <Expres> expres

%token <Expr> NUMBER VARIABLE
%token BLANKSPACE LFCR


%right '='
%left '+' '-'
%left '*' '/'


%%

expres :expr
    {
        $$=[]ast.Expr{$1}
        ParseResult=append(ParseResult,$$...)
    }
    |expres expr
    {
        $$=append($1,$2)
        ParseResult=$$
    };


expr : NUMBER
    |VARIABLE
    |VARIABLE '=' expr
    {
        $$=&ast.AssignExpr{$1,$3}
    }
    | expr '+' expr
    {
        $$=&ast.AddExpr{$1,$3}
    }
    |expr '-' expr
    {
        $$=&ast.SubExpr{$1,$3}
    }
    |expr '*' expr
    {
        $$=&ast.MultiExpr{$1,$3}
    }
    |expr '/' expr
    {
        $$=&ast.DivExpr{$1,$3}
    };
%%




