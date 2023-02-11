# Liverpool University Course Collector

This project is designed to download all postgraduate taught master's programs offered by the University of Liverpool and save the information as CSV files for convenient searching and selection by students.

## Table of Contents
- [Liverpool University Course Collector](#liverpool-university-course-collector)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [How to Run](#how-to-run)
  - [University of Liverpool Postgraduate Taught Master's Programs](#university-of-liverpool-postgraduate-taught-masters-programs)
  - [Artifacts](#artifacts)
  - [Non-Technical People](#non-technical-people)
  - [Contribution](#contribution)

## Overview
The Liverpool University Course Collector program uses a web scraper to collect information on all postgraduate taught master's programs offered by the University of Liverpool. The information collected includes the course name, description, and any other relevant details. The program saves the information as CSV files, making it easier for students to search and compare different courses.

## How to Run
To run the Liverpool University Course Collector program, you will need to have a recent version of [Go](https://golang.org/) installed on your computer. Once you have Go installed, you can run the program by cloning the repository and using the following command in your terminal:

    go run ./

The program will create a directory named `artifacts` in the current directory and save the CSV files of the course information in that directory.

## University of Liverpool Postgraduate Taught Master's Programs
The master's program index page of the University of Liverpool can be found at the following URL: [University of Liverpool Postgraduate Taught Master's Programs](https://www.liverpool.ac.uk/study/postgraduate-taught/courses/). This page lists all of the postgraduate taught master's programs offered by the university.

## Artifacts
The artifacts of the Liverpool University Course Collector program can be found in the `./artifacts` directory. This directory contains the CSV files with the information collected on the postgraduate taught master's programs offered by the University of Liverpool.

## Non-Technical People
For non-technical people, the information collected by the Liverpool University Course Collector program is available by clicking on the [Artifacts](https://github.com/y4code/liverpool-uni-course-collector/tree/main/artifacts) link. The information is presented in the form of CSV files, which can be easily viewed in the Github page. A useful tool for non-technical people is the Github CSV online explorer, which allows users to search for specific courses and quickly find the information they need.

## Contribution
If you would like to contribute to the Liverpool University Course Collector project, please fork the repository and submit a pull request with your changes. Your contributions are greatly appreciated!
