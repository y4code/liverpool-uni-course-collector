package main

type Department struct {
	Name     string
	Link     string
	Programs []*Program
	Collector
}

func (d *Department) Collect() {
	//TODO implement me
	panic("implement me")
}
