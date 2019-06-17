package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_If then returns
the lexer for a conditional statement
*/
func LexIf(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.If)
	lexer.Emit(lexertoken.TOKEN_If)
	return LexKey
}
