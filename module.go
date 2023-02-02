package main

type Module struct {
	Name                    string
	Code                    string // eg: (COMP516)
	IsCompulsory            bool   // true(Compulsory modules), false(Optional modules)
	Level                   string // eg: M, 3
	CreditLevel             int    // eg: 15
	Semester                string // eg: First Semester, Second Semester
	ExamCourseworkWeighting string // eg: 75:25(Exam:Coursework)
	Aims                    string
	LearningOutcomes        string
}
