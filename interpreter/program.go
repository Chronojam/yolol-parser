package program

type Program struct {
	State    map[string]interface{}
	Executor []func()
}

func (p *Program) Run() {

}
