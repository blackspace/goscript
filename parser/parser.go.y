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

%type <Expr> expr arithmetic_expr logical_expr assign_expr string_expr
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


expr : arithmetic_expr
    |logical_expr
    |assign_expr
    |string_expr;

logical_expr: BOOL
    |logical_expr AND logical_expr
    {
        $$=&ast.ANDExpr{$1,$3}
    }
    |logical_expr OR logical_expr
    {
        $$=&ast.ORExpr{$1,$3}
    };

assign_expr: VARIABLE '=' arithmetic_expr
    {
        $$=&ast.AssignExpr{$1,$3}
    }

arithmetic_expr : NUMBER
    |VARIABLE
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

string_expr: STRING
    |string_expr '+' string_expr
    {
        $$=&ast.StringConcateExpr{$1,$3}
    }
%%




