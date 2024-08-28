package main

import (
	"log"
	"os"

	"github.com/readytotouch/golang-companies-organizer/internal/db"

	tv1 "github.com/readytotouch/golang-companies-organizer/internal/templates/v1"
)

func main() {
	v00_00_00()
	v00_00_01()
}

func v00_00_00() {
	f := mustOpen("./netlify-public/v00.00.00.html")

	tv1.WriteCompanies(f, db.GoCompanies(), db.Universities())

	mustClose(f)
}

func v00_00_01() {
	{
		f := mustOpen("./netlify-public/v00.00.01.html")

		tv1.WriteOrganizer(f, db.GoCompanies(), db.Universities())

		mustClose(f)
	}

	// for Worldwide for Reddit article
	{
		f := mustOpen("./netlify-public/index.html")

		tv1.WriteOrganizer(f, db.GoCompanies(), nil)

		mustClose(f)
	}

	// for Ukraine for DOU article
	{
		f := mustOpen("./netlify-public/ukraine.html")

		tv1.WriteOrganizer(f, db.GoCompanies(), db.Universities())

		mustClose(f)
	}
}

func mustOpen(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func mustClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
