package main

import (
	"fmt"

	"github.com/chronojam/yolol-parser/lexer"
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func main() {
	l := &lexer.Lexer{
		Name: "Foobar",
		Input: `foo = bar
foo == pineapples
IF foo == bar THEN foo = apples END`,
		State:  lexer.LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}

	for {
		tok := l.NextToken()
		if tok.Type == lexertoken.TOKEN_EOF {
			break
		}
		fmt.Printf("%v\n", tok)
	}
}
