package main

import (
	"encoding/csv"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

const (
	CSVFileName  = "data.csv"
	ColumnLength = 9
	ColumnOffset = 3

	artifactPath = "artifacts"

	EndpointCourse = "https://www.liverpool.ac.uk/study/postgraduate-taught/courses/"
)

var (
	csvHeader    = []string{"Name", "Module Code", "Compulsory or Optional", "Level", "Credit Level", "Semester", "Exam Coursework Weighting", "Aims", "Learning Outcomes"}
	NewCollector = func() *colly.Collector {
		return colly.NewCollector(
			colly.AllowedDomains("www.liverpool.ac.uk", "liverpool.ac.uk"),
			colly.AllowURLRevisit(),
		)
	}
)

func main() {
	var (
		course = &Course{
			Name: CourseNamePostgraduateTaught,
		}
	)
	//delete all files in artifactPath
	err := os.RemoveAll(artifactPath)
	if err != nil {
		panic(err)

	}
	//collect
	err = course.Collect()
	if err != nil {
		panic(err)
	}

	//csv
	for _, department := range course.Departments {
		for _, program := range department.Programs {
			err = saveToCSV(department, program)
			if err != nil {
				panic(err)
			}
		}
	}
}

func saveToCSV(department *Department, program *Program) error {
	// create directory if not exists
	var (
		dataPath = artifactPath + "/" + department.Name
	)
	err := os.MkdirAll(dataPath, os.ModePerm)
	if err != nil {
		return err
	}
	// create file if not exists
	fileName := strings.ReplaceAll(program.Name, "/", "|")
	file, err := os.Create(dataPath + "/" + fileName + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()
	// prepare csv stuff
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// write csv header
	err = writer.Write(csvHeader)
	if err != nil {
		return err
	}
	// write csv content
	for _, module := range program.Overview.Modules {
		err = writer.Write([]string{
			module.Name,
			module.Code,
			Compulsory(module.IsCompulsory),
			module.Level,
			module.CreditLevel,
			module.Semester,
			module.ExamCourseworkWeighting,
			module.Aims,
			module.LearningOutcomes,
		})
		if err != nil {
			return err
		}
	}
	// do something to finish this
	return nil
}
