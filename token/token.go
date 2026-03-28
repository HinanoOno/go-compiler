package token

type TokenType int

const (
	// 特殊トークン
	ILLEGAL TokenType = iota
	EOF
	COMMENT

	// 識別子・リテラル
	IDENT  // add, x, y...
	INT    // 123
	FLOAT  // 12.3
	STRING // "hello"

	// 演算子
	ADD            // +
	SUB            // -
	MUL            // *
	DIV            // /
	MOD            // %
	AND            // &
	OR             // |
	XOR            // ^
	SHL            // <<
	SHR            // >>
	AND_NOT        // &^
	ADD_ASSIGN     // +=
	SUB_ASSIGN     // -=
	MUL_ASSIGN     // *=
	DIV_ASSIGN     // /=
	MOD_ASSIGN     // %=
	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=
	LAND           // &&
	LOR            // ||
	ARROW          // <-
	INC            // ++
	DEC            // --
	EQL            // ==
	LSS            // <
	GTR            // >
	ASSIGN         // =
	NOT            // !
	NEQ            // !=
	LEQ            // <=
	GEQ            // >=
	DEFINE         // :=

	// 区切り記号
	LPAREN    // (
	LBRACK    // [
	LBRACE    // {
	COMMA     // ,
	PERIOD    // .
	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	// キーワード
	keyword_beg
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
	keyword_end
)

type Pos struct {
	Line   int
	Column int
}

// トークンのデータ構造
type Token struct {
	Type    TokenType
	Literal string
	Pos     Pos
}

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	STRING: "STRING",

	ADD:            "+",
	SUB:            "-",
	MUL:            "*",
	DIV:            "/",
	MOD:            "%",
	AND:            "&",
	OR:             "|",
	XOR:            "^",
	SHL:            "<<",
	SHR:            ">>",
	AND_NOT:        "&^",
	ADD_ASSIGN:     "+=",
	SUB_ASSIGN:     "-=",
	MUL_ASSIGN:     "*=",
	DIV_ASSIGN:     "/=",
	MOD_ASSIGN:     "%=",
	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",
	LAND:           "&&",
	LOR:            "||",
	ARROW:          "<-",
	INC:            "++",
	DEC:            "--",
	EQL:            "==",
	LSS:            "<",
	GTR:            ">",
	ASSIGN:         "=",
	NOT:            "!",
	NEQ:            "!=",
	LEQ:            "<=",
	GEQ:            ">=",
	DEFINE:         ":=",

	LPAREN:    "(",
	LBRACK:    "[",
	LBRACE:    "{",
	COMMA:     ",",
	PERIOD:    ".",
	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:       "break",
	CASE:        "case",
	CHAN:        "chan",
	CONST:       "const",
	CONTINUE:    "continue",
	DEFAULT:     "default",
	DEFER:       "defer",
	ELSE:        "else",
	FALLTHROUGH: "fallthrough",
	FOR:         "for",
	FUNC:        "func",
	GO:          "go",
	GOTO:        "goto",
	IF:          "if",
	IMPORT:      "import",
	INTERFACE:   "interface",
	MAP:         "map",
	PACKAGE:     "package",
	RANGE:       "range",
	RETURN:      "return",
	SELECT:      "select",
	STRUCT:      "struct",
	SWITCH:      "switch",
	TYPE:        "type",
	VAR:         "var",
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = TokenType(i)
	}
}

// String はデバッグ・エラーメッセージ用の表示名を返す
func (t TokenType) String() string {
	if int(t) < len(tokens) {
		if s := tokens[t]; s != "" {
			return s
		}
	}
	return "UNKNOWN"
}

// 識別子がキーワードかユーザー定義の名前かを判定する
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// TriggersSemicolon はGo仕様の自動セミコロン挿入対象トークンかを返す
func (t TokenType) TriggersSemicolon() bool {
	switch t {
	case IDENT, INT, FLOAT, STRING,
		BREAK, CONTINUE, FALLTHROUGH, RETURN,
		INC, DEC,
		RPAREN, RBRACK, RBRACE:
		return true
	}
	return false
}
