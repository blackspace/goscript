//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:3
import (
	"goscript/ast"
)

var ParseResult []ast.Expr

//line parser/parser.go.y:14
type yySymType struct {
	yys     int
	Expr    ast.Expr
	Params  ast.Params
	Exprs   []ast.Expr
	String  string
	Strings []string
}

const NUMBER = 57346
const BOOL = 57347
const STRING = 57348
const BREAK = 57349
const WORD = 57350
const BLANKSPACE = 57351
const LFCR = 57352
const IF = 57353
const FOR = 57354
const SWITCH = 57355
const FUNCTION = 57356
const CLASS = 57357
const DOUBLEADD = 57358
const DOUBLESUB = 57359
const GREATEREQUAL = 57360
const LESSEQUAL = 57361
const ELSE = 57362
const AND = 57363
const OR = 57364

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"BOOL",
	"STRING",
	"BREAK",
	"WORD",
	"BLANKSPACE",
	"LFCR",
	"IF",
	"FOR",
	"SWITCH",
	"FUNCTION",
	"CLASS",
	"DOUBLEADD",
	"DOUBLESUB",
	"'>'",
	"GREATEREQUAL",
	"'<'",
	"LESSEQUAL",
	"ELSE",
	"AND",
	"OR",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"','",
	"'{'",
	"'}'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:231

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 20,
	30, 14,
	-2, 51,
	-1, 56,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	-2, 47,
	-1, 57,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	-2, 48,
	-1, 58,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	-2, 49,
	-1, 59,
	18, 0,
	19, 0,
	20, 0,
	21, 0,
	-2, 50,
}

const yyNprod = 52
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 201

var yyAct = [...]int{

	9, 66, 69, 49, 45, 13, 3, 13, 7, 24,
	19, 74, 31, 32, 33, 34, 49, 25, 26, 62,
	27, 28, 29, 30, 13, 35, 4, 75, 48, 88,
	47, 11, 10, 12, 43, 20, 85, 86, 38, 61,
	13, 19, 60, 76, 77, 68, 13, 6, 24, 44,
	72, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 79, 11, 10, 12, 75, 20, 29, 30,
	67, 46, 82, 84, 70, 40, 83, 36, 37, 87,
	81, 65, 91, 1, 83, 83, 92, 2, 89, 90,
	95, 96, 73, 83, 71, 97, 78, 94, 93, 14,
	31, 32, 33, 34, 67, 25, 26, 41, 27, 28,
	29, 30, 8, 31, 32, 33, 34, 80, 25, 26,
	5, 27, 28, 29, 30, 36, 37, 21, 19, 11,
	10, 12, 15, 20, 35, 17, 22, 23, 16, 18,
	11, 10, 12, 15, 20, 0, 39, 22, 23, 0,
	18, 0, 0, 11, 10, 12, 15, 20, 19, 63,
	22, 23, 0, 18, 11, 10, 12, 0, 20, 19,
	42, 27, 28, 29, 30, 31, 32, 33, 34, 0,
	25, 26, 19, 27, 28, 29, 30, 0, 0, 25,
	26, 64, 27, 28, 29, 30, 11, 10, 12, 0,
	20,
}
var yyPact = [...]int{

	149, -1000, 149, -1000, 157, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 109, -1000, -1000, 16, -1000, 67, 136,
	-1000, 4, 192, -5, -1000, 192, 192, 192, 192, 192,
	192, 192, 192, 192, 192, 149, -1000, -1000, -23, -11,
	-1000, 125, -1000, 160, 95, -1000, -33, 59, 0, -1000,
	145, 145, 40, 40, -1000, -1000, 166, 166, 166, 166,
	-1000, -1000, 19, -1000, -1000, 12, -1000, 157, -1000, 27,
	82, 8, -23, 5, -1000, -1000, -1000, 192, -6, 43,
	43, -23, -1000, 61, -1000, -23, 58, -1000, 43, -23,
	-23, -1000, -1000, -1000, -23, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 6, 47, 26, 4, 138, 135, 120, 0, 8,
	112, 99, 1, 127, 11, 92, 87, 83, 81,
}
var yyR1 = [...]int{

	0, 17, 16, 16, 1, 1, 1, 1, 1, 1,
	10, 10, 11, 11, 13, 15, 15, 14, 18, 18,
	12, 8, 8, 7, 7, 7, 2, 9, 9, 6,
	6, 6, 6, 6, 5, 5, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	5, 6, 3, 4, 1, 1, 3, 1, 1, 3,
	1, 3, 2, 1, 1, 1, 3, 2, 2, 7,
	6, 5, 4, 6, 3, 3, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1,
}
var yyChk = [...]int{

	-1000, -17, -16, -1, -3, -7, -2, -9, -10, -8,
	5, 4, 6, -4, -11, 7, -5, -6, 14, 33,
	8, -13, 11, 12, -1, 23, 24, 26, 27, 28,
	29, 18, 19, 20, 21, 25, 16, 17, 22, -13,
	8, -16, 34, 30, -3, -4, -2, 35, -4, 8,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-1, -8, 30, 34, 31, -18, -12, -3, -8, 35,
	-3, 35, 31, -15, -14, 8, 31, 32, -3, 35,
	35, -9, -8, -4, -8, 31, 32, -12, 35, -9,
	-9, -8, -8, -14, -9, -8, -8, -8,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	36, 37, 38, 39, 40, 23, 24, 25, 0, 0,
	-2, 0, 0, 0, 3, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 27, 28, 0, 0,
	14, 0, 22, 0, 0, 39, 0, 0, 0, 51,
	41, 42, 43, 44, 45, 46, -2, -2, -2, -2,
	26, 35, 0, 21, 12, 0, 18, 20, 34, 0,
	0, 0, 0, 0, 15, 17, 13, 0, 0, 0,
	0, 0, 32, 0, 10, 0, 0, 19, 0, 0,
	0, 31, 11, 16, 0, 33, 30, 29,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 31, 28, 26, 32, 27, 3, 29, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 35,
	20, 25, 18, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 33, 3, 34,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 19, 21, 22, 23,
	24,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:52
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:57
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:61
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:74
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:78
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:83
		{
			yyVAL.Expr = &ast.FuncCallExpr{Name: yyDollar[1].String}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:87
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:94
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:98
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:105
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:110
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:117
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:121
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:126
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:134
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:139
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:143
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:149
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:153
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:157
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:161
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:165
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:172
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:176
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:185
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:189
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:193
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:197
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:201
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:205
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:209
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:213
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:217
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:221
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:226
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
