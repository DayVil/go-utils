package main

import (
	"log"

	"github.com/DayVil/go-utils/indexer"
)

func main() {
	ind, _ := indexer.IndexFiles("/mnt/f/")

	log.Printf("Files: %d\n", len(ind.Files))
}
