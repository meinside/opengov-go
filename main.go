package main

import (
	"flag"
	"log"
)

var asSqlite, asCsv bool

func init() {
	flag.BoolVar(&asSqlite, "sqlite", false, "Save result as .sqlite")
	flag.BoolVar(&asCsv, "csv", false, "Save result as .csv")
}

func main() {
	flag.Parse()

	if asSqlite || asCsv {
		log.Println("> Starting...")

		log.Println("> Crawling info list...")
		infoList := crawlInfoList()
		log.Printf("> Total %d info list(s)\n", len(infoList))

		log.Println("> Crawling public list...")
		publicList := crawlPublicList()
		log.Printf("> Total %d public list(s)\n", len(publicList))

		log.Println("> Crawling research list...")
		researchList := crawlResearchList()
		log.Printf("> Total %d research list(s)\n", len(researchList))

		log.Println("> Saving...")

		// .sqlite
		if asSqlite {
			if err := save(TypeSqlite, infoList, publicList, researchList); err == nil {
				log.Println("> Finished!")
			} else {
				log.Printf("* Error while saving: %s\n", err)
			}
		}

		// .csv
		if asCsv {
			if err := save(TypeCsv, infoList, publicList, researchList); err == nil {
				log.Println("> Finished!")
			} else {
				log.Printf("* Error while saving: %s\n", err)
			}
		}
	} else {
		flag.PrintDefaults()
	}
}
