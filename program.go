package main

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

const (
	IndexName = iota - 3
	IndexModuleCode
	IndexCompulsoryOrOptional
	IndexLevel
	IndexCreditLevel
	IndexSemester
	IndexExamCourseworkWeighting
	IndexAims
	IndexLearningOutcomes
)

const (
	overview      = "/overview"
	moduleDetails = "/module-details"
)

type Program struct {
	Name string
	Link string
	*Overview
	Collector
}

// Collect save the programs information and execute the Collect method of Overview.
func (p *Program) Collect() error {
	var (
		collector = NewCollector()
		err       error
	)
	collector.OnHTML("#course-tabs", func(e1 *colly.HTMLElement) {
		p.Overview = &Overview{
			ModulesLink: strings.ReplaceAll(p.Link, overview, moduleDetails),
			Modules:     nil,
		}
	})
	err = collector.Visit(p.Link)
	if err != nil {
		return err
	}
	if p.Overview == nil {
		return nil
	}
	err = p.Overview.Collect()
	if err != nil {
		return err
	}
	return nil
}
