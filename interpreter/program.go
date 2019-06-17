package program

type Program struct {
	State    map[string]interface{}
	Executor []func()
}
