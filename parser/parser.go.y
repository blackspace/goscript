%{

package parser

import (
    "goscript/ast"
)

var ResultExpr ast.Expr

%}


%union {
    Expr ast.Expr
}

%type <Expr> expr top

%token <Expr> NUMBER

%right '='
%left '+' '-'
%left '*' '/'


%%

top :  expr
    {

    	ResultExpr = $1

    };


expr : NUMBER
    | expr '+' expr
    {
        $$=&ast.AddExpr{$1,$3}
    };


%%




