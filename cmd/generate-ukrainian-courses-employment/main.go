package main

import (
	"log"
	"os"
	"sort"

	"github.com/readytotouch/golang-companies-organizer/internal/db"

	tv1 "github.com/readytotouch/golang-companies-organizer/internal/templates/v1"
)

func main() {
	f, err := os.Create("./netlify-public/ukrainian-courses-employment.html")
	if err != nil {
		log.Fatal(err)
	}

	courses := db.Courses()
	sort.Slice(courses, func(i, j int) bool {
		return courses[i].AlumniCount > courses[j].AlumniCount
	})

	tv1.WriteUkrainianCoursesEmployment(f, courses, db.DouTop50Companies(), db.FaangCompanyGroup())

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
