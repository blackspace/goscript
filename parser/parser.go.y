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
    String string
    Strings []string
}

%type <Expr> expr literal_expr simple_expr
%type <Expr> if_expr for_expr stmt_expr  increment_decrement_expr


%type <Expr>  get_expr set_expr

%type <Expr> func_define_expr lambda_define_expr
%type <Expr> block_expr empty_block_expr not_empty_block_expr
%type <Expr> class_expr inclass_expr inmethod_expr method_define_expr
%type <Expr> call_expr
%type <Expr> array_expr empty_array not_empty_array array_set_expr
%type <Expr> object_expr empty_object_expr not_empty_object_expr
%type <Expr> key_value_expr

%type <String> func_name method_name
%type <String> param
%type <Strings> params member_path


%type <Exprs> exprs  top
%type <Exprs> values_expr array_element_list object_element_list
%type <Exprs> inclass_exprs inmethod_exprs


%token <Expr> INT BOOL STRING BREAK RETURN
%token <String> WORD
%token BLANKSPACE LFCR   FOR SWITCH  DOUBLEAT
%token FUNCTION CLASS DEF END DO POUNDCOMMENT DOUBLESLASHCOMMENT MULTILINECOMMENT LAMBDA
%token DOUBLEADD DOUBLESUB
%token IF ELSE
%token '='


%left AND OR
%left '+' '-'
%left '*' '/'

%nonassoc '>' GREATEREQUAL '<' LESSEQUAL
%nonassoc LOWER_SQUARE
%nonassoc '[' ']'

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

////////////////////////////////////////////////////////////////////////////////////////

expr : simple_expr
    |stmt_expr
    |set_expr
    |array_set_expr
    |increment_decrement_expr
    |func_define_expr
    |lambda_define_expr
    |class_expr


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
class_expr : CLASS WORD  '{'  inclass_exprs '}'
    {
        $$=&ast.Class{$2,$4}
    }
    |CLASS WORD  '{' '}'
    {
        $$=&ast.Class{$2,nil}
    };

inclass_expr: stmt_expr|set_expr
            |increment_decrement_expr
            |class_expr | method_define_expr

inclass_exprs: inclass_expr
    {
       $$=append(make([]ast.Expr,0),$1)
    }
    |inclass_exprs inclass_expr
    {
       $$=append($1,$2)
    };

method_define_expr: DEF method_name '(' ')' '{' inmethod_exprs '}'
    {
        $$=&ast.MethodDefineExpr{$2,nil,$6}
    }
    |DEF method_name '(' ')' '{'  '}'
    {
        $$=&ast.MethodDefineExpr{$2,nil,nil}
    }
    |DEF method_name '(' params ')' '{' inmethod_exprs '}'
    {
        $$=&ast.MethodDefineExpr{$2,$4,$7}
    }
    |DEF method_name '(' params ')' '{'  '}'
    {
        $$=&ast.MethodDefineExpr{$2,$4,nil}
    }

method_name: WORD

inmethod_expr:simple_expr|stmt_expr|set_expr
        |increment_decrement_expr

inmethod_exprs:inmethod_expr
    {
        $$=append(make([]ast.Expr,0),$1)
    }
    |inmethod_exprs inmethod_expr
    {
        $$=append($1,$2)
    };


call_expr:  member_path '('   ')'
    {
        $$=&ast.CalledExpr{$1,nil}
    }
    | member_path '(' values_expr  ')'
    {
        $$=&ast.CalledExpr{$1,$3}
    };

get_expr: member_path %prec LOWER_SQUARE
    {
        $$=&ast.GetExpr{$1}
    }
    |member_path '[' INT ']'
    {
        $$=&ast.ArrayGetExpr{$1,$3}
    }

set_expr: member_path '=' expr
    {
        $$=&ast.SetExpr{$1,$3}
    };


member_path: WORD
    {
        $$=append(make([]string,0,16),$1)
    }
    |member_path '.' WORD
    {
        $$=append($1,$3)
    }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func_define_expr : FUNCTION  func_name '(' ')' '{' exprs '}'
    {
        $$=&ast.FuncDefineExpr{$2,nil,$6}
    }
    |FUNCTION  func_name '(' params ')' '{' exprs '}'
    {
        $$=&ast.FuncDefineExpr{$2,$4,$7}
    }

lambda_define_expr:LAMBDA  '(' params ')' '{' exprs '}'
    {
        $$=&ast.LambdaDefineExpr{$3,$6}
    }
    |LAMBDA  '('')' '{' exprs '}'
    {
       $$=&ast.LambdaDefineExpr{nil,$5}
    };

func_name: WORD;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
params: param
    {
        $$=append(make([]string,0),$1)
    }
    | params ',' param
    {
        $$=append($1,$3)
    };

param:WORD;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
stmt_expr: BREAK
    {
        $$=&ast.BreakExpr{}
    }
    |RETURN simple_expr
    {
        $$=&ast.ReturnExpr{$2}
    }
    |if_expr
    |for_expr
    |not_empty_block_expr;

for_expr : FOR set_expr ';' simple_expr ';' increment_decrement_expr block_expr
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
    |FOR  set_expr ';' ';' increment_decrement_expr block_expr
    {
           $$=&ast.ForExpr{$2,nil,$5,$6}
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

block_expr: empty_block_expr|not_empty_block_expr


empty_block_expr: '{' '}'
    {
        $$=&ast.BlockExpr{nil}
    }
not_empty_block_expr: '{' exprs '}'
    {
        $$=&ast.BlockExpr{$2}
    }


////////////////////////////////////////////////////////////////////////////////////////
increment_decrement_expr: member_path DOUBLEADD
{
    $$=&ast.SubfixDoubleAddExpr{$1}
}
|member_path DOUBLESUB
{
    $$=&ast.SubfixDoubleSubExpr{$1}
};


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

simple_expr: literal_expr| call_expr | get_expr
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

literal_expr:BOOL|INT|STRING|array_expr|object_expr


array_expr: empty_array | not_empty_array

empty_array: '[' ']'
    {
        $$=&ast.ArrayExpr{nil}
    }
not_empty_array: '[' array_element_list ']'
    {
        $$=&ast.ArrayExpr{$2}
    }
array_element_list: simple_expr
    {
        $$=append(make([]ast.Expr,0,1024),$1)
    }
    |array_element_list ',' simple_expr
    {
        $$=append($1,$3)
    }

array_set_expr: member_path '[' INT ']' '=' simple_expr
    {
        $$=&ast.ArraySetExpr{$1,$3,$6}
    }

object_expr: empty_object_expr | not_empty_object_expr

empty_object_expr: '{' '}'
{
    $$=&ast.ObjectExpr{nil}
}

not_empty_object_expr:'{' object_element_list '}'
{
    $$=&ast.ObjectExpr{$2}
}

object_element_list: key_value_expr
    {
        $$=append(make([]ast.Expr,0,1024),$1)
    }
    |object_element_list ',' key_value_expr
    {
        $$=append($1,$3)
    }


key_value_expr : WORD ':' simple_expr
    {
        $$=&ast.KeyValueExpr{$1,$3}
    }
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
values_expr: simple_expr
    {

        $$=append(make([]ast.Expr,0),$1)
    }
    | values_expr ',' simple_expr
    {
        $$=append($1,$3)
    };
%%