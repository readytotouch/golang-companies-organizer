package main

import (
	"log"
	"os"

	tv1 "github.com/readytotouch/gocompanies/internal/templates/v1"
)

func main() {
	f, err := os.Create("./netlify-public/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tv1.WriteCompanies(f)

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
