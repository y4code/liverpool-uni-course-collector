package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

type Overview struct {
	ModulesLink string
	Modules     []*Module
	Collector
}

func (o *Overview) Collect() error {
	var (
		collector = NewCollector()
	)
	collector.OnHTML("#module-details", func(e1 *colly.HTMLElement) {
		e1.ForEach("h2", func(i int, e2 *colly.HTMLElement) {
			e2.DOM.Next().Children().Each(func(i int, selection *goquery.Selection) {
				if selection.Is("h5") {
					table := selection.Next()
					aims, _ := table.Find("tr").Eq(IndexAims).Find("td").Html()
					learningOutcomes, _ := table.Find("tr").Eq(IndexLearningOutcomes).Find("td").Html()
					module := &Module{
						Name:                    selection.Text()[:strings.LastIndex(selection.Text(), "(")],
						Code:                    selection.Text()[strings.LastIndex(selection.Text(), "("):],
						IsCompulsory:            e2.Text == CompulsoryModules,
						Level:                   table.Find("tr").Eq(IndexLevel).Find("td").Text(),
						CreditLevel:             table.Find("tr").Eq(IndexCreditLevel).Find("td").Text(),
						Semester:                table.Find("tr").Eq(IndexSemester).Find("td").Text(),
						ExamCourseworkWeighting: table.Find("tr").Eq(IndexExamCourseworkWeighting).Find("td").Text(),
						Aims:                    replaceHTMLWithLiteralFormat.Replace(aims),
						LearningOutcomes:        replaceHTMLWithLiteralFormat.Replace(learningOutcomes),
					}
					o.Modules = append(o.Modules, module)
				}
			})
		})
	})
	err := collector.Visit(o.ModulesLink)
	if err != nil {
		log.Println(err, " for ", o.ModulesLink)
		return nil
	}
	return nil
}

var (
	replaceHTMLWithLiteralFormat = strings.NewReplacer(
		"<p>", "",
		"</p>", "\n",
		"<br />", "\n",
		"<br>", "\n",
		"<br/>", "\n",
		" 2. ", "\n2. ",
		" 3. ", "\n3. ",
		" 4. ", "\n4. ",
		" 5. ", "\n5. ",
		" 6. ", "\n6. ",
		" 7. ", "\n7. ",
		" 8. ", "\n8. ",
		" 9. ", "\n9. ",
		" 10. ", "\n10. ",
		" 11. ", "\n11. ",
		" 12. ", "\n12. ",
		" 13. ", "\n13. ",
		" 14. ", "\n14. ",
		" 15. ", "\n15. ",
	)
)
