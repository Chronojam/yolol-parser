package program

import (
	"strings"

	"github.com/chronojam/yolol-parser/lexer"
	"github.com/chronojam/yolol-parser/lexer/lexertoken"
)

func Parse() {
	l := &lexer.Lexer{
		Name: "Parser",
		Input: `foo = bar
		foo == pineapples
		IF foo == bar THEN foo = apples END`,
		State:  lexer.LexBegin,
		Tokens: make(chan lexertoken.Token, 3),
	}

	p := Program{
		// GOTO support
		Registers: make([]interface{}, 4),
	}

	for {
		tok := l.NextToken()
		if tok.Type == lexertoken.TOKEN_EOF {
			break
		}
		startOfLine := false
		if strings.HasPrefix(tok.Value.(string), lexertoken.NewLine) {
			// We're the start of a newline.
			startOfLine = true
		}

		switch tok.Type {
		case lexertoken.TOKEN_Key:
			// Read forwards 2 operations
			// Operator token
			operatorTok := l.NextToken()
			valueTok := l.NextToken()

			keyValue := tok.Value.(string)

			if operatorTok.Type == lexertoken.TOKEN_Assignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = valueTok.Value.(string)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_AdditionAssignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = p.GlobalState[keyValue].(int) + valueTok.Value.(int)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_SubtractionAssignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = p.GlobalState[keyValue].(int) - valueTok.Value.(int)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_MultiplicationAssignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = p.GlobalState[keyValue].(int) * valueTok.Value.(int)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_DivisionAssignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = p.GlobalState[keyValue].(int) / valueTok.Value.(int)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_ModuloAssignment {
				p.Functions = append(p.Functions, func() {
					p.GlobalState[keyValue] = p.GlobalState[keyValue].(int) % valueTok.Value.(int)
				})
			}

			// Registers, Eph. State
			if operatorTok.Type == lexertoken.TOKEN_EqualToo {
				p.Functions = append(p.Functions, func() {
					p.Registers[0] = (tok.Value == valueTok.Value)
				})
			}
			if operatorTok.Type == lexertoken.TOKEN_GreaterThan {
				p.Functions = append(p.Functions, func() {
					p.Registers[0] = (tok.Value.(int) > valueTok.Value.(int))
				})
			}
			// Etc

		}
	}
}
