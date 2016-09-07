%{

package parser

import (
    "goscript/ast"
)

var ResultExpr ast.Expr

%}


%union {
    Expr ast.Expr
    String string
}

%type <Expr> expr top

%token <Expr> NUMBER
%token <String> WORD

%right '='
%left '+' '-'
%left '*' '/'


%%

top :  expr
    {
    	ResultExpr = $1

    };


expr : NUMBER
    | WORD '=' expr
    {
        $$=$3
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




