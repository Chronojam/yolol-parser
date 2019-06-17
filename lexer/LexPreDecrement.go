package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexPreDecrement(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.PreDecrement)
	lexer.Emit(lexertoken.TOKEN_PreDecrement)
	return LexBegin
}
