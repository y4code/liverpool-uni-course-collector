package main

type Collector interface {
	Collect() error
}
