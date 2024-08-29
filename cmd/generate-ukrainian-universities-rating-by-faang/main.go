package main

import (
	"log"
	"os"

	"github.com/readytotouch/golang-companies-organizer/internal/db"

	tv1 "github.com/readytotouch/golang-companies-organizer/internal/templates/v1"
)

func main() {
	f, err := os.Create("./netlify-public/ukrainian-universities-rating-by-faang.html")
	if err != nil {
		log.Fatal(err)
	}

	tv1.WriteUkrainianUniversitiesRatingByFAANG(f, db.Universities(), db.FaangCompanyGroup())

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
