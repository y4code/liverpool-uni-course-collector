package main

import "github.com/gocolly/colly/v2"

const CourseNamePostgraduateTaught = "Postgraduate Taught"

type Course struct {
	Name        string
	Departments []*Department
	Collector
}

// Collect save the departments information and execute the Collect method of each department.
func (c *Course) Collect() error {
	var (
		collector = NewCollector()
		err       error
	)
	collector.OnHTML("#departments", func(e1 *colly.HTMLElement) {
		e1.ForEach("#courseslist > tbody > tr", func(i int, e2 *colly.HTMLElement) {
			department := &Department{
				Name:     e2.ChildText("td:nth-child(1) > a"),
				Link:     EndpointCourse + e2.ChildAttr("td:nth-child(1) > a", "href"),
				Programs: nil,
			}
			c.Departments = append(c.Departments, department)
		})
	})
	err = collector.Visit(EndpointCourse)
	if err != nil {
		return err
	}
	for _, department := range c.Departments {
		err = department.Collect()
		if err != nil {
			return err
		}
	}
	return nil
}
