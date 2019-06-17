package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexThen(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.Then)
	lexer.Emit(lexertoken.TOKEN_Then)
	return LexBegin
}
