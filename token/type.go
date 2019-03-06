package token

type Type int

// Token Types
const (
	EOF Type = iota
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Reserved Words
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
	Error
)

var typeNames = map[Type]string{
	EOF:          "EOF",
	LeftParen:    "LeftParen",
	RightParen:   "RightParen",
	LeftBrace:    "LeftBrace",
	RightBrace:   "RightBrace",
	Comma:        "Comma",
	Dot:          "Dot",
	Minus:        "Minus",
	Plus:         "Plus",
	Semicolon:    "Semicolon",
	Slash:        "Slash",
	Star:         "Star",
	Bang:         "Bang",
	BangEqual:    "BangEqual",
	Equal:        "Equal",
	EqualEqual:   "EqualEqual",
	Greater:      "Greater",
	GreaterEqual: "GreaterEqual",
	Less:         "Less",
	LessEqual:    "LessEqual",
	Identifier:   "Identifier",
	String:       "String",
	Number:       "Number",
	And:          "And",
	Class:        "Class",
	Else:         "Else",
	False:        "False",
	Fun:          "Fun",
	For:          "For",
	If:           "If",
	Nil:          "Nil",
	Or:           "Or",
	Print:        "Print",
	Return:       "Return",
	Super:        "Super",
	This:         "This",
	True:         "True",
	Var:          "Var",
	While:        "While",
}

// string to token type for keywords
var Keywords = map[string]Type{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"fun":    Fun,
	"for":    For,
	"if":     If,
	"nil":    Nil,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}
