package lexer

import (
	"strings"

	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	remainingInput := lexer.InputToEnd()

	// Try and determine what instruction is next.
	// This is gross, but yolo.
	if strings.HasPrefix(remainingInput, lexertoken.If) {
		return LexIf
	}
	if strings.HasPrefix(remainingInput, lexertoken.End) {
		return LexEnd
	}

	return LexKey
}
