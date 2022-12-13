package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
	"strings"
)

const (
	CSVFileName  = "data.csv"
	ColumnLength = 9
	ColumnOffset = 3
)

const (
	NameIndex = iota
	ModuleCodeIndex
	CompulsoryOrOptionalIndex
	LevelIndex
	CreditLevelIndex
	SemesterIndex
	ExamCourseworkWeightingIndex
	AimsIndex
	LearningOutcomesIndex
)

func main() {
	var (
		file       *os.File
		err        error
		csvContent [][]string
	)
	//delete file if exists
	if _, err = os.Stat(CSVFileName); err == nil {
		if err = os.Remove(CSVFileName); err != nil {
			panic(err)
		}
	}
	//create file if not exists
	if file, err = os.OpenFile(CSVFileName, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()
	if csvContent, err = csv.NewReader(file).ReadAll(); err != nil {
		log.Fatal("Cannot read file", err)
	}
	csvContent = append(csvContent, []string{"Name", "Module Code", "Compulsory or Optional", "Level", "Credit Level", "Semester", "Exam Coursework Weighting", "Aims", "Learning Outcomes"})
	c := colly.NewCollector(
		colly.AllowedDomains("liverpool.ac.uk", "www.liverpool.ac.uk"),
	)
	c.OnHTML("section[id=module-details]", func(e *colly.HTMLElement) {
		e.ForEach("h2", func(i int, el *colly.HTMLElement) {
			el.DOM.Next().Children().Each(func(i int, selection *goquery.Selection) {
				var row = make([]string, ColumnLength)
				row[CompulsoryOrOptionalIndex] = el.Text
				if selection.Is("h5") {
					row[ModuleCodeIndex] = selection.Text()[strings.LastIndex(selection.Text(), "("):]
					row[NameIndex] = selection.Text()[:strings.LastIndex(selection.Text(), "(")]
					table := selection.Next()
					row[LevelIndex] = table.Find("tr").Eq(LevelIndex - ColumnOffset).Find("td").Text()
					row[CreditLevelIndex] = table.Find("tr").Eq(CreditLevelIndex - ColumnOffset).Find("td").Text()
					row[SemesterIndex] = table.Find("tr").Eq(SemesterIndex - ColumnOffset).Find("td").Text()
					row[ExamCourseworkWeightingIndex] = table.Find("tr").Eq(ExamCourseworkWeightingIndex - ColumnOffset).Find("td").Text()
					aims, _ := table.Find("tr").Eq(AimsIndex - ColumnOffset).Find("td").Html()
					row[AimsIndex] = replaceHTMLWithLiteralFormat.Replace(aims)
					learningOutcomes, _ := table.Find("tr").Eq(LearningOutcomesIndex - ColumnOffset).Find("td").Html()
					row[LearningOutcomesIndex] = replaceHTMLWithLiteralFormat.Replace(learningOutcomes)
					csvContent = append(csvContent, row)
				}
			})
		})
	})
	err = c.Visit("https://www.liverpool.ac.uk/study/postgraduate-taught/taught/advanced-computer-science-msc/module-details/")
	if err != nil {
		panic(err)
	}
	if err = csv.NewWriter(file).WriteAll(csvContent); err != nil {
		log.Fatalln("Cannot write to file", err)
	}
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
