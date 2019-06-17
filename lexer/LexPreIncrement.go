package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexPreIncrement(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.PreIncrement)
	lexer.Emit(lexertoken.TOKEN_PreIncrement)
	return LexBegin
}
