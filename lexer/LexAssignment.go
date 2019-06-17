package lexer

import (
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexAssignment(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.Assignment)
	lexer.Emit(lexertoken.TOKEN_Assignment)
	return LexValue
}
