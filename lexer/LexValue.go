package lexer

import (
	"strings"

	"github.com/chronojam/yolol-parser/lexer/errors"
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexValue(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.End) {
			lexer.Emit(lexertoken.TOKEN_Value)
			return LexEnd
		}
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.Then) {
			lexer.Emit(lexertoken.TOKEN_Value)
			return LexThen
		}
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NewLine) {
			lexer.Emit(lexertoken.TOKEN_Value)
			return LexBegin
		}

		lexer.Inc()
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}
