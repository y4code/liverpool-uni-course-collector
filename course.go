package main

import "github.com/gocolly/colly/v2"

const CourseNamePostgraduateTaught = "Postgraduate Taught"

type Course struct {
	Name        string
	Departments []*Department
	Collector
}

// Collect save the departments information and execute the Collect method of each department.
func (c *Course) Collect() {
	var (
		collector = NewCollector()
		course    = &Course{
			Name:        CourseNamePostgraduateTaught,
			Departments: []*Department{},
		}
		err error
	)
	collector.OnHTML("#departments", func(e1 *colly.HTMLElement) {
		e1.ForEach("#courseslist > tbody > tr", func(i int, e2 *colly.HTMLElement) {
			department := &Department{
				Name:     e2.ChildText("td:nth-child(1) > a"),
				Link:     e2.ChildAttr("td:nth-child(1) > a", "href"),
				Programs: nil,
			}
			course.Departments = append(course.Departments, department)
		})
	})
	err = collector.Visit(EndpointCourse)
	if err != nil {
		panic(err)
	}
	for _, department := range course.Departments {
		department.Collect()
	}
}
