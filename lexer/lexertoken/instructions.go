package lexertoken

// InstructionType describes a YOLOL instruction
type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	// COMMANDS
	// TOKEN_If IF
	TOKEN_If
	// TOKEN_Then THEN
	TOKEN_Then
	// TOKEN_End END
	TOKEN_End

	TOKEN_Key
	TOKEN_Value

	// TOKEN_Goto GOTO
	TOKEN_Goto

	// OPERATORS

	// TOKEN_Add +
	TOKEN_Add
	// TOKEN_Subtract -
	TOKEN_Subtract
	// TOKEN_Multiply *
	TOKEN_Multiply
	// TOKEN_Divide /
	TOKEN_Divide
	// TOKEN_PostIncrement A++
	TOKEN_PostIncrement
	// TOKEN_PostDecrement A--
	TOKEN_PostDecrement
	// TOKEN_PreIncrement ++A
	TOKEN_PreIncrement
	// TOKEN_PreDecrement --A
	TOKEN_PreDecrement
	// TOKEN_Assignment =
	TOKEN_Assignment
	// TOKEN_AdditionAssignment +=
	TOKEN_AdditionAssignment
	// TOKEN_SubtractionAssignment -=
	TOKEN_SubtractionAssignment
	// TOKEN_MultiplicationAssignment *=
	TOKEN_MultiplicationAssignment
	// TOKEN_DivisionAssignment /=
	TOKEN_DivisionAssignment
	// TOKEN_ModuloAssignment %=
	TOKEN_ModuloAssignment
	// TOKEN_Exponation A ^ B
	TOKEN_Exponation
	// TOKEN_Modulo A % B
	TOKEN_Modulo
	// TOKEN_Modulus ABS A
	TOKEN_Modulus
	// TOKEN_Factorial !
	TOKEN_Factorial
	// TOKEN_SquareRoot SQRT A
	TOKEN_SquareRoot
	// TOKEN_Sin SIN A
	TOKEN_Sin
	// TOKEN_Cos COS A
	TOKEN_Cos
	// TOKEN_Tan TAN A
	TOKEN_Tan
	// TOKEN_ArcSin ARCSIN A
	TOKEN_ArcSin
	// TOKEN_ArcCos ARCCOS A
	TOKEN_ArcCos
	// TOKEN_ArcTan ARCTAN A
	TOKEN_ArcTan
	// TOKEN_LessThan A < B
	TOKEN_LessThan
	// TOKEN_GreaterThan A > B
	TOKEN_GreaterThan
	// TOKEN_LessThanOrEqualToo A <= B
	TOKEN_LessThanOrEqualToo
	// TOKEN_GreaterThanOrEqualToo A >= B
	TOKEN_GreaterThanOrEqualToo
	// TOKEN_NotEqualToo A ~= B
	TOKEN_NotEqualToo
	// TOKEN_EqualToo A == B
	TOKEN_EqualToo
)
