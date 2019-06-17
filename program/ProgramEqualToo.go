package program

func (p *Program) ProgramEqualToo(a, b interface{}, reg int) func() {
	return func() {
		p.Registers[reg] = (a == b)
	}
}
