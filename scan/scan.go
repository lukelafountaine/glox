package scan

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"../token"
)

const eof = -1

// Scanner struct data
type Scanner struct {
	source  string
	Tokens  []token.Token
	start   int
	current int
	width   int
	line    int
}

// NewScanner constructor
func NewScanner(source string) *Scanner {
	return &Scanner{
		source,
		make([]token.Token, 0),
		0,
		0,
		0,
		1,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, token.Token{Type: token.EOF, Lexeme: "", Literal: nil, Line: s.line})

	for _, t := range s.Tokens {
		fmt.Println(t)
	}

	return s.Tokens
}

func (s *Scanner) scanToken() {

	r := s.advance()

	switch {
	case r == '(':
		s.addToken(token.LeftParen, nil)
	case r == ')':
		s.addToken(token.RightParen, nil)
	case r == '{':
		s.addToken(token.LeftBrace, nil)
	case r == '}':
		s.addToken(token.RightBrace, nil)
	case r == ',':
		s.addToken(token.Comma, nil)
	case r == '.':
		s.addToken(token.Dot, nil)
	case r == '-':
		s.addToken(token.Minus, nil)
	case r == '+':
		s.addToken(token.Plus, nil)
	case r == ';':
		s.addToken(token.Semicolon, nil)
	case r == '*':
		s.addToken(token.Star, nil)
	case r == '!':
		if s.match('=') {
			s.addToken(token.BangEqual, nil)
		} else {
			s.addToken(token.Bang, nil)
		}
	case r == '=':
		if s.match('=') {
			s.addToken(token.EqualEqual, nil)
		} else {
			s.addToken(token.Equal, nil)
		}
	case r == '<':
		if s.match('=') {
			s.addToken(token.LessEqual, nil)
		} else {
			s.addToken(token.Less, nil)
		}
	case r == '>':
		if s.match('=') {
			s.addToken(token.GreaterEqual, nil)
		} else {
			s.addToken(token.Greater, nil)
		}
	case r == '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.Slash, nil)
		}
	case r == '"':
		s.string()
	case r == '\n':
		s.line++
	case unicode.IsSpace(r):
		// do nothing
	default:
		if unicode.IsDigit(r) {
			s.number()
		} else if unicode.IsLetter(r) {
			s.identifier()
		} else {
			fmt.Println("[ line ", s.line, "] Unexpected character:", string(r))
		}
	}
}

func (s *Scanner) number() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
		s.advance()

		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	literal, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addToken(token.Number, literal)
}

func (s *Scanner) identifier() {
	for unicode.IsLetter(s.peek()) || unicode.IsNumber(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]

	keywordType, isKeyword := token.Keywords[text]

	if isKeyword {
		s.addToken(keywordType, text)
		return
	}

	s.addToken(token.Identifier, nil)
}

func (s *Scanner) string() {

	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		fmt.Println("[ line ", s.line, "] Unterminated string")
	}

	s.advance()
	literal := strings.Trim(s.source[s.start:s.current], "\"")
	s.addToken(token.String, literal)
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return eof
	}

	r, _ := utf8.DecodeRuneInString(s.source[s.current:])
	return r
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return eof
	}

	r, _ := utf8.DecodeRuneInString(s.source[s.current+s.width:])
	return r
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	r, w := utf8.DecodeRuneInString(s.source[s.current:])

	if r != expected {
		return false
	}

	s.width = w
	s.current += s.width

	return true
}

func (s *Scanner) advance() rune {
	r, w := utf8.DecodeRuneInString(s.source[s.current:])
	s.width = w
	s.current += s.width
	return r
}

func (s *Scanner) addToken(tokenType token.Type, literal interface{}) {
	s.Tokens = append(s.Tokens, token.Token{Type: tokenType, Lexeme: s.source[s.start:s.current], Literal: literal, Line: s.line})
}
