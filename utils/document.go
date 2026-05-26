package utils

import (
	"encoding/csv"
	"os"
)

// Document represents searchable content
type Document struct {
	ID   int
	Text string
}

// LoadDocuments loads CSV documents
func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var docs []Document

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 3 {
			continue
		}
		// combine title + description
		text := row[1] + " " + row[2]

		docs = append(docs, Document{
			ID:   i - 1,
			Text: text,
		})
	}
	return docs, nil
}
