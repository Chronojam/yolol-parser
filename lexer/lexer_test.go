package lexer

import (
	"fmt"
	"testing"

	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func TestLexer(t *testing.T) {
	tok := make(chan lexertoken.Token)
	l := &Lexer{
		Name:   "Foobar",
		Input:  "IF foo == bar THEN foo == apples END",
		State:  LexBegin,
		Tokens: tok,
	}

	l.Run()

	for {
		token := <-tok
		fmt.Printf("%v", token)
	}
}
