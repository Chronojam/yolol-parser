package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexEnd(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.End)
	lexer.Emit(lexertoken.TOKEN_End)
	return LexBegin
}
