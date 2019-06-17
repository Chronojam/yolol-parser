package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexMultiply(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.Multiply)
	lexer.Emit(lexertoken.TOKEN_Multiply)
	return LexValue
}
