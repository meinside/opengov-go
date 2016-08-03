package main

import (
	"log"
)

func main() {
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
	if err := save(TypeSqlite, infoList, publicList, researchList); err == nil {
		log.Println("> Finished!")
	} else {
		log.Printf("* Error while saving: %s\n", err)
	}
}
