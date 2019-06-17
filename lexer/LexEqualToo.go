package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexEqualToo(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.EqualToo)
	lexer.Emit(lexertoken.TOKEN_EqualToo)
	return LexValue
}
