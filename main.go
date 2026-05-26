// Program ka entry point. CLI flags (-p, -q) parse karta hai, timing measures karta hai, utils ke functions call karta hai, results print karta hai.

package main

import (
	"flag"
	"log"
	utils "serch-engine/utils"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "ag_news_test.csv", "Miami")

	flag.StringVar(&query, "q", "chemistry researcher at the University of Louisville", "search query")
	flag.Parse()

	log.Println("full text search in progress")

	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
		// %d Integer ,  \t tab space , %s string
	}
}
