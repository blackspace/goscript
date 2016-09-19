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
const DOUBLEADD = 57357
const DOUBLESUB = 57358
const DOUBLEAT = 57359
const FUNCTION = 57360
const CLASS = 57361
const DEF = 57362
const END = 57363
const DO = 57364
const GREATEREQUAL = 57365
const LESSEQUAL = 57366
const ELSE = 57367
const AND = 57368
const OR = 57369

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
	"DOUBLEADD",
	"DOUBLESUB",
	"DOUBLEAT",
	"FUNCTION",
	"CLASS",
	"DEF",
	"END",
	"DO",
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
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'.'",
	"','",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:313

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	37, 48,
	39, 35,
	-2, 87,
	-1, 67,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 81,
	-1, 68,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 82,
	-1, 69,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 83,
	-1, 70,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 84,
	-1, 76,
	37, 33,
	-2, 34,
}

const yyNprod = 88
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 447

var yyAct = [...]int{

	4, 45, 15, 92, 15, 93, 142, 82, 144, 97,
	74, 81, 60, 11, 141, 130, 121, 109, 72, 10,
	44, 85, 120, 9, 121, 108, 15, 109, 94, 56,
	48, 59, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 133, 15, 58, 94, 13, 12, 14, 6,
	7, 26, 89, 15, 77, 55, 83, 137, 60, 86,
	151, 13, 12, 14, 140, 24, 26, 73, 79, 8,
	46, 22, 60, 22, 91, 90, 84, 35, 36, 57,
	98, 15, 3, 111, 24, 30, 110, 5, 41, 115,
	83, 47, 15, 105, 94, 22, 98, 15, 87, 104,
	2, 117, 27, 103, 88, 123, 76, 114, 54, 105,
	83, 119, 22, 115, 115, 104, 50, 125, 124, 103,
	42, 43, 22, 95, 71, 51, 49, 132, 115, 100,
	101, 42, 43, 129, 30, 41, 1, 138, 75, 53,
	131, 145, 15, 145, 15, 100, 101, 135, 136, 102,
	22, 153, 145, 15, 139, 145, 15, 113, 154, 106,
	150, 22, 150, 153, 17, 102, 22, 99, 16, 21,
	20, 150, 0, 118, 150, 13, 12, 14, 0, 0,
	26, 127, 128, 99, 13, 12, 14, 0, 0, 26,
	147, 148, 147, 148, 31, 32, 134, 33, 34, 35,
	36, 147, 148, 0, 147, 148, 0, 0, 0, 116,
	149, 22, 149, 22, 33, 34, 35, 36, 80, 0,
	0, 149, 22, 0, 149, 22, 0, 0, 146, 0,
	146, 0, 0, 0, 0, 37, 38, 39, 40, 146,
	31, 32, 146, 33, 34, 35, 36, 0, 0, 37,
	38, 39, 40, 126, 31, 32, 0, 33, 34, 35,
	36, 13, 12, 14, 18, 19, 26, 112, 0, 28,
	29, 0, 0, 0, 0, 23, 25, 107, 0, 0,
	13, 12, 14, 18, 19, 26, 0, 0, 28, 29,
	0, 0, 24, 122, 23, 25, 107, 0, 0, 13,
	12, 14, 18, 19, 26, 0, 0, 28, 29, 0,
	0, 24, 96, 23, 25, 0, 0, 0, 13, 12,
	14, 18, 19, 26, 0, 0, 28, 29, 0, 0,
	24, 78, 23, 25, 13, 12, 14, 18, 19, 26,
	0, 0, 28, 29, 0, 0, 0, 0, 0, 24,
	52, 13, 12, 14, 18, 19, 26, 0, 0, 28,
	29, 0, 0, 0, 0, 24, 156, 13, 12, 14,
	18, 19, 26, 0, 0, 28, 29, 0, 0, 0,
	0, 0, 24, 155, 13, 12, 14, 18, 19, 26,
	0, 0, 28, 29, 0, 13, 12, 14, 24, 152,
	26, 0, 0, 13, 12, 14, 18, 19, 26, 0,
	0, 28, 29, 0, 0, 24, 143, 23, 25, 37,
	38, 39, 40, 0, 31, 32, 0, 33, 34, 35,
	36, 24, 0, 0, 24, 37, 38, 39, 40, 0,
	31, 32, 0, 33, 34, 35, 36,
}
var yyPact = [...]int{

	399, -1000, 399, -1000, 412, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 105, -1000, -1000, -1000, 391,
	64, -1000, -9, 107, 314, 99, -1000, 18, 391, 3,
	-1000, 391, 391, 391, 391, 391, 391, 391, 391, 391,
	391, 399, -1000, -1000, 412, -1000, -21, 30, 97, 17,
	-1000, 295, -1000, 33, -1000, 180, 396, -20, 57, 58,
	-1000, 183, 183, 44, 44, -1000, -1000, 166, 166, 166,
	166, -1000, 95, -1000, 15, 45, -1000, 36, -1000, 276,
	-1000, -13, -1000, 412, -1000, 42, 226, 49, -1000, 171,
	399, 30, -16, -1000, -1000, 257, -1000, -1000, 412, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 95, -1000, 391,
	212, 63, 63, 30, -1000, 116, -1000, -23, -1000, -1000,
	30, 85, -1000, -1000, 5, -1000, 63, 30, 30, -1000,
	-1000, -1000, -1000, 19, 30, -1000, -1000, 29, -24, -1000,
	380, 25, 363, -1000, -1000, 412, -1000, -1000, -1000, -1000,
	-1000, 347, -1000, -1000, 330, -1000, -1000,
}
var yyPgo = [...]int{

	0, 82, 49, 0, 1, 170, 169, 87, 19, 69,
	23, 168, 7, 13, 9, 8, 164, 50, 159, 102,
	139, 10, 138, 70, 5, 3, 100, 136, 11, 123,
	6,
}
var yyR1 = [...]int{

	0, 27, 26, 26, 13, 13, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 29, 29, 18, 18, 18,
	18, 15, 15, 15, 15, 15, 15, 30, 30, 16,
	16, 17, 20, 21, 22, 23, 1, 1, 1, 1,
	1, 1, 1, 1, 10, 10, 11, 11, 19, 25,
	25, 24, 28, 28, 12, 7, 7, 7, 7, 2,
	9, 9, 6, 6, 6, 6, 6, 5, 5, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 8, 8, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 7, 6, 8,
	7, 1, 1, 1, 1, 1, 1, 1, 2, 5,
	6, 5, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 5, 6, 3, 4, 1, 1,
	3, 1, 1, 3, 1, 1, 2, 1, 1, 3,
	2, 2, 7, 6, 5, 4, 6, 3, 3, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 1,
}
var yyChk = [...]int{

	-1000, -27, -26, -1, -3, -7, -2, -17, -9, -10,
	-8, -13, 5, 4, 6, -4, -11, -16, 7, 8,
	-5, -6, -23, 18, 35, 19, 9, -19, 12, 13,
	-1, 28, 29, 31, 32, 33, 34, 23, 24, 25,
	26, 30, 15, 16, -3, -4, -23, 27, 39, -19,
	9, -26, 36, -20, 9, 37, -3, -2, 41, -4,
	9, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -1, 39, -8, -21, -22, 9, 37, 36, 35,
	38, -28, -12, -3, -8, 41, -3, 41, 9, 37,
	30, 38, -25, -24, 9, -29, 36, -14, -3, -7,
	-2, -17, -9, -10, -8, -13, -18, 20, 38, 40,
	-3, 41, 41, -9, -8, -4, 38, -28, -1, -8,
	38, 40, 36, -14, -21, -12, 41, -9, -9, -8,
	38, -8, -24, 37, -9, -8, -8, 38, -25, -8,
	35, 38, -30, 36, -15, -3, -7, -2, -17, -9,
	-8, 35, 36, -15, -30, 36, 36,
}
var yyDef = [...]int{

	0, -2, 1, 2, 36, 37, 38, 39, 40, 41,
	42, 43, 69, 70, 71, 72, 73, 74, 55, 0,
	57, 58, 0, 0, 0, 0, -2, 0, 0, 0,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 60, 61, 56, 72, 0, 0, 0, 0,
	48, 0, 86, 0, 32, 0, 0, 0, 0, 0,
	87, 75, 76, 77, 78, 79, 80, -2, -2, -2,
	-2, 59, 0, 68, 0, 0, -2, 0, 85, 0,
	46, 0, 52, 54, 67, 0, 0, 0, 33, 0,
	0, 0, 0, 49, 51, 0, 5, 15, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 0, 47, 0,
	0, 0, 0, 0, 65, 0, 29, 0, 31, 44,
	0, 0, 4, 16, 0, 53, 0, 0, 0, 64,
	30, 45, 50, 0, 0, 66, 63, 0, 0, 62,
	0, 0, 0, 18, 27, 21, 22, 23, 24, 25,
	26, 0, 17, 28, 0, 20, 19,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	37, 38, 33, 31, 40, 32, 39, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 41,
	25, 30, 23, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 3, 36,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 24, 26, 27, 28, 29,
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
		//line parser/parser.go.y:56
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:61
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:65
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:72
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:76
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:85
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:89
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:94
		{
			yyVAL.Expr = &ast.ObjectMethod{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:98
		{
			yyVAL.Expr = &ast.ObjectMethod{yyDollar[2].String, nil, nil}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:102
		{
			yyVAL.Expr = &ast.ObjectMethod{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.Expr = &ast.ObjectMethod{yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:114
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:118
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:124
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, nil}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:128
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, yyDollar[5].Exprs}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:132
		{

		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:154
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:158
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:163
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, nil}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:167
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:174
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:178
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:185
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:190
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:198
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:202
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:210
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:215
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:219
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 62:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:225
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 63:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:229
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:233
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:237
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:241
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:246
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:250
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:260
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:264
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:268
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:272
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:276
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:280
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:284
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:288
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:292
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:296
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:301
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:305
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:310
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
