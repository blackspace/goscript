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

%type <Expr> expr assign_expr simple_expr if_expr for_expr stmt_expr block_expr increment_decrement_expr func_expr
%type <Exprs> exprs  top

%token <Expr> NUMBER VARIABLE BOOL STRING BREAK
%token BLANKSPACE LFCR WORD IF FOR SWITCH FUNCTION CLASS DOUBLEADD DOUBLESUB

%nonassoc '>' GREATEREQUAL '<' LESSEQUAL

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


expr : BREAK
    {
        $$=&ast.BreakExpr{}
    }
    |simple_expr
    |stmt_expr
    |assign_expr
    |increment_decrement_expr
    |func_expr
    |block_expr;


block_expr : '{' exprs '}'
    {
        $$=&ast.BlockExpr{$2}
    }
    | '{' '}'
    {
        $$=&ast.BlockExpr{}
    }

stmt_expr: if_expr
    |for_expr;


assign_expr: VARIABLE '=' expr
    {
        $$=&ast.AssignExpr{$1,$3}
    };

increment_decrement_expr: VARIABLE DOUBLEADD
    {
        $$=&ast.SubfixDoubleAddExpr{$1}
    }
    |VARIABLE DOUBLESUB
    {
           $$=&ast.SubfixDoubleSubExpr{$1}
    };


for_expr : FOR assign_expr ';' simple_expr ';' increment_decrement_expr block_expr
    {
        $$=&ast.ForExpr{$2,$4,$6,$7}
    }
    |FOR  ';' simple_expr ';' increment_decrement_expr block_expr
    {
             $$=&ast.ForExpr{nil,$3,$5,$6}
    }
    |FOR  ';'  ';' increment_decrement_expr block_expr
    {
            $$=&ast.ForExpr{nil,nil,$4,$5}
    }
    |FOR  ';' ';'  block_expr
    {
          $$=&ast.ForExpr{nil,nil,nil,$4}
    }
    |FOR  assign_expr ';' ';' increment_decrement_expr  block_expr
    {
           $$=&ast.ForExpr{$2,nil,$5,$6}
    };



func_expr : FUNCTION '('  ')' block_expr
    {
        $$=&ast.FuncExpr{}
    };

if_expr : IF simple_expr block_expr
    {
        $$=&ast.IFExpr{$2,$3,nil}
    }
    |if_expr ELSE block_expr
    {
        ($1.(*ast.IFExpr)).Expr2=$3

        $$=$1
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
    }
    |simple_expr '>' simple_expr
    {
        $$=&ast.GreaterThanExpr{$1,$3}
    }
    |simple_expr GREATEREQUAL simple_expr
    {
        $$=&ast.GreaterEqualExpr{$1,$3}
    }
    |simple_expr '<' simple_expr
    {
        $$=&ast.LessThanExpr{$1,$3}
    }
    |simple_expr LESSEQUAL simple_expr
    {
        $$=&ast.LessEqualExpr{$1,$3}
    };
%%




