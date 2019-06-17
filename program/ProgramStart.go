package program

// Default start
func ProgramStart(p *Program) ProgramFn {
	return p.Functions[0]()
}
