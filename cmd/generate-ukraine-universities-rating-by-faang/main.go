package main

import (
	"log"
	"os"

	"github.com/readytotouch/gocompanies/internal/db"

	tv1 "github.com/readytotouch/gocompanies/internal/templates/v1"
)

func main() {
	f, err := os.Create("./netlify-public/ukrainian-universities-rating-by-faang.html")
	if err != nil {
		log.Fatal(err)
	}

	tv1.WriteUkraineUniversitiesRatingByFAANG(f, db.Universities())

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
