package token

import "fmt"

type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) String() string {
	return fmt.Sprint(typeNames[t.Type], " ", t.Lexeme, " ", t.Literal, " ", t.Line)
}
