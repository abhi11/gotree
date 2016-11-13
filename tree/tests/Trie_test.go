package testSuite

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	//"sort"
	"testing"
)

func TestTrie_development(t *testing.T) {
	assert := assert.New(t)

	// Create a Trie
	trieObj := tree.CreateTrie()
	assert.NotEmpty(trieObj)

	// Insertion
	assert.True(trieObj.Insert("basic Str"))

	// Searching
	assert.True(trieObj.HasVal("basic Str"))
	assert.True(trieObj.HasVal("BASIC STR"))
	assert.True(trieObj.HasVal("basic"))
	assert.False(trieObj.HasVal("avengers"))

	// Create a trie with custom options
	// No substring matching. Only complete words.
	// Case sensitive
	trieObjOpt := tree.CreateTrieWithOptions(false, false)
	assert.NotEmpty(trieObjOpt)

	// Insertion
	assert.True(trieObjOpt.Insert("New String"))

	// Searching
	assert.True(trieObjOpt.HasVal("New String"))
	assert.False(trieObjOpt.HasVal("BASIC STRING"))
	assert.False(trieObjOpt.HasVal("New"))
}

func TestTrie_multiple(t *testing.T) {
	assert := assert.New(t)

	// Create a simple trie
	trieObj := tree.CreateTrie()

	// Inserting
	insertStatus := trieObj.Insert(
		"Letme",
		"tell",
		"you",
		"a",
		"story",
		"to",
		"chill",
		"your",
		"bones",
		"About",
		"a",
		"thing",
		"that",
		"I",
		"saw",
	)
	assert.True(insertStatus)

	assert.True(trieObj.HasVal("tel"))
	assert.True(trieObj.HasVal("let"))
	assert.True(trieObj.HasVal("a"))
	assert.True(trieObj.HasVal("i"))
	assert.True(trieObj.HasVal("chil"))
	assert.True(trieObj.HasVal("bones"))

	assert.False(trieObj.HasVal("tella"))
	assert.False(trieObj.HasVal(""))
	assert.False(trieObj.HasVal("wooot"))
	assert.False(trieObj.HasVal("chiller"))
}

func TestTrie_withOptionsMap(t *testing.T) {
	assert := assert.New(t)

	// Creating a tree with options
	options := map[string]bool{
		"partial_match":      false,
		"case_insensitive":   false,
		"strip_stopwords":    false,
		"strip_punctuations": true,
	}

	trieObj := tree.CreateTrieWithOptionsMap(options)
	trieObj.Insert(
		"Orion", "is", "a", "freaking", "masterpiece!",
	)

	assert.False(trieObj.HasVal("orion"))
	assert.False(trieObj.HasVal("freak"))
	assert.True(trieObj.HasVal("masterpiece"))
}

func TestTrie_stopWords(t *testing.T) {
	assert := assert.New(t)

	options := map[string]bool{
		"strip_stopwords":    true,
		"strip_punctuations": true,
	}

	trieObj := tree.CreateTrieWithOptionsMap(options)

	insertStatus := trieObj.InsertStr(
		`Darkness, Imprisoning me.
		 All that I see, Absolute horror! 
		 I cannot live, 
		 I cannot die,
		 Trapped in myself, 
		 Body my holding cell`,
	)

	assert.True(insertStatus)
}

// Let's Benchmark
func BenchmarkExp_trieSearch(b *testing.B) {
	trieObj := tree.CreateTrie()
	ipList := []interface{}{
		"This", "has", "to", "be", "a", "line", "with", "a", "lot", "of", "words", "and", "some", "numbers", "like", "this", "12", "bigggggword",
	}

	trieObj.Insert(ipList...)
	for i := 0; i < b.N; i++ {
		trieObj.HasVal("This")
	}
}