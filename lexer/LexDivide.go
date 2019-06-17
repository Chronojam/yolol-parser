package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexDivide(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.Divide)
	lexer.Emit(lexertoken.TOKEN_Divide)
	return LexValue
}
