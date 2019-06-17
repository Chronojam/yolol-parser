package lexertoken

const (
	EOF     rune   = 0
	If      string = "IF"
	Then    string = "THEN"
	End     string = "END"
	NewLine string = "\n"

	Assignment string = "="
	EqualToo   string = "=="
)

type Token struct {
	Type  TokenType
	Value interface{}
}
