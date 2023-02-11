package main

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

type Department struct {
	Name     string
	Link     string
	Programs []*Program
	Collector
}

var (
	suffixes = []string{"(Online)", "Robotics, Automation and Artificial Intelligence MSc"}
)

// Collect save the programs information and execute the Collect method of each program.
func (d *Department) Collect() error {
	var (
		collector = NewCollector()
		err       error
	)
	collector.OnHTML(".department-list", func(e1 *colly.HTMLElement) {
		e1.ForEach("li", func(i int, e2 *colly.HTMLElement) {
			link := e2.ChildAttr("h2 > a", "href")
			if strings.HasPrefix(link, "/study") {
				link = "https://www.liverpool.ac.uk" + link
			}
			program := &Program{
				Name:     e2.ChildText("h2 > a"),
				Link:     link,
				Overview: &Overview{},
			}
			if hasSuffix(program.Name, suffixes) {
				return
			}
			d.Programs = append(d.Programs, program)
		})
	})
	err = collector.Visit(d.Link)
	if err != nil {
		return err
	}
	for _, program := range d.Programs {
		err = program.Collect()
		if err != nil {
			return err
		}
	}
	return nil
}

func hasSuffix(name string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(name, suffix) {
			return true
		}
	}
	return false
}
