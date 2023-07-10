package agent

type Agent struct {
	address string
}

func New(address string) *Agent {
	return &Agent{
		address: address,
	}
}

func (a *Agent) Run() {

}
