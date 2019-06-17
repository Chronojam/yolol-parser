package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexPostDecrement(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.PostDecrement)
	lexer.Emit(lexertoken.TOKEN_PostDecrement)
	return LexBegin
}
