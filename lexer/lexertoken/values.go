package lexertoken

const (
	EOF          rune   = 0
	LeftBracket  string = "("
	RightBracket string = ")"

	If      string = "IF"
	Then    string = "THEN"
	End     string = "END"
	NewLine string = "\n"

	Goto string = "goto"

	Add                      string = "+"
	Subtract                 string = "-"
	Multiply                 string = "*"
	Divide                   string = "/"
	PostIncrement            string = "++"
	PostDecrement            string = "--"
	PreIncrement             string = "++"
	PreDecrement             string = "--"
	Assignment               string = "="
	AdditionAssignment       string = "+="
	SubtractionAssignment    string = "-="
	MultiplicationAssignment string = "*="
	DivisionAssignment       string = "/="
	ModuloAssignment         string = "%="
	Exponation               string = "^"
	Modulo                   string = "%"
	Modulus                  string = "ABS"
	Factorial                string = "!"
	SquareRoot               string = "SQRT"
	Sin                      string = "SIN"
	Cos                      string = "COS"
	Tan                      string = "TAN"
	ArcSin                   string = "ARCSIN"
	ArcCos                   string = "ARCCOS"
	ArcTan                   string = "ARCTAN"

	LessThan              string = "<"
	GreaterThan           string = ">"
	LessThanOrEqualToo    string = "<="
	GreaterThanOrEqualToo string = ">="
	NotEqualToo           string = "~="
	EqualToo              string = "=="
)

type Token struct {
	Type  TokenType
	Value interface{}
}
