package main

import (
	"log"
	"os"

	"github.com/readytotouch/gocompanies/internal/db"

	tv1 "github.com/readytotouch/gocompanies/internal/templates/v1"
)

func main() {
	f, err := os.Create("./netlify-public/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tv1.WriteCompanies(f, db.Companies(), db.Universities())

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
