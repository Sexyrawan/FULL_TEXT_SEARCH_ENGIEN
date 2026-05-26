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

// Search finds matching document IDs using intersection (AND search)
// Isse wahi documents return honge jisme query ke saare words present honge.
func (idx Index) Search(query string) []int {
	tokens := Tokenizer(query)
	filtered := Filter(tokens)

	if len(filtered) == 0 {
		return []int{}
	}

	// Pehle token ke document IDs se shuru karte hain
	result := idx[filtered[0]]

	// Baaki tokens ke document IDs ke sath intersection (AND) karte hain
	for _, token := range filtered[1:] {
		result = intersect(result, idx[token])
		// Agar beech me hi koi list empty ho jaye, toh aage search karne ki zarurat nahi hai
		if len(result) == 0 {
			break
		}
	}

	return result
}

// intersect sorted slices a aur b ka intersection (common elements) nikalta hai.
// Kyunki hamare index me IDs sorted order me append hote hain, hum O(N) time complexity me intersection nikal sakte hain.
func intersect(a, b []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			res = append(res, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return res
}
