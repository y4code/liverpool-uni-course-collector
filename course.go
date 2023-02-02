package main

const CourseNamePostgraduateTaught = "Postgraduate Taught"

type Course struct {
	Name        string
	Departments []*Department
}
