package lexer

type Lexer struct {
	input	string
	position	string
	readPosition	int
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
