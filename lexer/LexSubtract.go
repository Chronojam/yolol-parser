package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexSubtract(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.Subtract)
	lexer.Emit(lexertoken.TOKEN_Subtract)
	return LexValue
}
