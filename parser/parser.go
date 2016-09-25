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
	Exprs   []ast.Expr
	String  string
	Strings []string
}

const INT = 57346
const BOOL = 57347
const STRING = 57348
const BREAK = 57349
const RETURN = 57350
const WORD = 57351
const BLANKSPACE = 57352
const LFCR = 57353
const FOR = 57354
const SWITCH = 57355
const DOUBLEAT = 57356
const FUNCTION = 57357
const CLASS = 57358
const DEF = 57359
const END = 57360
const DO = 57361
const POUNDCOMMENT = 57362
const DOUBLESLASHCOMMENT = 57363
const MULTILINECOMMENT = 57364
const LAMBDA = 57365
const DOUBLEADD = 57366
const DOUBLESUB = 57367
const IF = 57368
const ELSE = 57369
const AND = 57370
const OR = 57371
const GREATEREQUAL = 57372
const LESSEQUAL = 57373
const LOWER_SQUARE = 57374

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT",
	"BOOL",
	"STRING",
	"BREAK",
	"RETURN",
	"WORD",
	"BLANKSPACE",
	"LFCR",
	"FOR",
	"SWITCH",
	"DOUBLEAT",
	"FUNCTION",
	"CLASS",
	"DEF",
	"END",
	"DO",
	"POUNDCOMMENT",
	"DOUBLESLASHCOMMENT",
	"MULTILINECOMMENT",
	"LAMBDA",
	"DOUBLEADD",
	"DOUBLESUB",
	"IF",
	"ELSE",
	"'='",
	"AND",
	"OR",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'>'",
	"GREATEREQUAL",
	"'<'",
	"LESSEQUAL",
	"LOWER_SQUARE",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'.'",
	"','",
	"';'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:386

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 81,
	35, 0,
	36, 0,
	37, 0,
	38, 0,
	-2, 74,
	-1, 82,
	35, 0,
	36, 0,
	37, 0,
	38, 0,
	-2, 75,
	-1, 83,
	35, 0,
	36, 0,
	37, 0,
	38, 0,
	-2, 76,
	-1, 84,
	35, 0,
	36, 0,
	37, 0,
	38, 0,
	-2, 77,
}

const yyNprod = 99
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 576

var yyAct = [...]int{

	50, 20, 98, 20, 4, 6, 177, 125, 11, 8,
	179, 110, 176, 100, 121, 70, 32, 67, 2, 57,
	58, 49, 111, 55, 5, 25, 24, 26, 112, 108,
	32, 66, 20, 109, 63, 85, 64, 57, 58, 53,
	104, 56, 74, 164, 75, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 3, 65, 20, 38, 93, 56,
	147, 37, 121, 51, 120, 115, 121, 116, 20, 135,
	105, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 97, 25, 24, 26, 61, 55, 32, 185, 19,
	155, 20, 57, 58, 87, 85, 55, 32, 175, 53,
	119, 56, 101, 131, 56, 101, 90, 139, 127, 134,
	95, 129, 128, 161, 56, 141, 137, 142, 37, 101,
	51, 144, 38, 20, 131, 140, 106, 126, 148, 127,
	90, 152, 129, 128, 20, 149, 139, 139, 171, 86,
	150, 118, 89, 146, 122, 156, 157, 20, 126, 20,
	159, 20, 143, 89, 102, 99, 139, 117, 103, 145,
	52, 20, 20, 20, 160, 165, 162, 172, 32, 154,
	20, 101, 86, 68, 96, 62, 184, 60, 184, 169,
	180, 182, 180, 182, 113, 183, 184, 183, 187, 184,
	180, 182, 188, 180, 182, 183, 89, 94, 183, 187,
	181, 138, 181, 123, 69, 38, 45, 46, 47, 48,
	181, 73, 92, 181, 1, 38, 153, 38, 25, 24,
	26, 59, 36, 32, 38, 35, 28, 89, 7, 34,
	33, 27, 158, 13, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 130, 88, 89, 89, 10, 9,
	14, 166, 167, 136, 37, 89, 51, 18, 17, 91,
	173, 25, 24, 26, 15, 16, 32, 12, 0, 30,
	0, 0, 21, 23, 25, 24, 26, 15, 16, 32,
	22, 0, 30, 29, 0, 21, 23, 43, 44, 45,
	46, 47, 48, 22, 0, 0, 29, 37, 0, 31,
	174, 0, 0, 25, 24, 26, 15, 16, 32, 0,
	37, 30, 31, 170, 21, 23, 25, 24, 26, 15,
	16, 32, 22, 0, 30, 29, 0, 21, 23, 0,
	0, 0, 0, 0, 0, 22, 0, 0, 29, 37,
	0, 31, 168, 0, 0, 25, 24, 26, 15, 16,
	32, 0, 37, 30, 31, 163, 21, 23, 25, 24,
	26, 15, 16, 32, 22, 0, 30, 29, 0, 21,
	23, 0, 0, 0, 0, 0, 0, 22, 0, 0,
	29, 37, 0, 31, 114, 0, 0, 25, 24, 26,
	15, 16, 71, 0, 37, 30, 31, 107, 21, 23,
	25, 24, 26, 15, 16, 32, 22, 0, 30, 29,
	0, 21, 23, 25, 24, 26, 15, 16, 32, 22,
	0, 30, 29, 37, 0, 31, 68, 0, 0, 0,
	0, 0, 0, 0, 0, 29, 37, 0, 31, 25,
	24, 26, 15, 16, 32, 0, 0, 30, 0, 37,
	0, 31, 190, 0, 25, 24, 26, 15, 16, 32,
	0, 29, 30, 25, 24, 26, 15, 16, 32, 0,
	0, 30, 0, 0, 0, 37, 29, 31, 189, 15,
	16, 32, 0, 0, 30, 29, 0, 0, 23, 132,
	37, 0, 31, 186, 0, 25, 24, 26, 29, 37,
	32, 31, 178, 39, 40, 41, 42, 43, 44, 45,
	46, 47, 48, 0, 133, 151, 90, 15, 16, 32,
	0, 0, 30, 0, 57, 58, 23, 132, 55, 0,
	0, 37, 72, 51, 0, 0, 29, 25, 24, 26,
	54, 0, 32, 0, 53, 0, 56, 0, 0, 0,
	0, 0, 133, 124, 39, 40, 41, 42, 43, 44,
	45, 46, 47, 48, 0, 41, 42, 43, 44, 45,
	46, 47, 48, 37, 0, 51,
}
var yyPact = [...]int{

	396, -1000, 396, -1000, 525, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 533, 133, -1000, -1000,
	500, 168, 41, 166, -1000, -1000, -1000, -1000, -1000, 533,
	7, 383, -1000, -1000, -1000, -1000, -1000, 491, -1000, 533,
	533, 533, 533, 533, 533, 533, 533, 533, 533, 525,
	55, 130, 64, 214, 193, 396, 165, -1000, -1000, 37,
	-1000, 110, 112, 474, -8, 78, 58, 354, -1000, -14,
	-1000, -38, -1000, -19, 525, 534, 534, 254, 254, 171,
	171, -1000, -1000, -1000, -1000, 180, -38, -1000, -1000, -1000,
	341, -1000, 20, 525, 116, -1000, -1000, 96, 19, 102,
	-1000, -1000, 510, -1000, 21, 205, 88, -1000, -1000, 163,
	533, -1000, 533, 111, -1000, -1000, 533, 131, 101, 15,
	86, 162, 396, 472, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 68, 160, 396, 42, 159, 159, 64, -1000, 13,
	-1000, 525, 525, -1000, 525, 533, 396, 71, 396, -1000,
	312, -1000, -1000, -1, -1000, 159, 64, 64, -1000, 525,
	299, 396, 270, -1000, 93, 64, -1000, -1000, -1000, 257,
	-1000, 56, -33, -1000, -1000, 459, 46, 450, -1000, -1000,
	525, -1000, -1000, -1000, -5, 435, -1000, -1000, 409, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 54, 267, 4, 258, 257, 24, 9, 250, 5,
	249, 248, 94, 245, 89, 8, 7, 10, 244, 233,
	231, 230, 229, 228, 226, 225, 222, 15, 221, 216,
	13, 2, 0, 17, 214, 212, 211, 204, 203, 6,
}
var yyR1 = [...]int{

	0, 34, 33, 33, 1, 1, 1, 1, 1, 1,
	1, 1, 15, 15, 16, 16, 16, 16, 16, 38,
	38, 18, 18, 18, 18, 29, 17, 17, 17, 17,
	39, 39, 19, 19, 8, 8, 9, 32, 32, 10,
	10, 11, 11, 28, 31, 31, 30, 6, 6, 6,
	6, 6, 5, 5, 5, 5, 5, 4, 4, 12,
	12, 13, 14, 7, 7, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 2, 20, 20, 21, 22, 36, 36, 23,
	24, 24, 25, 26, 37, 37, 27, 35, 35,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 5, 4, 1, 1, 1, 1, 1, 1,
	2, 7, 6, 8, 7, 1, 1, 1, 1, 1,
	1, 2, 3, 4, 1, 4, 3, 1, 3, 7,
	8, 7, 6, 1, 1, 3, 1, 1, 2, 1,
	1, 1, 7, 6, 5, 4, 6, 3, 3, 1,
	1, 2, 3, 2, 2, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 6,
	1, 1, 2, 3, 1, 3, 3, 1, 3,
}
var yyChk = [...]int{

	-1000, -34, -33, -1, -3, -6, -9, -23, -7, -10,
	-11, -15, -2, -19, -8, 7, 8, -4, -5, -14,
	-32, 15, 23, 16, 5, 4, 6, -20, -24, 26,
	12, 42, 9, -21, -22, -25, -26, 40, -1, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, -3,
	-32, 42, 27, 44, 40, 28, 46, 24, 25, -28,
	9, 44, 9, -3, -9, 48, -32, -33, 43, -37,
	-27, 9, 41, -36, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, 40, 9, -12, -13, -14,
	42, 45, -35, -3, 4, -1, 9, 44, -31, 45,
	-30, 9, 42, -12, 48, -3, 48, 43, 43, 47,
	49, 41, 47, 4, 43, 45, 47, 41, 45, -31,
	45, 47, 42, -38, 43, -16, -6, -9, -7, -15,
	-18, -32, 17, 42, -3, 48, 48, -7, -12, -32,
	-27, -3, -3, 41, -3, 28, 42, 45, 42, -30,
	-33, 43, -16, -29, 9, 48, -7, -7, -12, -3,
	-33, 42, -33, 43, 44, -7, -12, -12, 43, -33,
	43, 45, -31, -12, 43, 42, 45, -39, 43, -17,
	-3, -6, -9, -7, -32, 42, 43, -17, -39, 43,
	43,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	10, 11, 65, 66, 67, 47, 0, 49, 50, 51,
	34, 0, 0, 0, 78, 79, 80, 81, 82, 0,
	0, 0, 37, 83, 84, 90, 91, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 48,
	34, 0, 0, 0, 0, 0, 0, 63, 64, 0,
	43, 0, 0, 0, 0, 0, 0, 0, 92, 0,
	94, 37, 85, 0, 87, 68, 69, 70, 71, 72,
	73, -2, -2, -2, -2, 0, 0, 58, 59, 60,
	0, 32, 0, 97, 0, 36, 38, 0, 0, 0,
	44, 46, 0, 57, 0, 0, 0, 62, 93, 0,
	0, 86, 0, 0, 61, 33, 0, 35, 0, 0,
	0, 0, 0, 0, 13, 19, 14, 15, 16, 17,
	18, 0, 0, 0, 0, 0, 0, 0, 55, 0,
	95, 96, 88, 35, 98, 0, 0, 0, 0, 45,
	0, 12, 20, 0, 25, 0, 0, 0, 54, 89,
	0, 0, 0, 42, 0, 0, 56, 53, 39, 0,
	41, 0, 0, 52, 40, 0, 0, 0, 22, 30,
	26, 27, 28, 29, 34, 0, 21, 31, 0, 24,
	23,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	44, 45, 33, 31, 47, 32, 46, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 48,
	37, 28, 35, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 40, 3, 41, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 42, 3, 43,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 29, 30, 36, 38,
	39,
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
		//line parser/parser.go.y:65
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:70
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:74
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:92
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:96
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:105
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:109
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:114
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:118
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, nil}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:122
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:126
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:136
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:140
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:146
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, nil}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:150
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, yyDollar[3].Exprs}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:155
		{
			yyVAL.Expr = &ast.GetExpr{yyDollar[1].Strings}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:159
		{
			yyVAL.Expr = &ast.ArrayGetExpr{yyDollar[1].Strings, yyDollar[3].Expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:164
		{
			yyVAL.Expr = &ast.SetExpr{yyDollar[1].Strings, yyDollar[3].Expr}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:170
		{
			yyVAL.Strings = append(make([]string, 0, 16), yyDollar[1].String)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:174
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:181
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:185
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 41:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:190
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{yyDollar[3].Strings, yyDollar[6].Exprs}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:194
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{nil, yyDollar[5].Exprs}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:201
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:205
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:212
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:216
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 52:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:224
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:228
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:232
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:236
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 56:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:240
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:245
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:249
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:259
		{
			yyVAL.Expr = &ast.BlockExpr{nil}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:263
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:270
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Strings}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:274
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Strings}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:283
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:287
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:291
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:295
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:299
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:303
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:307
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:311
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:315
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:319
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:329
		{
			yyVAL.Expr = &ast.ArrayExpr{nil}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:333
		{
			yyVAL.Expr = &ast.ArrayExpr{yyDollar[2].Exprs}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:337
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0, 1024), yyDollar[1].Expr)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:341
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 89:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:346
		{
			yyVAL.Expr = &ast.ArraySetExpr{yyDollar[1].Strings, yyDollar[3].Expr, yyDollar[6].Expr}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:353
		{
			yyVAL.Expr = &ast.ObjectExpr{nil}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:358
		{
			yyVAL.Expr = &ast.ObjectExpr{yyDollar[2].Exprs}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:363
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0, 1024), yyDollar[1].Expr)
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:367
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:373
		{
			yyVAL.Expr = &ast.KeyValueExpr{yyDollar[1].String, yyDollar[3].Expr}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:378
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:383
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	}
	goto yystack /* stack new state and value */
}
