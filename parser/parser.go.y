%{

package parser

import (
    "goscript/ast"
)

var ParseResult []ast.Expr

%}


%union {
    Expr ast.Expr
    Params ast.Params
    Exprs []ast.Expr
    String string
    Strings []string
}

%type <Expr> expr assign_expr simple_expr var_expr
%type <Expr> if_expr for_expr stmt_expr block_expr increment_decrement_expr


%type <Expr> func_define_expr  func_call_expr value_expr
%type <Expr> class_expr inclass_expr inmethod_expr
%type <Expr> method_call_expr attribute_set_expr method_define_expr attribute_define_expr
%type <String> func_name class_name method_name attribute_name
%type <String> param
%type <Strings> params


%type <Exprs> exprs  top
%type <Exprs> values_expr inclass_exprs inmethod_exprs


%token <Expr> NUMBER BOOL STRING BREAK RETURN
%token <String> WORD
%token BLANKSPACE LFCR  IF FOR SWITCH DOUBLEADD DOUBLESUB DOUBLEAT
%token FUNCTION CLASS DEF END DO

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


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
class_expr : CLASS class_name '{'  inclass_exprs '}'
    {
        $$=&ast.Class{}
    };

inclass_expr: simple_expr|stmt_expr|assign_expr|attribute_set_expr
            |increment_decrement_expr|func_define_expr |block_expr
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
        $$=&ast.Method{}
    };
    |DEF method_name '(' params ')' '{' inmethod_exprs '}'
    {
        $$=&ast.Method{}
    };

inmethod_expr:simple_expr|stmt_expr|assign_expr|attribute_set_expr
        |increment_decrement_expr |block_expr | attribute_define_expr;

inmethod_exprs:inmethod_expr
        {
            $$=append(make([]ast.Expr,0),$1)
        }
        |inmethod_exprs inmethod_expr
        {
            $$=append($1,$2)
        };

attribute_define_expr: '@' attribute_name '=' expr
      {
          $$=&ast.Attribute{}
      }
      | DOUBLEAT attribute_name '=' expr
      {
          $$=&ast.ClassAttribute{}
      };


method_call_expr:  var_expr '.' method_name '('   ')'
      | var_expr '.' method_name '(' values_expr  ')';
attribute_set_expr: var_expr '.' attribute_name '=' expr;

class_name: WORD;
method_name: WORD;
attribute_name: WORD;


////////////////////////////////////////////////////////////////////////////////////////

expr : simple_expr
    |stmt_expr
    |assign_expr
    |attribute_set_expr
    |increment_decrement_expr
    |func_define_expr
    |block_expr
    |class_expr;

func_define_expr : FUNCTION  func_name '(' ')' block_expr
    {
        $$=&ast.FuncDefineExpr{$2,nil,$5}
    }
    |FUNCTION  func_name '(' params ')' block_expr
    {
        $$=&ast.FuncDefineExpr{$2,$4,$6}
    };

func_call_expr : func_name '(' ')'
    {
        $$=&ast.FuncCallExpr{$1,nil}
    }
    |func_name '(' values_expr ')'
    {
        $$=&ast.FuncCallExpr{$1,$3}
    };

func_name: WORD;

params: param
    {
        $$=append(make([]string,0),$1)
    }
    | params ',' param
    {
        $$=append($1,$3)
    };

param:WORD;

values_expr: value_expr
    {

        $$=append(make([]ast.Expr,0),$1)
    }
    | values_expr ',' value_expr
    {
        $$=append($1,$3)
    };

value_expr: simple_expr;

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
    |for_expr;


assign_expr: var_expr '=' expr
    {
        $$=&ast.AssignExpr{$1,$3}
    };

increment_decrement_expr: var_expr DOUBLEADD
    {
        $$=&ast.SubfixDoubleAddExpr{$1}
    }
    |var_expr DOUBLESUB
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



if_expr : IF simple_expr block_expr
    {
        $$=&ast.IFExpr{$2,$3,nil}
    }
    |if_expr ELSE block_expr
    {
        ($1.(*ast.IFExpr)).Expr2=$3

        $$=$1
    };


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

simple_expr: BOOL|NUMBER|STRING | var_expr | func_call_expr | method_call_expr
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

block_expr : '{' exprs '}'
    {
        $$=&ast.BlockExpr{$2}
    }
    | '{' '}'
    {
        $$=&ast.BlockExpr{}
    };

var_expr : WORD
    {
        $$=&ast.Variable{$1}
    };
%%




