package lexer

import (
	"strings"

	"github.com/chronojam/yolol-parser/lexer/errors"
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_KEY with the name of an
key that will be assigned a value.
*/
func LexKey(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EqualToo) {
			lexer.Emit(lexertoken.TOKEN_Key)
			return LexEqualToo
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.Assignment) {
			lexer.Emit(lexertoken.TOKEN_Key)
			return LexAssignment
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}
