%{

package parser

import (
    "goscript/ast"
)

var ParseResult []ast.Expr

%}


%union {
    Expr ast.Expr
    Exprs []ast.Expr
}

%type <Expr> expr assign_expr simple_expr if_expr for_expr stmt_expr
%type <Exprs> exprs top

%token <Expr> NUMBER VARIABLE BOOL STRING
%token BLANKSPACE LFCR WORD IF FOR SWITCH FUNCTION CLASS DOUBLEADD DOUBLESUB

%right ELSE

%left AND OR

%right '='
%left '+' '-'
%left '*' '/'


%%

top: exprs
    {
        ParseResult=$1
    }

exprs :expr
    {
        $$=[]ast.Expr{$1}
    }
    |exprs expr
    {
        $$=append($1,$2)

    };


expr : simple_expr
    |stmt_expr
    |assign_expr;


stmt_expr: if_expr
    |for_expr;

for_expr : FOR expr ';' expr ';' expr '{' exprs '}'
    {
        $$=&ast.ForExpr{}
    }


if_expr : IF simple_expr '{' exprs '}'
    {
        $$=&ast.IFExpr{$2,$4,nil}
    }
    |if_expr ELSE '{' exprs '}'
    {
        ($1.(*ast.IFExpr)).Expr2=$4

        $$=$1
    };

assign_expr: VARIABLE '=' expr
    {
        $$=&ast.AssignExpr{$1,$3}
    }
    |VARIABLE DOUBLEADD
    {
        $$=&ast.SubfixDoubleAddExpr{$1}
    }
    |VARIABLE DOUBLESUB
    {
           $$=&ast.SubfixDoubleSubExpr{$1}
    };


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




