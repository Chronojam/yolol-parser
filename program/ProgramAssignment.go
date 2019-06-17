package program

func (p *Program) ProgramAssignment(key string, value interface{}) func() {
	return func() {
		p.GlobalState[key] = value
	}
}
