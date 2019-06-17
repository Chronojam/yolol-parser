package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexPostIncrement(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.PostIncrement)
	lexer.Emit(lexertoken.TOKEN_PostIncrement)
	return LexBegin
}
