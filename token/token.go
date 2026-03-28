package lexer

import "fmt"

// TokenType はトークンの種別を表す
type TokenType int

const (
	// 特殊
	ILLEGAL TokenType = iota
	EOF

	// リテラル
	IDENT      // main, x, foo
	INT        // 42, 0x1F, 0b1010, 0o77
	FLOAT      // 3.14, 1e10
	IMAG       // 3.14i
	RUNE       // 'a', '\n'
	STRING     // "hello", `raw`

	// キーワード（Go仕様準拠）
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE
	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR
	FUNC
	GO
	GOTO
	IF
	IMPORT
	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN
	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR

	// 演算子・区切り記号
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	REM    // %
	AND    // &
	OR     // |
	XOR    // ^
	SHL    // <<
	SHR    // >>
	AND_NOT // &^

	ADD_ASSIGN    // +=
	SUB_ASSIGN    // -=
	MUL_ASSIGN    // *=
	QUO_ASSIGN    // /=
	REM_ASSIGN    // %=
	AND_ASSIGN    // &=
	OR_ASSIGN     // |=
	XOR_ASSIGN    // ^=
	SHL_ASSIGN    // <<=
	SHR_ASSIGN    // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
)

var tokenNames = map[TokenType]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	RUNE:   "RUNE",
	STRING: "STRING",

	BREAK: "break", CASE: "case", CHAN: "chan",
	CONST: "const", CONTINUE: "continue", DEFAULT: "default",
	DEFER: "defer", ELSE: "else", FALLTHROUGH: "fallthrough",
	FOR: "for", FUNC: "func", GO: "go",
	GOTO: "goto", IF: "if", IMPORT: "import",
	INTERFACE: "interface", MAP: "map", PACKAGE: "package",
	RANGE: "range", RETURN: "return", SELECT: "select",
	STRUCT: "struct", SWITCH: "switch", TYPE: "type",
	VAR: "var",

	ADD: "+", SUB: "-", MUL: "*", QUO: "/", REM: "%",
	AND: "&", OR: "|", XOR: "^", SHL: "<<", SHR: ">>", AND_NOT: "&^",

	ADD_ASSIGN: "+=", SUB_ASSIGN: "-=", MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=", REM_ASSIGN: "%=", AND_ASSIGN: "&=",
	OR_ASSIGN: "|=", XOR_ASSIGN: "^=", SHL_ASSIGN: "<<=",
	SHR_ASSIGN: ">>=", AND_NOT_ASSIGN: "&^=",

	LAND: "&&", LOR: "||", ARROW: "<-", INC: "++", DEC: "--",
	EQL: "==", LSS: "<", GTR: ">", ASSIGN: "=", NOT: "!",
	NEQ: "!=", LEQ: "<=", GEQ: ">=", DEFINE: ":=", ELLIPSIS: "...",

	LPAREN: "(", LBRACK: "[", LBRACE: "{", COMMA: ",", PERIOD: ".",
	RPAREN: ")", RBRACK: "]", RBRACE: "}", SEMICOLON: ";", COLON: ":",
}

func (t TokenType) String() string {
	if s, ok := tokenNames[t]; ok {
		return s
	}
	return fmt.Sprintf("TokenType(%d)", int(t))
}

// keywords はGoのキーワード一覧
var keywords = map[string]TokenType{
	"break": BREAK, "case": CASE, "chan": CHAN,
	"const": CONST, "continue": CONTINUE, "default": DEFAULT,
	"defer": DEFER, "else": ELSE, "fallthrough": FALLTHROUGH,
	"for": FOR, "func": FUNC, "go": GO,
	"goto": GOTO, "if": IF, "import": IMPORT,
	"interface": INTERFACE, "map": MAP, "package": PACKAGE,
	"range": RANGE, "return": RETURN, "select": SELECT,
	"struct": STRUCT, "switch": SWITCH, "type": TYPE,
	"var": VAR,
}

// LookupIdent は識別子がキーワードかどうかを判定する
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Pos はソースコード上の位置
type Pos struct {
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

// Token はレキサーが返すトークン
type Token struct {
	Type    TokenType
	Literal string
	Pos     Pos
}

func (t Token) String() string {
	return fmt.Sprintf("Token{%s, %q, %s}", t.Type, t.Literal, t.Pos)
}
