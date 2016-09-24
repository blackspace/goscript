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

const NUMBER = 57346
const BOOL = 57347
const STRING = 57348
const BREAK = 57349
const RETURN = 57350
const WORD = 57351
const BLANKSPACE = 57352
const LFCR = 57353
const IF = 57354
const FOR = 57355
const SWITCH = 57356
const DOUBLEAT = 57357
const FUNCTION = 57358
const CLASS = 57359
const DEF = 57360
const END = 57361
const DO = 57362
const POUNDCOMMENT = 57363
const DOUBLESLASHCOMMENT = 57364
const MULTILINECOMMENT = 57365
const LAMBDA = 57366
const GREATEREQUAL = 57367
const LESSEQUAL = 57368
const ELSE = 57369
const AND = 57370
const OR = 57371
const DOUBLEADD = 57372
const DOUBLESUB = 57373

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"BOOL",
	"STRING",
	"BREAK",
	"RETURN",
	"WORD",
	"BLANKSPACE",
	"LFCR",
	"IF",
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
	"'>'",
	"GREATEREQUAL",
	"'<'",
	"LESSEQUAL",
	"ELSE",
	"AND",
	"OR",
	"DOUBLEADD",
	"DOUBLESUB",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"','",
	"';'",
	"'['",
	"']'",
	"':'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:390

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 82,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 74,
	-1, 83,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 75,
	-1, 84,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 76,
	-1, 85,
	25, 0,
	26, 0,
	27, 0,
	28, 0,
	-2, 77,
}

const yyNprod = 100
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 658

var yyAct = [...]int{

	51, 21, 99, 21, 4, 6, 178, 11, 71, 8,
	180, 126, 111, 57, 58, 55, 101, 68, 2, 144,
	56, 118, 50, 54, 5, 105, 113, 86, 33, 112,
	165, 56, 67, 21, 54, 64, 102, 65, 86, 177,
	122, 148, 122, 75, 98, 76, 77, 78, 79, 80,
	81, 82, 83, 84, 85, 62, 21, 109, 3, 94,
	110, 39, 186, 33, 66, 176, 20, 87, 88, 21,
	172, 106, 102, 46, 47, 48, 49, 146, 40, 41,
	121, 122, 91, 42, 43, 44, 45, 102, 57, 58,
	55, 162, 21, 156, 91, 56, 116, 117, 54, 69,
	149, 120, 59, 53, 132, 147, 119, 123, 140, 128,
	135, 130, 103, 129, 95, 33, 142, 138, 143, 141,
	90, 100, 145, 155, 21, 132, 102, 39, 127, 87,
	128, 90, 130, 104, 129, 21, 153, 140, 140, 150,
	96, 151, 42, 43, 44, 45, 157, 158, 21, 127,
	21, 160, 21, 40, 41, 44, 45, 140, 42, 43,
	44, 45, 21, 21, 21, 161, 166, 163, 173, 63,
	55, 21, 61, 114, 90, 56, 139, 185, 97, 185,
	170, 181, 183, 181, 183, 124, 184, 185, 184, 188,
	185, 181, 183, 189, 181, 183, 184, 57, 58, 184,
	188, 182, 70, 182, 56, 90, 74, 159, 93, 1,
	39, 182, 154, 60, 182, 37, 57, 58, 55, 36,
	39, 29, 39, 56, 90, 90, 167, 168, 35, 39,
	34, 28, 13, 90, 131, 174, 26, 25, 27, 16,
	17, 33, 89, 10, 30, 31, 9, 14, 22, 24,
	26, 25, 27, 16, 17, 33, 23, 19, 30, 31,
	18, 7, 22, 24, 15, 12, 0, 0, 0, 0,
	23, 0, 32, 175, 0, 0, 0, 0, 38, 0,
	0, 0, 0, 0, 0, 0, 32, 171, 0, 0,
	0, 0, 38, 26, 25, 27, 16, 17, 33, 0,
	0, 30, 31, 0, 0, 22, 24, 26, 25, 27,
	16, 17, 33, 23, 0, 30, 31, 0, 0, 22,
	24, 0, 0, 0, 0, 0, 0, 23, 0, 32,
	169, 0, 0, 0, 0, 38, 0, 0, 0, 0,
	0, 0, 0, 32, 164, 0, 0, 0, 0, 38,
	26, 25, 27, 16, 17, 33, 0, 0, 30, 31,
	0, 0, 22, 24, 26, 25, 27, 16, 17, 33,
	23, 0, 30, 31, 0, 0, 22, 24, 0, 0,
	0, 0, 0, 0, 23, 0, 32, 115, 0, 0,
	0, 0, 38, 0, 0, 0, 0, 0, 0, 0,
	32, 108, 0, 0, 0, 0, 38, 26, 25, 27,
	16, 17, 72, 0, 0, 30, 31, 0, 0, 22,
	24, 26, 25, 27, 16, 17, 33, 23, 0, 30,
	31, 0, 0, 22, 24, 0, 0, 0, 0, 0,
	0, 23, 0, 32, 69, 0, 0, 0, 0, 38,
	26, 25, 27, 16, 17, 33, 0, 32, 30, 31,
	0, 0, 0, 38, 26, 25, 27, 16, 17, 33,
	0, 0, 30, 31, 26, 25, 27, 16, 17, 33,
	0, 0, 30, 31, 0, 0, 32, 191, 0, 0,
	0, 0, 38, 0, 26, 25, 27, 16, 17, 33,
	32, 190, 30, 31, 0, 0, 38, 0, 0, 0,
	32, 187, 46, 47, 48, 49, 38, 40, 41, 0,
	0, 0, 42, 43, 44, 45, 0, 0, 0, 0,
	32, 179, 137, 26, 25, 27, 38, 0, 33, 0,
	0, 26, 25, 27, 0, 0, 33, 0, 26, 25,
	27, 0, 0, 33, 0, 26, 25, 27, 0, 0,
	33, 0, 26, 25, 27, 0, 0, 33, 0, 52,
	0, 0, 0, 0, 0, 38, 73, 52, 0, 0,
	0, 0, 136, 38, 52, 0, 0, 0, 0, 107,
	38, 52, 0, 0, 92, 0, 0, 38, 52, 0,
	46, 47, 48, 49, 38, 40, 41, 0, 0, 0,
	42, 43, 44, 45, 0, 91, 16, 17, 33, 0,
	0, 30, 31, 16, 17, 33, 24, 133, 30, 31,
	0, 0, 0, 24, 133, 46, 47, 48, 49, 0,
	40, 41, 0, 0, 0, 42, 43, 44, 45, 134,
	152, 0, 0, 0, 0, 0, 134, 125,
}
var yyPact = [...]int{

	417, -1000, 417, -1000, 610, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 558, 74, -1000,
	-1000, 56, 163, 13, 160, -1000, -1000, -1000, -1000, -1000,
	558, 19, 403, -1000, -1000, -1000, -1000, -1000, 529, -1000,
	558, 558, 558, 558, 558, 558, 558, 558, 558, 558,
	610, -8, 58, 42, 551, 417, 131, -1000, -1000, 174,
	2, -1000, 78, 72, 575, -20, 544, 136, 360, -1000,
	16, -1000, -36, -1000, -18, 610, 107, 107, 118, 118,
	-1000, -1000, 123, 123, 123, 123, 169, -36, -1000, -1000,
	-1000, 346, -1000, 53, 610, -1000, -1000, -26, 63, 37,
	67, -1000, -1000, 616, -1000, 537, 487, 54, -1000, -1000,
	120, 558, -1000, 558, -28, -1000, -1000, 558, 43, 65,
	-2, 60, 117, 417, 609, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 184, 114, 417, 48, 106, 106, 42, -1000,
	165, -1000, 610, 610, -1000, 610, 558, 417, 51, 417,
	-1000, 303, -1000, -1000, -12, -1000, 106, 42, 42, -1000,
	610, 289, 417, 246, -1000, 27, 42, -1000, -1000, -1000,
	232, -1000, 25, -4, -1000, -1000, 490, 22, 470, -1000,
	-1000, 610, -1000, -1000, -1000, -19, 460, -1000, -1000, 446,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 58, 265, 4, 264, 261, 260, 257, 24, 9,
	247, 5, 246, 243, 68, 242, 66, 7, 11, 10,
	234, 232, 231, 230, 228, 221, 219, 215, 8, 213,
	212, 16, 2, 0, 17, 209, 208, 206, 202, 185,
	6,
}
var yyR1 = [...]int{

	0, 35, 34, 34, 1, 1, 1, 1, 1, 1,
	1, 1, 17, 17, 18, 18, 18, 18, 18, 39,
	39, 20, 20, 20, 20, 30, 19, 19, 19, 19,
	40, 40, 21, 21, 10, 11, 33, 33, 12, 12,
	13, 13, 29, 32, 32, 31, 8, 8, 8, 8,
	8, 7, 7, 7, 7, 7, 6, 6, 14, 14,
	15, 16, 9, 9, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 2, 22, 22, 23, 24, 37, 37, 4,
	5, 25, 25, 26, 27, 38, 38, 28, 36, 36,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 5, 4, 1, 1, 1, 1, 1, 1,
	2, 7, 6, 8, 7, 1, 1, 1, 1, 1,
	1, 2, 3, 4, 1, 3, 1, 3, 7, 8,
	7, 6, 1, 1, 3, 1, 1, 2, 1, 1,
	1, 7, 6, 5, 4, 6, 3, 3, 1, 1,
	2, 3, 2, 2, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 2, 3, 1, 3, 4,
	6, 1, 1, 2, 3, 1, 3, 3, 1, 3,
}
var yyChk = [...]int{

	-1000, -35, -34, -1, -3, -8, -11, -5, -9, -12,
	-13, -17, -2, -21, -10, -4, 7, 8, -6, -7,
	-16, -33, 16, 24, 17, 5, 4, 6, -22, -25,
	12, 13, 40, 9, -23, -24, -26, -27, 46, -1,
	30, 31, 35, 36, 37, 38, 25, 26, 27, 28,
	-3, -33, 40, 29, 42, 34, 39, 32, 33, 46,
	-29, 9, 42, 9, -3, -11, 45, -33, -34, 41,
	-38, -28, 9, 47, -37, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, 46, 9, -14, -15,
	-16, 40, 43, -36, -3, -1, 9, 4, 42, -32,
	43, -31, 9, 40, -14, 45, -3, 45, 41, 41,
	44, 48, 47, 44, 4, 41, 43, 44, 47, 43,
	-32, 43, 44, 40, -39, 41, -18, -8, -11, -9,
	-17, -20, -33, 18, 40, -3, 45, 45, -9, -14,
	-33, -28, -3, -3, 47, -3, 34, 40, 43, 40,
	-31, -34, 41, -18, -30, 9, 45, -9, -9, -14,
	-3, -34, 40, -34, 41, 42, -9, -14, -14, 41,
	-34, 41, 43, -32, -14, 41, 40, 43, -40, 41,
	-19, -3, -8, -11, -9, -33, 40, 41, -19, -40,
	41, 41,
}
var yyDef = [...]int{

	0, -2, 1, 2, 4, 5, 6, 7, 8, 9,
	10, 11, 64, 65, 66, 67, 46, 0, 48, 49,
	50, 34, 0, 0, 0, 78, 79, 80, 81, 82,
	0, 0, 0, 36, 83, 84, 91, 92, 0, 3,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	47, 34, 0, 0, 0, 0, 0, 62, 63, 0,
	0, 42, 0, 0, 0, 0, 0, 0, 0, 93,
	0, 95, 36, 85, 0, 87, 68, 69, 70, 71,
	72, 73, -2, -2, -2, -2, 0, 0, 57, 58,
	59, 0, 32, 0, 98, 35, 37, 0, 0, 0,
	0, 43, 45, 0, 56, 0, 0, 0, 61, 94,
	0, 0, 86, 0, 0, 60, 33, 0, 89, 0,
	0, 0, 0, 0, 0, 13, 19, 14, 15, 16,
	17, 18, 0, 0, 0, 0, 0, 0, 0, 54,
	0, 96, 97, 88, 89, 99, 0, 0, 0, 0,
	44, 0, 12, 20, 0, 25, 0, 0, 0, 53,
	90, 0, 0, 0, 41, 0, 0, 55, 52, 38,
	0, 40, 0, 0, 51, 39, 0, 0, 0, 22,
	30, 26, 27, 28, 29, 34, 0, 21, 31, 0,
	24, 23,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	42, 43, 37, 35, 44, 36, 39, 38, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 48, 45,
	27, 34, 25, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 46, 3, 47, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 40, 3, 41,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 26, 28, 29, 30, 31, 32, 33,
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
		//line parser/parser.go.y:66
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:71
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:75
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:93
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:97
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:110
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 21:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:115
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:119
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, nil, nil}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:123
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 24:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:127
		{
			yyVAL.Expr = &ast.MethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:137
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:141
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:147
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, nil}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:151
		{
			yyVAL.Expr = &ast.CalledExpr{yyDollar[1].Strings, yyDollar[3].Exprs}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:156
		{
			yyVAL.Expr = &ast.GetExpr{yyDollar[1].Strings}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:161
		{
			yyVAL.Expr = &ast.SetExpr{yyDollar[1].Strings, yyDollar[3].Expr}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:167
		{
			yyVAL.Strings = append(make([]string, 0, 16), yyDollar[1].String)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:171
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 38:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:178
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 39:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:182
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:187
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{yyDollar[3].Strings, yyDollar[6].Exprs}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:191
		{
			yyVAL.Expr = &ast.LambdaDefineExpr{nil, yyDollar[5].Exprs}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:198
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:202
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:211
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:215
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 51:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:223
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:227
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:231
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:235
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:239
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:244
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:248
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:258
		{
			yyVAL.Expr = &ast.BlockExpr{nil}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:262
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:269
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Strings}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:273
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Strings}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:282
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:286
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:290
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:294
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:298
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:302
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:306
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:310
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:314
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:318
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:328
		{
			yyVAL.Expr = &ast.ArrayExpr{nil}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:332
		{
			yyVAL.Expr = &ast.ArrayExpr{yyDollar[2].Exprs}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:336
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0, 1024), yyDollar[1].Expr)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:340
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:345
		{
			yyVAL.Expr = &ast.ArrayGetExpr{yyDollar[1].Strings, yyDollar[3].Expr}
		}
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:350
		{
			yyVAL.Expr = &ast.ArraySetExpr{yyDollar[1].Strings, yyDollar[3].Expr, yyDollar[6].Expr}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:357
		{
			yyVAL.Expr = &ast.ObjectExpr{nil}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:362
		{
			yyVAL.Expr = &ast.ObjectExpr{yyDollar[2].Exprs}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:367
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0, 1024), yyDollar[1].Expr)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:371
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:377
		{
			yyVAL.Expr = &ast.KeyValueExpr{yyDollar[1].String, yyDollar[3].Expr}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:382
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:387
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	}
	goto yystack /* stack new state and value */
}
