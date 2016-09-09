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

%type <Expr> expr assign_expr simple_expr
%type <Expres> exprs

%token <Expr> NUMBER VARIABLE BOOL STRING
%token BLANKSPACE LFCR WORD
%token IF FOR SWITCH
%left AND OR
%token FUNCTION CLASS


%right '='
%left '+' '-'
%left '*' '/'


%%

exprs :expr
    {
        $$=[]ast.Expr{$1}
        ParseResult=append(ParseResult,$$...)
    }
    |exprs expr
    {
        $$=append($1,$2)
        ParseResult=$$
    };


expr : simple_expr
    |assign_expr;

assign_expr: VARIABLE '=' expr
    {
        $$=&ast.AssignExpr{$1,$3}
    }

simple_expr: BOOL|NUMBER|STRING |VARIABLE
    |simple_expr AND simple_expr
    {
        $$=&ast.ANDExpr{$1,$3}
    }
    |simple_expr OR simple_expr
    {
        $$=&ast.ORExpr{$1,$3}
    }
    | simple_expr '+' simple_expr
    {
        $$=&ast.AddExpr{$1,$3}
    }
    |simple_expr '-' simple_expr
    {
        $$=&ast.SubExpr{$1,$3}
    }
    |simple_expr '*' simple_expr
    {
        $$=&ast.MultiExpr{$1,$3}
    }
    |simple_expr '/' simple_expr
    {
        $$=&ast.DivExpr{$1,$3}
    };
%%




