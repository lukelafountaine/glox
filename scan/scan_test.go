package scan

func ExampleScanTokens() {
	input := `and class else
	false fun for if nil
	or print return super this true var
	while
	1 2 3 1.1 0.0 0
	(){},.-+;*!====<<=>=>///comment`

	scanner := NewScanner(input)
	scanner.ScanTokens()
	// Output:
	// And and and 1
	// Class class class 1
	// Else else else 1
	// False false false 2
	// Fun fun fun 2
	// For for for 2
	// If if if 2
	// Nil nil nil 2
	// Or or or 3
	// Print print print 3
	// Return return return 3
	// Super super super 3
	// This this this 3
	// True true true 3
	// Var var var 3
	// While while while 4
	// Number 1 1 5
	// Number 2 2 5
	// Number 3 3 5
	// Number 1.1 1.1 5
	// Number 0.0 0 5
	// Number 0 0 5
	// LeftParen ( <nil> 6
	// RightParen ) <nil> 6
	// LeftBrace { <nil> 6
	// RightBrace } <nil> 6
	// Comma , <nil> 6
	// Dot . <nil> 6
	// Minus - <nil> 6
	// Plus + <nil> 6
	// Semicolon ; <nil> 6
	// Star * <nil> 6
	// BangEqual != <nil> 6
	// EqualEqual == <nil> 6
	// Equal = <nil> 6
	// Less < <nil> 6
	// LessEqual <= <nil> 6
	// GreaterEqual >= <nil> 6
	// Greater > <nil> 6
	// EOF  <nil> 6
}
