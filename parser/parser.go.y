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

%type <Expr> expr arithmetic_expr bool_expr
%type <Expres> exprs

%token <Expr> NUMBER VARIABLE BOOL
%token BLANKSPACE LFCR WORD
%token IF FOR SWITCH
%left AND OR
%token FUNCTION CLASS


%right '='
%left '+' '-'
%left '*' '/'


%%

exprs :arithmetic_expr
    {
        $$=[]ast.Expr{$1}
        ParseResult=append(ParseResult,$$...)
    }
   |bool_expr
    {
        $$=[]ast.Expr{$1}
        ParseResult=append(ParseResult,$$...)
    }
    |exprs expr
    {
        $$=append($1,$2)
        ParseResult=$$
    };


expr : arithmetic_expr
    | bool_expr;

bool_expr: BOOL
    |bool_expr AND bool_expr
    {
        $$=&ast.ANDExpr{$1,$3}
    }
    |bool_expr OR bool_expr
    {
        $$=&ast.ORExpr{$1,$3}
    };

arithmetic_expr : NUMBER
    |VARIABLE
    |VARIABLE '=' arithmetic_expr
    {
        $$=&ast.AssignExpr{$1,$3}
    }
    | arithmetic_expr '+' arithmetic_expr
    {
        $$=&ast.AddExpr{$1,$3}
    }
    |arithmetic_expr '-' arithmetic_expr
    {
        $$=&ast.SubExpr{$1,$3}
    }
    |arithmetic_expr '*' arithmetic_expr
    {
        $$=&ast.MultiExpr{$1,$3}
    }
    |arithmetic_expr '/' arithmetic_expr
    {
        $$=&ast.DivExpr{$1,$3}
    };
%%




