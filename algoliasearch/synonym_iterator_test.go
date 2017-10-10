package algoliasearch

import "testing"

func TestSynonymIterator(t *testing.T) {
	t.Parallel()
	_, i := initClientAndIndex(t, "TestSynonymIterator")

	synonyms := addObjectsAndSynonyms(t, i, "TestSynonymIterator")

	it := NewSynonymIterator(i)

	var foundSynonyms []Synonym
	var synonym *Synonym
	var err error

	for {
		synonym, err = it.Next()
		if err != nil {
			break
		}
		foundSynonyms = append(foundSynonyms, *synonym)
	}

	if err != NoMoreSynonymsErr {
		t.Fatalf("TestSynonymIterator: Should have stopped iterating because of a %s error but got %s instead", NoMoreSynonymsErr, err)
	}

	if !synonymSlicesAreEqual(synonyms, foundSynonyms) {
		t.Fatalf("TestSynonymIterator: Synonym slices are not equal:\n%v\n%v\n", synonyms, foundSynonyms)
	}
}
