package program

type Program struct {
	GlobalState map[string]interface{}
	LocalState  map[string]interface{}

	// lol
	Registers []interface{}

	Functions []func()
}
