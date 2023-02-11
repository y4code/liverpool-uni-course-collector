package main

type Program struct {
	Name    string
	Modules []*Module
	Collector
}

func (p *Program) Collect() {
	//TODO implement me
	panic("implement me")
}
