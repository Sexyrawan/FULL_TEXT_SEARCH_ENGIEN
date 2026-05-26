// NOW what this file does
// This is the:
// INVERTED INDEX ENGINE 🔥

package utils

// Index maps words to document IDs
type Index map[string][]int

// Add adds documents to index
func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		tokens := Tokenizer(doc.Text)
		filtered := Filter(tokens)

		for _, token := range filtered {
			// Avoid adding the same document ID multiple times for the same word
			// in the same document (simple deduplication within a document)
			list := idx[token]
			if len(list) == 0 || list[len(list)-1] != doc.ID {
				idx[token] = append(idx[token], doc.ID)
			}
		}
	}
}

// Search finds matching document IDs
func (idx Index) Search(query string) []int {
	tokens := Tokenizer(query)
	filtered := Filter(tokens)

	if len(filtered) == 0 {
		return []int{}
	}
	return idx[filtered[0]]
}
