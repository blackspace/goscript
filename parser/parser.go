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

//line parser/parser.go.y:319

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	37, 49,
	39, 36,
	-2, 89,
	-1, 68,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 83,
	-1, 69,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 84,
	-1, 70,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 85,
	-1, 71,
	23, 0,
	24, 0,
	25, 0,
	26, 0,
	-2, 86,
	-1, 77,
	37, 34,
	-2, 35,
}

const yyNprod = 90
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 463

var yyAct = [...]int{

	4, 46, 15, 93, 15, 144, 146, 83, 75, 94,
	98, 76, 143, 11, 122, 82, 132, 61, 110, 10,
	86, 45, 9, 7, 8, 73, 121, 15, 122, 109,
	57, 110, 60, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 6, 15, 49, 13, 12, 14, 59,
	5, 27, 135, 90, 15, 95, 78, 84, 56, 95,
	87, 38, 39, 40, 41, 153, 32, 33, 74, 34,
	35, 36, 37, 61, 58, 61, 142, 85, 25, 128,
	80, 99, 15, 112, 139, 89, 91, 111, 92, 42,
	116, 84, 48, 15, 106, 36, 37, 99, 15, 3,
	105, 25, 31, 104, 102, 103, 118, 124, 115, 2,
	106, 84, 120, 114, 116, 116, 105, 125, 127, 104,
	102, 103, 47, 23, 101, 23, 34, 35, 36, 37,
	116, 100, 134, 95, 131, 52, 28, 129, 130, 140,
	101, 133, 72, 147, 15, 147, 15, 100, 23, 137,
	138, 155, 31, 136, 147, 15, 141, 147, 15, 156,
	126, 50, 152, 155, 152, 23, 150, 151, 150, 151,
	43, 44, 77, 152, 55, 23, 152, 150, 151, 51,
	150, 151, 43, 44, 96, 42, 149, 1, 149, 54,
	18, 119, 107, 148, 17, 148, 16, 149, 22, 21,
	149, 0, 0, 23, 148, 0, 0, 148, 13, 12,
	14, 0, 0, 27, 23, 38, 39, 40, 41, 23,
	32, 33, 0, 34, 35, 36, 37, 13, 12, 14,
	19, 20, 27, 113, 0, 29, 30, 0, 0, 0,
	0, 24, 26, 108, 0, 88, 13, 12, 14, 19,
	20, 27, 0, 0, 29, 30, 0, 0, 25, 123,
	24, 26, 108, 0, 0, 23, 0, 23, 32, 33,
	0, 34, 35, 36, 37, 0, 23, 25, 97, 23,
	13, 12, 14, 19, 20, 27, 0, 0, 29, 30,
	0, 13, 12, 14, 24, 26, 27, 0, 0, 13,
	12, 14, 19, 20, 27, 0, 0, 29, 30, 0,
	0, 25, 79, 24, 26, 13, 12, 14, 19, 20,
	27, 0, 0, 29, 30, 117, 0, 0, 0, 0,
	25, 53, 13, 12, 14, 19, 20, 27, 0, 0,
	29, 30, 0, 0, 0, 0, 25, 158, 13, 12,
	14, 19, 20, 27, 0, 0, 29, 30, 0, 0,
	0, 0, 0, 25, 157, 13, 12, 14, 19, 20,
	27, 0, 0, 29, 30, 0, 13, 12, 14, 25,
	154, 27, 0, 0, 13, 12, 14, 19, 20, 27,
	0, 0, 29, 30, 0, 0, 25, 145, 24, 26,
	38, 39, 40, 41, 0, 32, 33, 0, 34, 35,
	36, 37, 25, 0, 0, 25, 38, 39, 40, 41,
	0, 32, 33, 0, 34, 35, 36, 37, 13, 12,
	14, 0, 0, 27, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 81,
}
var yyPact = [...]int{

	380, -1000, 380, -1000, 393, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 155, -1000, -1000, -1000, -1000,
	372, 65, -1000, 6, 170, 295, 165, -1000, 21, 372,
	8, -1000, 372, 372, 372, 372, 372, 372, 372, 372,
	372, 372, 380, -1000, -1000, 393, -1000, -14, 43, 163,
	19, -1000, 276, -1000, 45, -1000, 424, 377, -21, 204,
	59, -1000, 95, 95, 62, 62, -1000, -1000, 240, 240,
	240, 240, -1000, 163, -1000, 16, 56, -1000, 50, -1000,
	242, -1000, -9, -1000, 393, -1000, 42, 192, 66, -1000,
	287, 380, 43, -12, -1000, -1000, 223, -1000, -1000, 393,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 151, -1000,
	372, 38, 64, 64, 43, -1000, 167, -1000, -22, -1000,
	-1000, 43, 124, -1000, -1000, 15, -1000, -1000, 64, 43,
	43, -1000, -1000, -1000, -1000, 46, 43, -1000, -1000, 41,
	-26, -1000, 361, 30, 344, -1000, -1000, 393, -1000, -1000,
	-1000, -1000, -1000, 328, -1000, -1000, 311, -1000, -1000,
}
var yyPgo = [...]int{

	0, 99, 43, 0, 1, 199, 198, 50, 19, 24,
	22, 196, 7, 13, 10, 6, 194, 192, 23, 190,
	136, 189, 8, 11, 122, 9, 3, 109, 187, 15,
	184, 5,
}
var yyR1 = [...]int{

	0, 28, 27, 27, 13, 13, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 30, 30, 17, 17, 17,
	17, 15, 15, 15, 15, 15, 15, 31, 31, 16,
	16, 18, 19, 21, 22, 23, 24, 1, 1, 1,
	1, 1, 1, 1, 1, 10, 10, 11, 11, 20,
	26, 26, 25, 29, 29, 12, 7, 7, 7, 7,
	2, 9, 9, 6, 6, 6, 6, 6, 5, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 8, 8, 4,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 7, 6, 8,
	7, 1, 1, 1, 1, 1, 1, 1, 2, 5,
	6, 5, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 5, 6, 3, 4, 1,
	1, 3, 1, 1, 3, 1, 1, 2, 1, 1,
	3, 2, 2, 7, 6, 5, 4, 6, 3, 3,
	1, 1, 1, 1, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
}
var yyChk = [...]int{

	-1000, -28, -27, -1, -3, -7, -2, -18, -9, -10,
	-8, -13, 5, 4, 6, -4, -11, -16, -19, 7,
	8, -5, -6, -24, 18, 35, 19, 9, -20, 12,
	13, -1, 28, 29, 31, 32, 33, 34, 23, 24,
	25, 26, 30, 15, 16, -3, -4, -24, 27, 39,
	-20, 9, -27, 36, -21, 9, 37, -3, -2, 41,
	-4, 9, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -1, 39, -8, -22, -23, 9, 37, 36,
	35, 38, -29, -12, -3, -8, 41, -3, 41, -23,
	37, 30, 38, -26, -25, 9, -30, 36, -14, -3,
	-7, -2, -18, -9, -10, -8, -13, -17, 20, 38,
	40, -3, 41, 41, -9, -8, -4, 38, -29, -1,
	-8, 38, 40, 36, -14, -22, 9, -12, 41, -9,
	-9, -8, 38, -8, -25, 37, -9, -8, -8, 38,
	-26, -8, 35, 38, -31, 36, -15, -3, -7, -2,
	-18, -9, -8, 35, 36, -15, -31, 36, 36,
}
var yyDef = [...]int{

	0, -2, 1, 2, 37, 38, 39, 40, 41, 42,
	43, 44, 70, 71, 72, 73, 74, 75, 76, 56,
	0, 58, 59, 0, 0, 0, 0, -2, 0, 0,
	0, 3, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 61, 62, 57, 73, 0, 0, 0,
	0, 49, 0, 88, 0, 33, 0, 0, 0, 0,
	0, 89, 77, 78, 79, 80, 81, 82, -2, -2,
	-2, -2, 60, 0, 69, 0, 32, -2, 0, 87,
	0, 47, 0, 53, 55, 68, 0, 0, 0, 32,
	0, 0, 0, 0, 50, 52, 0, 5, 15, 6,
	7, 8, 9, 10, 11, 12, 13, 14, 0, 48,
	0, 0, 0, 0, 0, 66, 0, 29, 0, 31,
	45, 0, 0, 4, 16, 0, 34, 54, 0, 0,
	0, 65, 30, 46, 51, 0, 0, 67, 64, 0,
	0, 63, 0, 0, 0, 18, 27, 21, 22, 23,
	24, 25, 26, 0, 17, 28, 0, 20, 19,
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
		//line parser/parser.go.y:57
		{
			ParseResult = yyDollar[1].Exprs
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:62
		{
			yyVAL.Exprs = []ast.Expr{yyDollar[1].Expr}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:66
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:73
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, yyDollar[4].Exprs}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:77
		{
			yyVAL.Expr = &ast.Class{yyDollar[2].String, nil}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:86
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:90
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:95
		{
			yyVAL.Expr = &ast.ObjectMethodDefineExpr{yyDollar[2].String, nil, yyDollar[6].Exprs}
		}
	case 18:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:99
		{
			yyVAL.Expr = &ast.ObjectMethodDefineExpr{yyDollar[2].String, nil, nil}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser/parser.go.y:103
		{
			yyVAL.Expr = &ast.ObjectMethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[7].Exprs}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:107
		{
			yyVAL.Expr = &ast.ObjectMethodDefineExpr{yyDollar[2].String, yyDollar[4].Strings, nil}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:115
		{
			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:119
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[2].Expr)
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:125
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, nil}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:129
		{
			yyVAL.Expr = &ast.MethodCalledExpr{yyDollar[1].String, yyDollar[3].String, yyDollar[5].Exprs}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:133
		{
			yyVAL.Expr = &ast.AttributeSetExpr{yyDollar[1].String, yyDollar[3].String, yyDollar[5].Expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:138
		{
			yyVAL.Expr = &ast.AttributeExpr{yyDollar[1].String, yyDollar[3].String}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:160
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, nil, yyDollar[5].Expr}
		}
	case 46:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:164
		{
			yyVAL.Expr = &ast.FuncDefineExpr{yyDollar[2].String, yyDollar[4].Strings, yyDollar[6].Expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:169
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, nil}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:173
		{
			yyVAL.Expr = &ast.FuncCallExpr{yyDollar[1].String, yyDollar[3].Exprs}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:180
		{
			yyVAL.Strings = append(make([]string, 0), yyDollar[1].String)
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:184
		{
			yyVAL.Strings = append(yyDollar[1].Strings, yyDollar[3].String)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:191
		{

			yyVAL.Exprs = append(make([]ast.Expr, 0), yyDollar[1].Expr)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:196
		{
			yyVAL.Exprs = append(yyDollar[1].Exprs, yyDollar[3].Expr)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:204
		{
			yyVAL.Expr = &ast.BreakExpr{}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:208
		{
			yyVAL.Expr = &ast.ReturnExpr{yyDollar[2].Expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:216
		{
			yyVAL.Expr = &ast.AssignExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:221
		{
			yyVAL.Expr = &ast.SubfixDoubleAddExpr{yyDollar[1].Expr}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:225
		{
			yyVAL.Expr = &ast.SubfixDoubleSubExpr{yyDollar[1].Expr}
		}
	case 63:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser/parser.go.y:231
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, yyDollar[4].Expr, yyDollar[6].Expr, yyDollar[7].Expr}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:235
		{
			yyVAL.Expr = &ast.ForExpr{nil, yyDollar[3].Expr, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:239
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, yyDollar[4].Expr, yyDollar[5].Expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:243
		{
			yyVAL.Expr = &ast.ForExpr{nil, nil, nil, yyDollar[4].Expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:247
		{
			yyVAL.Expr = &ast.ForExpr{yyDollar[2].Expr, nil, yyDollar[5].Expr, yyDollar[6].Expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:252
		{
			yyVAL.Expr = &ast.IFExpr{yyDollar[2].Expr, yyDollar[3].Expr, nil}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:256
		{
			(yyDollar[1].Expr.(*ast.IFExpr)).Expr2 = yyDollar[3].Expr

			yyVAL.Expr = yyDollar[1].Expr
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:266
		{
			yyVAL.Expr = &ast.ANDExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:270
		{
			yyVAL.Expr = &ast.ORExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:274
		{
			yyVAL.Expr = &ast.AddExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:278
		{
			yyVAL.Expr = &ast.SubExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:282
		{
			yyVAL.Expr = &ast.MultiExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:286
		{
			yyVAL.Expr = &ast.DivExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:290
		{
			yyVAL.Expr = &ast.GreaterThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:294
		{
			yyVAL.Expr = &ast.GreaterEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:298
		{
			yyVAL.Expr = &ast.LessThanExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:302
		{
			yyVAL.Expr = &ast.LessEqualExpr{yyDollar[1].Expr, yyDollar[3].Expr}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:307
		{
			yyVAL.Expr = &ast.BlockExpr{yyDollar[2].Exprs}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:311
		{
			yyVAL.Expr = &ast.BlockExpr{}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:316
		{
			yyVAL.Expr = &ast.Variable{yyDollar[1].String}
		}
	}
	goto yystack /* stack new state and value */
}
